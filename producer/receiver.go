package main

import (
  "fmt"
  "github.com/go-redis/redis"
)
var Queue = make(chan string, 100)
func StartReceiving() {
  
  client := redis.NewClient(&redis.Options{
    Addr:     *RedisAddr,
    Password: *Password, 
    DB:       0,  // use default DB
  })

  for {
    
    command,_ := client.BLPop(0, *QueueName + "results").Result()
    res := command[1]
    select {
      case Queue <- res:
        fmt.Println(command[1])
    }
  }
  return
}
