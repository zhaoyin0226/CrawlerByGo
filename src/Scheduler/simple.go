package Scheduler

import "crawlerByGo/src/Engine"

type SimpleScheduler struct {
	workerChan chan Engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan Engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(requests chan Engine.Request) {}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan Engine.Request)
}

func (s *SimpleScheduler) Submit(request Engine.Request) {
	go func() {
		s.workerChan <- request
	}()
}
