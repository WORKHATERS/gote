package core

import (
	"context"
	"log/slog"
	"net/http"
	"os"
)

// Bot структура бота
type Bot struct {
	ctx    context.Context
	cancel context.CancelFunc

	token  string
	client HTTPClient
	logger Logger
	debug  bool
}

// NewBot функция для создания бота
func NewBot(ctx context.Context, token string, opts ...Option) *Bot {
	ctx, cancel := context.WithCancel(ctx)

	b := &Bot{
		ctx:    ctx,
		cancel: cancel,

		token: token,
		debug: false,
	}

	for _, opt := range opts {
		opt(b)
	}

	if b.client == nil {
		b.client = &http.Client{}
	}

	if b.logger == nil {
		b.logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	}

	return b
}

// Option тип функциональных параметров
type Option func(*Bot)

// WithLogger функция установки значения для логгера
func WithLogger(l Logger) Option {
	return func(b *Bot) { b.logger = l }
}

// WithHTTPClient функция установки значения для HTTP-клиента
func WithHTTPClient(c HTTPClient) Option {
	return func(b *Bot) { b.client = c }
}

// WithDebug функция установки значения для дебаг режима
func WithDebug(on bool) Option {
	return func(b *Bot) { b.debug = on }
}

// Context метод получения контекста
func (b *Bot) Context() context.Context { return b.ctx }

// Logger метод получения логгера
func (b *Bot) Logger() Logger { return b.logger }

// Debug метод получения значения дебаг режима
func (b *Bot) Debug() bool { return b.debug }

// Stop метод остановки бота
func (b *Bot) Stop() {
	b.cancel()
}
