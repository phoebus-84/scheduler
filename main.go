package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"os"

	"github.com/phoebus-84/scheduler/temporal"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"go.temporal.io/sdk/client"
)

func main() {
	app := pocketbase.New()

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/{path...}", apis.Static(os.DirFS("./pb_public"), false))
		se.Router.GET("/api/create-schedule", func(e *core.RequestEvent) error {
			scheduleId := e.Request.URL.Query().Get("scheduleId")
			if scheduleId == "" {
				return e.JSON(http.StatusBadRequest, "scheduleId is required")

			}
			createSchedule(app, scheduleId)
			return e.JSON(http.StatusOK, "Schedule created")
		})
		return se.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

func createSchedule(app *pocketbase.PocketBase, scheduleID string) {
	shedulerRecord, err := app.FindRecordById("features", "692se15r284284f")
	if err != nil {
		log.Fatal(err)
	}
	result := struct {
		Interval string `json:"interval"`
	}{}
	errJson := json.Unmarshal([]byte(shedulerRecord.GetString("envVariables")), &result)
	if errJson != nil {
		log.Fatal(errJson)
	}
	var interval time.Duration
	switch result.Interval {
	case "every_minute":
		interval = time.Minute
	case "hourly":
		interval = time.Hour
	case "daily":
		interval = time.Hour * 24
	case "weekly":
		interval = time.Hour * 24 * 7
	case "monthly":
		interval = time.Hour * 24 * 30
	default:
		interval = time.Hour
	}
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
		ID: scheduleID,
		Spec: client.ScheduleSpec{
			Intervals: []client.ScheduleIntervalSpec{
				{
					Every: interval,
				},
			},
		},
		Action: &client.ScheduleWorkflowAction{
			ID:        workflowID,
			Workflow:  temporal.Scheduler,
			TaskQueue: temporal.TaskQueue,
		},
	})

	if err != nil {
		log.Fatalln("Unable to create schedule", err)
	}
	log.Println("Schedule created", "ScheduleID", scheduleID)
	_, _ = scheduleHandle.Describe(ctx)
}
