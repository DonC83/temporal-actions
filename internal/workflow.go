package internal

import (
	"go.temporal.io/sdk/workflow"
	"time"
)

func WorkflowRun(ctx workflow.Context, name string) (string, error) {

	options := workflow.ActivityOptions{
		TaskQueue:           ActivityTaskQueue,
		StartToCloseTimeout: time.Second * 5,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	var result string

	err := workflow.ExecuteActivity(ctx, ComposeGreeting, name).Get(ctx, &result)

	err = workflow.ExecuteActivity(ctx, AnotherFunction, name).Get(ctx, &result)

	options2 := workflow.ActivityOptions{
		TaskQueue:           ComplainingTaskQueue,
		StartToCloseTimeout: time.Second * 5,
	}

	complainingCtx := workflow.WithActivityOptions(ctx, options2)
	err = workflow.ExecuteActivity(complainingCtx, ComplainingFunction, name).Get(ctx, &result)

	return result, err
}
