package main

import "fmt"

var WorkerQueue chan chan WorkRequest
var PublisherWorker Publisher

func StartDispatcher(nworkers int) {
  // initialize the channel we are going to but the workers' work channels into.
  WorkerQueue = make(chan chan WorkRequest, nworkers)
  
  // create all of our workers.
  for i := 0; i<nworkers; i++ {
    fmt.Println("Starting worker", i+1)
    worker := NewWorker(i+1, WorkerQueue)
    worker.Start()
  }

  fmt.Println("Starting the publisher")

  PublisherWorker = Publisher {
  	ResultQueue: make(chan Result),
  	QuitChan: make(chan bool)}

  PublisherWorker.Start();
  
  go func() {
    for {
      select {
      case work := <-WorkQueue:
        fmt.Println("Received work requeust")
        go func() {
          worker := <-WorkerQueue
          
          fmt.Println("Dispatching work request")
          worker <- work
        }()
      }
    }
  }()
}