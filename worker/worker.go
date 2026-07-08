package worker

import (
	"context"
	"fmt"
	"log"
	"oj/api"
	"oj/worker/helloworld"
	"oj/worker/notifydelivery"
	"oj/worker/notifyfriend"
	"oj/worker/notifykidfriend"
	"oj/worker/youtubedownload"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/riverqueue/river"
	"github.com/riverqueue/river/riverdriver/riverpgxv5"
)

var RiverClient *river.Client[pgx.Tx]

func Start(ctx context.Context, queries *api.Queries, conn *pgxpool.Pool) error {
	workers := river.NewWorkers()
	river.AddWorker(workers, &helloworld.Worker{})
	river.AddWorker(workers, youtubedownload.NewWorker(queries))
	river.AddWorker(workers, &notifydelivery.Worker{Queries: queries, Conn: conn})
	river.AddWorker(workers, &notifyfriend.Worker{Queries: queries, Conn: conn})
	river.AddWorker(workers, &notifykidfriend.Worker{Queries: queries, Conn: conn})

	var err error
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
	job, err := RiverClient.Insert(context.Background(), &notifydelivery.NotifyDeliveryArgs{ID: deliveryID}, nil)
	if err != nil {
		return "", err
	}
	return fmt.Sprint(job.Job.ID), nil
}

func NotifyKidFriend(friendID int64) (string, error) {
	log.Printf("Enqueue NotifyKidFriend %d", friendID)
	job, err := RiverClient.Insert(context.Background(), &notifykidfriend.NotifyKidFriendArgs{ID: friendID}, nil)
	if err != nil {
		return "", err
	}
	return fmt.Sprint(job.Job.ID), nil
}

func NotifyFriend(friendID int64) (string, error) {
	log.Printf("Enqueue NotifyFriend %d", friendID)
	job, err := RiverClient.Insert(context.Background(), &notifyfriend.NotifyFriendArgs{ID: friendID}, nil)
	if err != nil {
		return "", err
	}
	return fmt.Sprint(job.Job.ID), nil
}
