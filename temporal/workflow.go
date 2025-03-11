package temporal

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func Scheduler(ctx workflow.Context) error {

	retrypolicy := &temporal.RetryPolicy{
		InitialInterval:    time.Second,
		BackoffCoefficient: 2.0,
		MaximumInterval:    100 * time.Second,
		MaximumAttempts:    500,
	}

	options := workflow.ActivityOptions{
		TaskQueue: TaskQueue,
		// Timeout options specify when to automatically timeout Activity functions.
		StartToCloseTimeout: time.Minute,
		// Optionally provide a customized RetryPolicy.
		// Temporal retries failed Activities by default.
		RetryPolicy: retrypolicy,
	}

	// Apply the options.
	ctx = workflow.WithActivityOptions(ctx, options)

	err := workflow.ExecuteActivity(ctx, FetchIssuersActivity).Get(ctx, nil)

	if err != nil {
		return err
	}
	childWorkflow := workflow.ExecuteChildWorkflow(ctx, MyChildWorkflow)
    // Wait for child to start
    _ = childWorkflow.GetChildWorkflowExecution().Get(ctx, nil)

	return nil
}

func MyChildWorkflow(ctx workflow.Context) error {
	// This workflow will be executed as a child workflow
	// of the parent workflow Scheduler
	return nil
}
