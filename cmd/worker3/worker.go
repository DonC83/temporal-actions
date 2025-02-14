package main

import "temporal-multitaskqueue/internal"

func main() {
	gw := internal.ComplainingWorkerImpl{}
	gw.Start()
}
