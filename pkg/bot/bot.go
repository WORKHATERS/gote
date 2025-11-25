package bot

import (
	"context"
	"log"

	"gote/pkg/api"
	"gote/pkg/core"
	// "gote/pkg/types"
)

type Bot struct {
	ctx    context.Context
	cancel context.CancelFunc

	api    core.TelegramAPI
	logger core.Logger

	// router core.Router
	// state  core.MemoryState
	// store  core.MemoryStore
	// deps   core.Dependencies
}

func New(ctx context.Context, token string, opts ...Option) *Bot {
	ctx, cancel := context.WithCancel(ctx)

	b := &Bot{
		ctx:    ctx,
		cancel: cancel,
		api:    api.New(token),
		logger: log.Default(),
		// router: router.New(),
		// state:  state.New(),
		// store:  store.New(),
		// deps:   di.New(),
	}

	for _, opt := range opts {
		opt(b)
	}

	return b
}

type Option func(*Bot)

func WithLogger(l core.Logger) Option {
	return func(b *Bot) { b.logger = l }
}

// func WithState(s core.MemoryState) Option {
// 	return func(b *Bot) { b.state = s }
// }

// func WithStore(s core.MemoryStore) Option {
// 	return func(b *Bot) { b.store = s }
// }

// func WithDeps(d core.Dependencies) Option {
// 	return func(b *Bot) { b.deps = d }
// }

// func WithRouter(r core.Router) Option {
// 	return func(b *Bot) { b.router = r }
// }

// func (b *Bot) Handle(update types.Update) {
// 	b.router.Process(b.ctx, b, update)
// }

func (b *Bot) API() core.TelegramAPI    { return b.api }
func (b *Bot) Context() context.Context { return b.ctx }
func (b *Bot) Logger() core.Logger      { return b.logger }

// func (b *Bot) Router() core.Router      { return b.router }
// func (b *Bot) State() core.MemoryState  { return b.state }
// func (b *Bot) Store() core.MemoryStore  { return b.store }
// func (b *Bot) Deps() core.Dependencies  { return b.deps }

func (b *Bot) Stop() {
	b.cancel()
}
