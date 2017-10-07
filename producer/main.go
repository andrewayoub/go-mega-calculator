package main

import (
	"fmt"
	"flag"
	"bufio"
	"os"
	"strings"
)

var (
  RedisAddr = flag.String("addr", "localhost:6379", "Redit server address")
  Password = flag.String("pass", "", "Redit server address")
  QueueName = flag.String("queue", "queue1", "queue name")
)

func main() {
	//parse flags
	flag.Parse()


	/*//initialize redis client
    client := redis.NewClient(&redis.Options{
            Addr:     *RedisAddr,
            Password: *Password, 
            DB:       0,  // use default DB
        })


	fmt.Println("sending to " + *QueueName)*/

	sender := Sender{}
	sender.Init()

	
	
	go StartReceiving()

    reader := bufio.NewReader(os.Stdin)
    for {
		fmt.Println("Enter operation:")
		//get operation from user input
	    operation, _ := reader.ReadString('\n')

	    sender.Send(strings.TrimRight(operation, "\n"))
    }



}