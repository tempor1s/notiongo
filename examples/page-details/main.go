package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/tempor1s/notiongo/notion"
	"log"
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
	// retrieve page
	page, err := instance.Page.Retrieve(context.Background(), "20c4a81f-4d0f-4b7f-a2d4-27ff6582753d")
	if err != nil {
		log.Println("err:", err)
		return
	}

	fmt.Printf("Page: %+v \n", page)

}
