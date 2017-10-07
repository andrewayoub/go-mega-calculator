package main

import (
  "flag"
  "fmt" 
  "strconv"
)
import s "strings"
var (
  RedisAddr = flag.String("addr", "localhost:6379", "Redit server address")
  Password = flag.String("pass", "", "Redit server address")
  QueueAndNumber = flag.String("queue", "queue1:4", "queue name")
)
var NWorkers int
var QueueName string

func main() {
  // Parse the command-line flags.
  flag.Parse()
  queueValue := s.Split(*QueueAndNumber, ":")
  QueueName = queueValue[0]
  NWorkers,err := strconv.Atoi(queueValue[1])
  if err != nil || len(QueueName) < 1 {
    fmt.Println("Invalid args")
    return
  }
  // Start the dispatcher.
  fmt.Println("Starting the dispatcher")
  StartDispatcher(NWorkers)
  
  // Register our collector as a radis client.
  //collect := make(chan string)
  Collector()
  
}