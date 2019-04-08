package queue_go

var (
	MaxWorker = 3
	MaxQueue  = 20
	JobQueue  chan Job
)

func Run()  {
	JobQueue = make(chan Job, MaxQueue)
	//WorkerPool chan chan Job
	dispatcher := NewDispatcher(MaxWorker)

	dispatcher.Run()
}
