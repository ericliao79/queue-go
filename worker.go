package queue_go

type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

type Job struct {
	Uuid    string
	Data    interface{}
	Query   map[string]string
	PayLoad func(string, interface{})
}

func NewWorker(workerPool chan chan Job) *Worker {
	return &Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool)}
}

func (w Worker) Start() {
	go func() {
		for {
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				// we have received a work request.
				job.PayLoad(job.Uuid, job.Data)

			case <-w.quit:
				// we have received a signal to stop
				return
			}
		}
	}()
}
