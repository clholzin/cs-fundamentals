package main

import (
	"context"
	"cs-fundamentals/kube/redis_main_replicas/repo"
	"flag"
	"log"
	"strings"

	"github.com/go-redis/redis/v8"
)

var (
	input      = flag.String("input", "say,hello", "set input")
	stream     = flag.String("stream", "pubsub", "stream name")
	password   = flag.String("password", "4W9yCzVuLi", "redis password")
	id         = flag.String("id", "*", "stream id")
	connection = flag.String("addr", "127.0.0.1:6379", "client address")
)

func main() {

	dbclient := repo.NewRepo(*connection, *password, 0)
	sender := NewSender(*stream, *id, dbclient)
	flag.Parse()
	dataToSend := *input
	parsedkv := strings.Split(dataToSend, ",")
	err := sender.Write(parsedkv)
	if err != nil {
		log.Println(err)
	}
	log.Println(*input)
}

func NewSender(stream, id string, client *redis.Client) *Sender {
	sender := &Sender{}
	ctx := context.Background()
	sender.repo = client
	sender.ctx = ctx
	sender.name = stream
	sender.id = id
	return sender
}

type Sender struct {
	name string
	repo *redis.Client
	ctx  context.Context
	id   string
}

func (s *Sender) Write(input []string) error {
	args := &redis.XAddArgs{
		Stream: s.name,
		ID:     s.id,
	}
	kv := make(map[string]interface{})
	var i int
	for i < len(input) && i+1 < len(input) {
		kv[input[i]] = input[i+1]
		i = i + 2
	}
	args.Values = kv
	scmd := s.repo.XAdd(s.ctx, args)
	return scmd.Err()
}
