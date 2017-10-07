package main

import (
  "fmt"
)

// NewWorker creates, and returns a new Worker object. Its only argument
// is a channel that the worker can add itself to whenever it is done its
// work.
func NewWorker(id int, workerQueue chan chan WorkRequest) Worker {
  // Create, and return the worker.
  worker := Worker{
    ID:          id,
    Work:        make(chan WorkRequest),
    WorkerQueue: workerQueue,
    QuitChan:    make(chan bool)}
  
  return worker
}

type Worker struct {
  ID          int
  Work        chan WorkRequest
  WorkerQueue chan chan WorkRequest
  QuitChan    chan bool
}
// This function "starts" the worker by starting a goroutine,
func (w *Worker) Start() {
    go func() {
      for {
        // Add ourselves into the worker queue.
        w.WorkerQueue <- w.Work
        
        select {
        case work := <-w.Work:
          // Receive a work request.
          fmt.Printf("worker%d: Received work request, will excute the job { %s }\n", w.ID, work.ID)
          
          res, err := calculate(work.cmd)
          if err != nil {
            fmt.Println(err)
          } else {
            fmt.Printf("%s = %f {%s}\n", work.cmd, res, work.ID)
            result := Result{Value: res, ID:work.ID, Cmd: work.cmd}
            PublisherWorker.Publish(result)
          }
          //fmt.Printf("worker%d: Hello, %s!\n", w.ID, work.Name)
          
        case <-w.QuitChan:
          // We have been asked to stop.
          fmt.Printf("worker%d stopping\n", w.ID)
          return
        }
      }
    }()
}

// Stop tells the worker to stop listening for work requests.
//
// Note that the worker will only stop *after* it has finished its work.
func (w *Worker) Stop() {
  go func() {
    w.QuitChan <- true
  }()
}