package Engine

import (
	"fmt"
	"log"
)

type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkerCount      int
	ItemChan         chan interface{}
	RequestProcessor Processor
}
type Processor func(Request) (ParseResult, error)
type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}
type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (se *ConcurrentEngine) Run(seeds ...Request) {
	//file, err := os.OpenFile("fiba.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	//defer file.Close()
	//if err != nil {
	//	panic(errors.New("打开文件失败"))
	//}
	out := make(chan ParseResult)
	se.Scheduler.Run()
	for i := 0; i < se.WorkerCount; i += 1 {
		se.createWorker(se.Scheduler.WorkerChan(), out, se.Scheduler)
	}

	for _, seed := range seeds {
		fmt.Printf("seed is --%v\n", seed)
		if isDuplicate(seed.Url) {
			log.Printf("current Url is visited #%s", seed.Url)
			continue
		}
		se.Scheduler.Submit(seed)
	}
	//totalCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			//file.WriteString(fmt.Sprintf("file is -%v\n", item))
			go func() { se.ItemChan <- item }()
			//log.Printf("got item#%d: %v\n", totalCount, item)
			//totalCount++
		}
		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				log.Printf("current Url is visited #%s", request.Url)
				continue
			}
			se.Scheduler.Submit(request)
		}
	}
}

var isVisited = make(map[string]bool)

func isDuplicate(url string) bool {
	if isVisited[url] {
		return true
	}
	isVisited[url] = true
	return false
}
func (ce *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := ce.RequestProcessor(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
