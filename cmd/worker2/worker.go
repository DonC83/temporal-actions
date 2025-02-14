package main

import "temporal-multitaskqueue/internal"

func main() {
	gw := internal.ActivityWorkerImpl{}
	gw.Start()
}
