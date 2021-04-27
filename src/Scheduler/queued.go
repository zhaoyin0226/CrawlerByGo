package Scheduler

import "crawlerByGo/src/Engine"

type QueuedScheduler struct {
	requestChan chan Engine.Request
	workerChan  chan chan Engine.Request
}

func (qs *QueuedScheduler) WorkerChan() chan Engine.Request {
	return make(chan Engine.Request)
}

func (qs *QueuedScheduler) Submit(request Engine.Request) {
	qs.requestChan <- request
}

func (qs *QueuedScheduler) WorkerReady(worker chan Engine.Request) {
	qs.workerChan <- worker
}

func (qs *QueuedScheduler) Run() {
	qs.requestChan = make(chan Engine.Request)
	qs.workerChan = make(chan chan Engine.Request)
	go func() {
		var requestQ []Engine.Request
		var workerQ []chan Engine.Request
		for {
			var activeRequest Engine.Request
			var activeWorker chan Engine.Request
			if len(requestQ) > 0 &&
				len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case request := <-qs.requestChan:
				requestQ = append(requestQ, request)
			case worker := <-qs.workerChan:
				workerQ = append(workerQ, worker)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
