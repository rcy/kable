package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"oj/api"
	"oj/handlers"
	"oj/handlers/eventsource"
	"oj/services/email"
	"oj/worker"

	"github.com/alexandrevicenzi/go-sse"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed to connect to postgres: %s", err)
	}
	defer pool.Close()

	conn, err := pool.Acquire(ctx)
	if err != nil {
		log.Fatalf("Could not aquire connection", err)
	}

	err = worker.Start(context.Background(), api.New(pool), conn)
	if err != nil {
		log.Fatalf("could not start worker: %s", err)
	}

	go func() {
		count := 0
		for {
			id := fmt.Sprint(count)
			data := time.Now().Format(time.RFC3339Nano)
			eventsource.SSE.SendMessage("", sse.NewMessage(id, data, "KEEP_ALIVE"))
			count += 1
			time.Sleep(30 * time.Second)
		}
	}()

	err = email.Send("kable startup", "application started", os.Getenv("DEV_EMAIL"))
	if err != nil {
		log.Fatalf("failed to send startup email: %s", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	handler := handlers.Router(conn)

	log.Printf("listening on port %s", port)
	err = http.ListenAndServe(":"+port, handler)
	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("server closed\n")
	} else if err != nil {
		log.Printf("server closed unexpectedly: %v\n", err)
		os.Exit(1)
	}
}
