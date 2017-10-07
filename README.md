# go-mega-calculator
arithmatic calculator that implements producer consumer pattern

this project is done as an interview task for CodeScalers Egypt

this is a simple producer consumer implementation that uses Go Language and Redis

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

1- Go </br>
2- Redis Server </br>
3- [Go-Redis](https://github.com/go-redis/redis) (a redis client for go) </br>

### Installing
you need to build both Producer and Agent but first make sure that you have Go-Redis in your GOPATH
Run the following commands:

```shell
go build -o agent/agent agent/*.go
go build -o producer/producer producer/*.go
```
then run these binary files in different terminal windows

### How to use

to run the agent 
```shell
agent/agent [-queue {queue name}] [-n number of workers] [-addr {redis server address}] [-pass  {redi server password}]
```
to run the producer 
```shell
producer/producer [-queue {queue name}] [-addr {redis server address}] [-pass  {redi server password}]
```
now in producer screen you can enter an arithmatic operation and it will be sent to the calulator then to the result queue
example : 1+2+3*4/5



### TODO
this task still considered incomplete
* write unit tests
* allow multi-queues 
* change the CLI to be as mentioned in task description
