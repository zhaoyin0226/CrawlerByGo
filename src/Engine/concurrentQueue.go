package Engine

import (
	"log"
)

type ConcurrentQueueEngine struct {
	Scheduler   QScheduler
	WorkerCount int
}
type QScheduler interface {
	Submit(Request)
	WorkerChan() chan Request
	WorkerReady(chan Request)
	Run()
}

func (ce *ConcurrentQueueEngine) RunEngine(seeds ...Request) {

	out := make(chan ParseResult)
	ce.Scheduler.Run()
	for i := 0; i < ce.WorkerCount; i += 1 {
		createQWorker(out, ce.Scheduler)
	}

	for _, seed := range seeds {
		ce.Scheduler.Submit(seed)
	}
	totalCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("got item#%d: %v\n", totalCount, item)
			totalCount++
		}
		for _, request := range result.Requests {
			ce.Scheduler.Submit(request)
		}
	}

}
func createQWorker(out chan ParseResult, qs QScheduler) {
	in := make(chan Request)
	go func() {
		for {
			qs.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
