package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/NathanBak/cfgbuild"
	"github.com/joho/godotenv"

	"github.com/NathanBak/go-server/internal/server"
	"github.com/NathanBak/go-server/pkg/storage"
	"github.com/NathanBak/go-server/pkg/widget"
)

func main() {

	_ = godotenv.Load("local.env")
	_ = godotenv.Load()

	builder := cfgbuild.Builder[*server.Config]{}

	cfg, err := builder.Build()
	if err != nil {
		log.Fatal(err)
	}

	cfg.Storage = &storage.MapStorage[widget.Widget]{}

	s, err := server.New(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Start server running on separate thread
	go func() {
		err = s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// wait for signal and then shutdown cleanly
	quitchan := make(chan os.Signal, 1)
	signal.Notify(quitchan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quitchan
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = s.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
