package worker

import (
	"context"
	"log"
	"oj/api"
	"oj/worker/helloworld"
	"oj/worker/notifydelivery"
	"oj/worker/notifyfriend"
	"oj/worker/notifykidfriend"
	"oj/worker/youtubedownload"
	"time"

	"github.com/acaloiaro/neoq"
	"github.com/acaloiaro/neoq/handler"
	"github.com/acaloiaro/neoq/jobs"
	"github.com/acaloiaro/neoq/types"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/riverqueue/river"
	"github.com/riverqueue/river/riverdriver/riverpgxv5"
)

var Queue types.Backend
var RiverClient *river.Client[pgx.Tx]

func Start(ctx context.Context, queries *api.Queries, conn *pgxpool.Pool) error {
	var err error
	Queue, err = neoq.New(ctx)
	if err != nil {
		return err
	}

	Queue.Start(ctx, "notify-delivery", handler.New(notifydelivery.NewService(queries, conn).Handle))
	Queue.Start(ctx, "notify-friend", handler.New(notifyfriend.NewService(queries, conn).Handle))
	Queue.Start(ctx, "notify-kid-friend", handler.New(notifykidfriend.NewService(queries, conn).Handle))

	workers := river.NewWorkers()
	river.AddWorker(workers, &helloworld.Worker{})
	river.AddWorker(workers, youtubedownload.NewWorker(queries))

	RiverClient, err = river.NewClient(riverpgxv5.New(conn), &river.Config{
		Queues: map[string]river.QueueConfig{
			river.QueueDefault: {MaxWorkers: 1},
		},
		Workers: workers,
	})
	if err != nil {
		return err
	}

	if err := RiverClient.Start(ctx); err != nil {
		return err
	}

	log.Print("started worker")

	return nil
}

func NotifyDelivery(deliveryID int64) (string, error) {
	return Queue.Enqueue(context.Background(), &jobs.Job{
		Queue:    "notify-delivery",
		Payload:  map[string]any{"id": deliveryID},
		RunAfter: time.Now().Add(1 * time.Second),
	})
}

func NotifyFriend(friendID int64) (string, error) {
	log.Printf("Enqueue NotifyFriend %d", friendID)
	return Queue.Enqueue(context.Background(), &jobs.Job{
		Queue:   "notify-friend",
		Payload: map[string]any{"id": friendID},
	})
}

func NotifyKidFriend(friendID int64) (string, error) {
	log.Printf("Enqueue NotifyKidFriend %d", friendID)
	return Queue.Enqueue(context.Background(), &jobs.Job{
		Queue:   "notify-kid-friend",
		Payload: map[string]any{"id": friendID},
	})
}
