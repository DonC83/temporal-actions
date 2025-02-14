package internal

import (
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
)

type WorkerInterface interface {
	Start()
	Stop()
}

type WorkflowWorkerImpl struct {
}

func (g *WorkflowWorkerImpl) Start() {
	c, err := client.Dial(client.Options{
		HostPort: "127.0.0.1:7233",
	})
	if err != nil {
		log.Fatalln("unable to create temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, WorkflowTaskQueue, worker.Options{})
	w.RegisterWorkflow(WorkflowRun)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start greeting worker", err)
	}
}

func (g *WorkflowWorkerImpl) Stop() {
	//TODO implement me
	panic("implement me")
}

type ActivityWorkerImpl struct{}

func (c *ActivityWorkerImpl) Start() {
	cl, err := client.Dial(client.Options{
		HostPort: "127.0.0.1:7233",
	})
	if err != nil {
		log.Fatalln("unable to create temporal client", err)
	}
	defer cl.Close()

	w := worker.New(cl, ActivityTaskQueue, worker.Options{})
	w.RegisterActivity(ComposeGreeting)
	w.RegisterActivity(AnotherFunction)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start complain worker", err)
	}
}

func (c *ActivityWorkerImpl) Stop() {
}

type ComplainingWorkerImpl struct{}

func (c *ComplainingWorkerImpl) Start() {
	cl, err := client.Dial(client.Options{
		HostPort: "127.0.0.1:7233",
	})
	if err != nil {
		log.Fatalln("unable to create temporal client", err)
	}
	defer cl.Close()

	w := worker.New(cl, ComplainingTaskQueue, worker.Options{
		WorkerActivitiesPerSecond:    1,
		TaskQueueActivitiesPerSecond: 1,
	})
	w.RegisterActivity(ComplainingFunction)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start complain worker", err)
	}
}

func (c *ComplainingWorkerImpl) Stop() {
}
