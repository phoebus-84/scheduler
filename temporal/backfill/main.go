package main

import (
	"context"
	"log"
	"time"

	schedule "github.com/phoebus-84/scheduler/temporal"
	"go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/client"
)


func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create a Workflow to backfill
	temporalClient, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create Temporal Client", err)
	}
	defer temporalClient.Close()

	workflowID := "schedule_workflow_"
	// create paused Workflow
	now := time.Now()
	scheduleHandle, _ := temporalClient.ScheduleClient().Create(ctx, client.ScheduleOptions{
		ID: "backfill-schedule",
		Spec: client.ScheduleSpec{
			Intervals: []client.ScheduleIntervalSpec{
				{
					Every: time.Hour * 24,
				},
			},
		},
		Action: &client.ScheduleWorkflowAction{
			ID:        workflowID,
			Workflow:  schedule.Scheduler,
			TaskQueue: schedule.TaskQueue,
		},
		Paused: true,
	})

	err = scheduleHandle.Backfill(ctx, client.ScheduleBackfillOptions{
		Backfill: []client.ScheduleBackfill{
			{
				Start:   now.Add(-4 * time.Minute),
				End:     now.Add(-2 * time.Minute),
				Overlap: enums.SCHEDULE_OVERLAP_POLICY_ALLOW_ALL,
			},
			{
				Start:   now.Add(-2 * time.Minute),
				End:     now,
				Overlap: enums.SCHEDULE_OVERLAP_POLICY_ALLOW_ALL,
			},
		},
	})
	if err != nil {
		log.Fatalln("Unable to Backfill Schedule", err)
	}

}