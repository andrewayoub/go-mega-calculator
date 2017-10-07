package main

import (
    "fmt"
    "github.com/go-redis/redis"
)
type Result struct {
    Value   float64
    Cmd     string
    ID      string
}

type Publisher struct {
    ResultQueue chan Result
    QuitChan    chan bool
}

func (P *Publisher) Start(){
    go func() {
        client := redis.NewClient(&redis.Options{
            Addr:     *RedisAddr,
            Password: *Password, 
            DB:       0,  // use default DB
        })
        for {
            select{
            case result := <-P.ResultQueue:
                fmt.Printf("result received %f %s\n", result.Value, result.ID)
                client.LPush(*QueueName + "results", fmt.Sprintf("{'%s' : '%s = %f'}" , result.ID, result.Cmd, result.Value))
            }
        }
    }()
}

func (P *Publisher) Publish(result Result){
    P.ResultQueue <- result
}

// Stop tells the publisher to stop listening for work requests.
//
// Note that the publisher will only stop *after* it has finished its work.
func (P *Publisher) Stop() {
  go func() {
    P.QuitChan <- true
  }()
}