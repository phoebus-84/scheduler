package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/phoebus-84/scheduler/temporal"
)

func main() {

	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client.", err)
	}
	defer c.Close()

	w := worker.New(c, temporal.TaskQueue, worker.Options{})

	// This worker hosts both Workflow and Activity functions.
	w.RegisterWorkflow(temporal.Scheduler)
	w.RegisterActivity(temporal.FetchIssuersActivity)

	// Start listening to the Task Queue.
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}