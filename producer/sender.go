package main

import (
  "github.com/go-redis/redis"
  "fmt"
)

type Sender struct {
  Client *redis.Client
}

func (s *Sender) Init() {
  
  client := redis.NewClient(&redis.Options{
    Addr:     *RedisAddr,
    Password: *Password, 
    DB:       0,  // use default DB
  })
  s.Client = client
}

func (s *Sender) Send(command string) {
  id, err := s.Client.Get(*QueueName+"i").Result()
  if err!= nil{
    id="0";
    s.Client.Set(*QueueName+"i", id, 0)
  }
  s.Client.LPush(*QueueName,command + "&" + id)
  fmt.Println("sent with id:" + id)
  s.Client.Incr(*QueueName+"i")
}