package router

import (
	"context"
	"strings"
	"sync"

	"gote/pkg/core"
	"gote/pkg/types"
)

type HandlerFunc func(ctx context.Context, bot core.BotContext, update types.Update) error

type MiddlewareFunc func(next HandlerFunc) HandlerFunc

type Router struct {
	mu               sync.RWMutex
	handlers         []handlerEntry
	middlewares      []MiddlewareFunc
	globalMiddleware []MiddlewareFunc
}

type handlerEntry struct {
	kind    handlerKind
	pattern string
	handler HandlerFunc
}

type handlerKind int

const (
	kindCommand handlerKind = iota
	kindTextContains
	kindCallbackPrefix
	kindMessage
	kindAny
)

func New() *Router {
	return &Router{}
}

// Use подключает обработчики, middleware и модули
func (r *Router) Use(things ...any) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, t := range things {
		switch v := t.(type) {
		case func(HandlerFunc) HandlerFunc:
			r.middlewares = append(r.middlewares, MiddlewareFunc(v))
		case MiddlewareFunc:
			r.middlewares = append(r.middlewares, v)
		case func(*Registrar):
			v(&Registrar{r: r})
		case *Router:
			r.handlers = append(r.handlers, v.handlers...)
			r.middlewares = append(r.middlewares, v.middlewares...)
			r.globalMiddleware = append(r.globalMiddleware, v.globalMiddleware...)
		default:
			panic("router.Use: unsupported type")
		}
	}
}

func (r *Router) UseGlobal(mw MiddlewareFunc) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.globalMiddleware = append(r.globalMiddleware, mw)
}

// Registrar — для модулей
type Registrar struct{ r *Router }

func (reg *Registrar) Command(cmd string, h HandlerFunc) {
	reg.r.Command(cmd, h)
}

func (reg *Registrar) OnMessage(h HandlerFunc) {
	reg.r.HandleMessage(h)
}

func (reg *Registrar) OnTextContains(text string, h HandlerFunc) {
	reg.r.TextContains(text, h)
}

func (reg *Registrar) Middleware(mw MiddlewareFunc) {
	reg.r.Use(mw)
}

// Публичные методы
func (r *Router) Command(cmd string, h HandlerFunc) {
	r.add(kindCommand, normalizeCommand(cmd), h)
}

func (r *Router) TextContains(text string, h HandlerFunc) {
	r.add(kindTextContains, strings.ToLower(text), h)
}

func (r *Router) HandleMessage(h HandlerFunc) {
	r.add(kindMessage, "", h)
}

func (r *Router) HandleAny(h HandlerFunc) {
	r.add(kindAny, "", h)
}

func (r *Router) add(kind handlerKind, pattern string, h HandlerFunc) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.handlers = append(r.handlers, handlerEntry{kind: kind, pattern: pattern, handler: h})
}

// Process запускает обработчики и прослойки
func (r *Router) Process(ctx context.Context, bot core.BotContext, update types.Update) {
	// базовый next — просто пустая функция
	next := func(ctx context.Context, bot core.BotContext, update types.Update) error {
		handlers := r.matchHandlers(update)
		if len(handlers) == 0 {
			return nil
		}

		// строим цепочку обработчиков
		hNext := func(ctx context.Context, bot core.BotContext, update types.Update) error { return nil }
		for i := len(handlers) - 1; i >= 0; i-- {
			h := handlers[i]
			currentNext := hNext
			hNext = func(ctx context.Context, bot core.BotContext, update types.Update) error {
				_ = h(ctx, bot, update)
				return currentNext(ctx, bot, update)
			}
		}

		// применяем middleware для фильтрованных апдейтов
		mNext := hNext
		for i := len(r.middlewares) - 1; i >= 0; i-- {
			mNext = r.middlewares[i](mNext)
		}

		return mNext(ctx, bot, update)
	}

	// применяем глобальные middleware поверх всего
	for i := len(r.globalMiddleware) - 1; i >= 0; i-- {
		next = r.globalMiddleware[i](next)
	}

	_ = next(ctx, bot, update)
}

// matchHandlers возвращает обработчики в порядке регистрации
func (r *Router) matchHandlers(update types.Update) []HandlerFunc {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var matches []HandlerFunc
	for _, h := range r.handlers {
		if r.matches(h, update) {
			matches = append(matches, h.handler)
		}
	}
	return matches
}

// сопоставление типа обработчика с типом обновления
func (r *Router) matches(entry handlerEntry, update types.Update) bool {
	switch entry.kind {
	case kindCommand:
		if msg := update.Message; msg != nil && len(msg.Text) > 0 && string(msg.Text[0]) == "/" {
			return normalizeCommand(msg.Text) == entry.pattern
		}
	case kindTextContains:
		if msg := update.Message; msg != nil && msg.Text != "" {
			return strings.Contains(strings.ToLower(msg.Text), entry.pattern)
		}
	case kindCallbackPrefix:
		if cb := update.CallbackQuery; cb != nil {
			return strings.HasPrefix(cb.Data, entry.pattern)
		}
	case kindMessage:
		return update.Message != nil
	case kindAny:
		return true
	}
	return false
}

// нормализация команты /Text -> text
func normalizeCommand(cmd string) string {
	cmd = strings.TrimPrefix(cmd, "/")
	cmd = strings.ToLower(cmd)
	if i := strings.Index(cmd, "@"); i != -1 {
		cmd = cmd[:i]
	}
	return cmd
}
