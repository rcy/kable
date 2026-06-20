package helloworld

import (
	"context"
	"log"

	"github.com/riverqueue/river"
)

type HelloWorldArgs struct{}

func (HelloWorldArgs) Kind() string { return "hello_world" }

type Worker struct {
	river.WorkerDefaults[HelloWorldArgs]
}

func (w *Worker) Work(ctx context.Context, job *river.Job[HelloWorldArgs]) error {
	log.Println("Hello, world from River!")
	return nil
}
