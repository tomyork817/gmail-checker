package app

import (
	"context"
	"fmt"
	"log"
	gmail2 "route256-gmail-checker/internal/infrastructure/gmail"
	"route256-gmail-checker/pkg/config"
	"route256-gmail-checker/pkg/gmail"
)

func Run(cfg *config.Config) {
	ctx := context.Background()

	gmailService, err := gmail.NewGmailService(ctx, cfg.GoogleAPI)
	if err != nil {
		log.Fatal(err)
	}

	gmailClient := gmail2.NewClient(gmailService)

	messageList, err := gmailClient.GetLast10MessageIDs("")
	if err != nil {
		log.Fatal(err)
	}

	message, err := gmailClient.GetMessageByID(messageList[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", message)

	err = gmailClient.DeleteLabelByID(message.ID, "UNREAD")
	if err != nil {
		log.Fatal(err)
	}

	message, err = gmailClient.GetMessageByID(messageList[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", message)
}
