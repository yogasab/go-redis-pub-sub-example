package main

import (
	"encoding/json"
	config "go-redis-pub-sub-example/config"
	"go-redis-pub-sub-example/model"
	"log"
)

func main() {
	log.Println("Listening all the subscribed message ...")
	subscriber := config.RedisClient.Subscribe(config.Ctx, "send-user-data")

	user := model.User{}

	for {
		message, err := subscriber.ReceiveMessage(config.Ctx)

		if err != nil {
			panic(err)
		}

		if err := json.Unmarshal([]byte(message.Payload), &user); err != nil {
			panic(err)
		}

		log.Println("Received message from " + message.Channel + " channel.")
		log.Printf("%+v\n", user)
	}
}
