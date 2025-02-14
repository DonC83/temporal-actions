package main

import "temporal-multitaskqueue/internal"

func main() {
	gw := internal.WorkflowWorkerImpl{}
	gw.Start()
}
