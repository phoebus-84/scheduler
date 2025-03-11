package main

import (
	"context"
	"log"
	"time"

	schedule "github.com/phoebus-84/scheduler/temporal"
	"go.temporal.io/sdk/client"
)
 
func main() {
	scheduleID := "schedule_id"
	workflowID := "schedule_workflow_id"
	ctx := context.Background()
	
	temporalClient, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create Temporal Client", err)
	}
	defer temporalClient.Close()

	scheduleHandle, err := temporalClient.ScheduleClient().Create(ctx, client.ScheduleOptions{
		ID:   scheduleID,
		Spec: client.ScheduleSpec{
			Intervals: []client.ScheduleIntervalSpec{
				{
					Every: time.Hour,
				},
			},
		},
		Action: &client.ScheduleWorkflowAction{
			ID:        workflowID,
			Workflow:  schedule.Scheduler,
			TaskQueue: schedule.TaskQueue,
		},
	})

	if err != nil {
		log.Fatalln("Unable to create schedule", err)
	}
	log.Println("Schedule created", "ScheduleID", scheduleID)
	_, _ = scheduleHandle.Describe(ctx)
}