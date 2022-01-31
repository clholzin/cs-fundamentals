package main

import (
	"context"
	"cs-fundamentals/kube/redis_main_replicas/repo"
	"flag"
	"log"

	"github.com/go-redis/redis/v8"
)

var (
	stream     = flag.String("stream", "pubsub", "stream name")
	password   = flag.String("password", "4W9yCzVuLi", "redis password")
	id         = flag.String("id", "0", "stream id")
	connection = flag.String("addr", "127.0.0.1:6379", "client address")
)

func main() {
	log.Printf("Starting up on %s %s", *stream, *connection)
	flag.Parse()
	dbclient := repo.NewRepo(*connection, *password, 0)
	receiver := NewReceiver(*stream, *id, dbclient)
	receiver.Listener()
	log.Println("shutting down")
}

func NewReceiver(stream, id string, client *redis.Client) *Receiver {
	r := &Receiver{}
	r.name = stream
	ctx := context.Background()
	r.ctx = ctx
	r.repo = client
	r.id = id
	return r
}

type Receiver struct {
	name   string
	ctx    context.Context
	repo   *redis.Client
	id     string
	offset string
}

func (r *Receiver) Listener() {
	for {
		xstream := r.repo.XRead(r.ctx, &redis.XReadArgs{
			Streams: []string{r.name, r.id},
			Count:   1,
			Block:   0, //time.Duration(2 * time.Second),
		})
		streamData, err := xstream.Result()
		if err != nil {
			log.Printf("failed to receive result with: %v", err)
			continue
		}
		for _, xstream := range streamData {
			log.Printf("stream name: %s\n", xstream.Stream)
			for _, xmsg := range xstream.Messages {
				r.id = xmsg.ID
				log.Printf("received %+v\n", xmsg.Values)
			}
		}
	}
}
