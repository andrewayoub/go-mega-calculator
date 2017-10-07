package main

import (
  "fmt"
  "github.com/go-redis/redis"
)
import s "strings"
// A buffered channel that we can send work requests on.
var WorkQueue = make(chan WorkRequest, 100)

func Collector() {
  
  client := redis.NewClient(&redis.Options{
    Addr:     *RedisAddr,
    Password: *Password, 
    DB:       0,  // use default DB
  })

  for {
    fmt.Println("Collector is running")
    command,_ := client.BLPop(0, QueueName).Result()
    
    commands := s.Split(command[1],"&")

    work := WorkRequest{cmd: commands[0],ID: commands[1]}
    
    select {
      case WorkQueue <- work:
        fmt.Println("Work request queued")
    }
  }
  return
}
