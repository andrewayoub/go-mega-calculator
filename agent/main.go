package main

import (
  "flag"
  "fmt"
  
)

var (
  NWorkers = flag.Int("n", 4, "The number of workers to start")
  RedisAddr = flag.String("addr", "localhost:6379", "Redit server address")
  Password = flag.String("pass", "", "Redit server address")
  QueueName = flag.String("queue", "queue1", "queue name")
)

func main() {
  // Parse the command-line flags.
  flag.Parse()
  
  // Start the dispatcher.
  fmt.Println("Starting the dispatcher")
  StartDispatcher(*NWorkers)
  
  // Register our collector as a radis client.
  collect := make(chan string)
  Collector(collect)
  
}