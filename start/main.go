package main

import (
	"context"
	"fmt"
	"go.temporal.io/sdk/client"
	"log"
	"temporal-multitaskqueue/internal"
)

func main() {
	callWorkflow()
}

func callWorkflow() {
	c, err := client.Dial(client.Options{
		HostPort: "127.0.0.1:7233",
	})
	if err != nil {
		log.Fatalln("unable to create temporal client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		TaskQueue: internal.WorkflowTaskQueue,
	}

	ctx := context.Background()
	name := "World"
	we, err := c.ExecuteWorkflow(ctx, options, internal.WorkflowRun, name)
	if err != nil {
		log.Fatalln("unable to complete workflow", err)
	}

	var greeting string
	if err = we.Get(ctx, &greeting); err != nil {
		log.Fatalln("unable to get workflow result", err)
	}
	fmt.Println(we.GetID(), we.GetID())
	fmt.Printf("result %s", greeting)
}
