package main

import (
	"context"
	"gitgub.com/terratensor/quote-app/api/quotesrv/internal/api/server"
	"gitgub.com/terratensor/quote-app/api/quotesrv/internal/app/repos/quote"
	"gitgub.com/terratensor/quote-app/api/quotesrv/internal/app/starter"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	if tz := os.Getenv("TZ"); tz != "" {
		var err error
		time.Local, err = time.LoadLocation(tz)
		if err != nil {
			log.Printf("error loading location '%s': %v\n", tz, err)
		}
	}

	// output current time zone
	tnow := time.Now()
	tz, _ := tnow.Zone()
	log.Printf("Local time zone %s. Service started at %s", tz,
		tnow.Format("2006-01-02T15:04:05.000 MST"))

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	var qst quote.Store
	a := starter.NewApp(qst)
	qs := quote.NewQuotes(qst)
	// ToDo сделать handler и routeroapi
	h := handler.NewHandlers(qs)
	rh := routeroapi.NewRouterOpenAPI(h)

	srv := server.NewServer(os.Getenv("SERVER_ADDR"), rh)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go a.Serve(ctx, wg, srv)

	<-ctx.Done()
	cancel()
	wg.Wait()
}
