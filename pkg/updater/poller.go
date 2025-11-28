package updater

import (
	"time"

	"gote/pkg/core"
	"gote/pkg/types"
)

// Updater интерфейс для получателей обновлений
type Updater interface {
	Start() <-chan types.Update
}

// Poller структура для работы с Long polling запросами
type Poller struct {
	bot          *core.Bot
	params       types.GetUpdates
	errorBackoff time.Duration
	bufferSize   int64
}

// PollerOption тип функциональных параметров
type PollerOption func(*Poller)

// NewPoller функция-конструктор для Poller
func NewPoller(b *core.Bot, opts ...PollerOption) *Poller {
	p := &Poller{
		bot: b,
		params: types.GetUpdates{
			Timeout: 30,
			Limit:   100,
		},
		errorBackoff: 5 * time.Second,
		bufferSize:   100,
	}

	for _, opt := range opts {
		opt(p)
	}

	return p
}

// WithTimeout функция установки значения времени ожидания
func WithTimeout(t int64) PollerOption {
	return func(p *Poller) { p.params.Timeout = t }
}

// WithLimit функция установки значения количества получаемых обновлений
func WithLimit(l int64) PollerOption {
	return func(p *Poller) { p.params.Limit = l }
}

// WithAllowedUpdates функция установки значения принимаемых типов обновлений
func WithAllowedUpdates(au []string) PollerOption {
	return func(p *Poller) { p.params.AllowedUpdates = au }
}

// WithErrorBackoff функция установки значения времени ожидания при повторном запросе в случае ошибки
func WithErrorBackoff(d time.Duration) PollerOption {
	return func(p *Poller) { p.errorBackoff = d }
}

// WithUpdatesBufferSize функция установки размера буфера обновлений
func WithUpdatesBufferSize(size int64) PollerOption {
	return func(p *Poller) { p.bufferSize = size }
}

// Start метод получения обвновлений
func (p *Poller) Start() <-chan types.Update {
	ch := make(chan types.Update, p.bufferSize)

	go func() {
		defer close(ch)

		for {
			if p.bot.Context().Err() != nil {
				return
			}

			updates, err := p.bot.GetUpdates(p.bot.Context(), p.params)
			if err != nil {
				p.bot.Logger().Error("Ошибка получения обновлений", err)
				select {
				case <-time.After(p.errorBackoff):
				case <-p.bot.Context().Done():
					return
				}
				continue
			}

			if n := len(updates); n > 0 {
				for _, u := range updates {
					select {
					case ch <- u:
					case <-p.bot.Context().Done():
						return
					}
				}

				p.params.Offset = updates[n-1].UpdateId + 1
			}
		}
	}()

	return ch
}
