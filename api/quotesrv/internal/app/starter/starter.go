package starter

import (
	"context"
	"gitgub.com/terratensor/quote-app/api/quotesrv/internal/app/repos/quote"
	"sync"
)

type App struct {
	qs *quote.Quotes
}

func NewApp(qst quote.Store) *App {
	a := &App{
		qs: quote.NewQuotes(qst),
	}
	return a
}

type HTTPServer interface {
	Start(qs *quote.Quotes)
	Stop()
}

func (a *App) Serve(ctx context.Context, wg *sync.WaitGroup, hs HTTPServer) {
	defer wg.Done()
	hs.Start(a.qs)
	<-ctx.Done()
	hs.Stop()
}
