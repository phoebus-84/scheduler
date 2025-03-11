package main

import (
	// "context"
	"log"
	// "net/http"
	"os"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	// "go.temporal.io/sdk/client"
)

func main() {
	app := pocketbase.New()

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// serves static files from the provided public dir (if exists)
		se.Router.GET("/{path...}", apis.Static(os.DirFS("./pb_public"), false))
		// se.Router.GET("/api/validate-yaml", func(e *core.RequestEvent) error {
		// 	yaml := e.Request.URL.Query().Get("yaml")
		// 	if yaml == "" {
		// 		return e.JSON(http.StatusBadRequest, "yaml query param is required")
		// 	}
		// 	email := e.Request.URL.Query().Get("email")
		// 	if email == "" {
		// 		return e.JSON(http.StatusBadRequest, "email query param is required")
		// 	}
		// 	ctx, err := client.Dial(client.Options{})

		// 	if err != nil {
		// 		log.Fatalln("Unable to create Temporal client:", err)
		// 	}

		// 	defer ctx.Close()
		// 	input := update.AppInput{
		// 		Email: email,
		// 		Url:   yaml,
		// 		App: app,
		// 	}
		// 	options := client.StartWorkflowOptions{
		// 		ID:        "validate-id",
		// 		TaskQueue: update.TaskQueue,
		// 	}
		// 	log.Printf("Starting Workflow with ID: %s\n, input: %v", options.ID, input)
		// 	we, err := ctx.ExecuteWorkflow(context.Background(), options, update.Validation, input)
		// 	if err != nil {
		// 		log.Fatalln("Unable to start the Workflow:", err)
		// 	}
		// 	log.Printf("WorkflowID: %s RunID: %s\n", we.GetID(), we.GetRunID())
		// 	var result string
		// 	err = we.Get(context.Background(), &result)
		// 	if err != nil {
		// 		log.Fatalln("Unable to get Workflow result:", err)
		// 	}
		// 	if err != nil {
		// 		log.Fatalln("Unable to get Workflow result:", err)
		// 	}
		// 	log.Println(result)
		// 	return e.JSON(http.StatusOK, &result)
		// })
		return se.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
