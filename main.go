package main

import (
	"flag"
	"log"

	"github.com/neomede/chattp2/src/chattp2"
	"golang.org/x/net/context"
)

func main() {
	client := flag.Bool("client", false, "Client mode.")
	userSender := flag.String("sender", "sender", "User sender.")
	userReceiver := flag.String("receiver", "receiver", "User receiver.")
	flag.Parse()

	if !*client {
		server, err := chattp2.NewServer()
		if err != nil {
			log.Println(err.Error())
			return
		}

		server.Run(context.Background())
	} else {
		client, err := chattp2.NewClient(*userSender, *userReceiver)
		if err != nil {
			log.Println(err.Error())
			return
		}

		client.Run(context.Background())
	}
}
