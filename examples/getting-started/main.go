package main

import (
	"context"
	"flag"
	"log"

	"github.com/tempor1s/notiongo/notion"
)

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Notion Secret Token")
	flag.Parse()
}

func main() {
	// create the notion client to interact with the api
	instance := notion.NewClient(Token, nil)
	log.Printf("%+v\n", instance)

	users, err := instance.User.List(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("-- all users --")
	for _, user := range users {
		log.Println(user)
	}

	log.Println("-- single user --")
	user, err := instance.User.Get(context.Background(), "29dbbde7-e391-46eb-b2cf-e5a61748bd63")
	if err != nil {
		log.Println("err", err)
		return
	}

	log.Println(user)
}
