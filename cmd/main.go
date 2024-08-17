package main

import (
	"context"
	"fmt"

	"gitlab.com/hartsfield/gmailAPI"
	"gitlab.com/hartsfield/inboxer"
	gmail "google.golang.org/api/gmail/v1"
)

func main() {
	// Connect to the gmail API service.
	ctx := context.Background()
	srv := gmailAPI.ConnectToService(ctx, gmail.MailGoogleComScope)

	msgs, err := inboxer.Query(srv, "category:forums after:2017/01/01 before:2017/01/30")
	if err != nil {
		fmt.Println(err)
	}

	// Range over the messages
	for _, msg := range msgs {
		fmt.Println("========================================================")
		time, err := inboxer.ReceivedTime(msg.InternalDate)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Date: ", time)
		md := inboxer.GetPartialMetadata(msg)
		fmt.Println("From: ", md.From)
		fmt.Println("Sender: ", md.Sender)
		fmt.Println("Subject: ", md.Subject)
		fmt.Println("Delivered To: ", md.DeliveredTo)
		fmt.Println("To: ", md.To)
		fmt.Println("CC: ", md.CC)
		fmt.Println("Mailing List: ", md.MailingList)
		fmt.Println("Thread-Topic: ", md.ThreadTopic)
		fmt.Println("Snippet: ", msg.Snippet)
		body, err := inboxer.GetBody(msg, "text/plain")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(body)
	}
}
