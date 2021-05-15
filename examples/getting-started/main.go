package main

import (
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
}
