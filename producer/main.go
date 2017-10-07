package main

import (
	"fmt"
	"github.com/go-redis/redis"
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


	//initialize redis client
    client := redis.NewClient(&redis.Options{
            Addr:     *RedisAddr,
            Password: *Password, 
            DB:       0,  // use default DB
        })


    reader := bufio.NewReader(os.Stdin)
	fmt.Println("sending to " + *QueueName)
	
    for {
		fmt.Println("Enter operation:")
	    operation, _ := reader.ReadString('\n')

	    client.LPush(*QueueName,strings.TrimRight(operation, "\n") + "&1")
    }



}