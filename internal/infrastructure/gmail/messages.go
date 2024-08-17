package gmail

import (
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/googleapi"
	"route256-gmail-checker/internal/domain"
)

func (c *Client) GetLast10MessageIDs(searchQuery string) ([domain.MessagesListLen]string, error) {
	messageListResponse, err := c.service.Users.Messages.List(c.user).Do(googleapi.QueryParameter("q", searchQuery))
	if err != nil {
		return [domain.MessagesListLen]string{}, err
	}

	ids := [domain.MessagesListLen]string{}
	for i := 0; i < domain.MessagesListLen; i++ {
		ids[i] = messageListResponse.Messages[i].Id
	}

	return ids, nil
}

func (c *Client) GetMessageByID(id string) (domain.Message, error) {
	messageGetResponse, err := c.service.Users.Messages.Get(c.user, id).Do()
	if err != nil {
		return domain.Message{}, err
	}

	message := domain.Message{
		ID:      messageGetResponse.Id,
		Snippet: messageGetResponse.Snippet,
		Status:  domain.StatusRead,
	}

	for _, label := range messageGetResponse.LabelIds {
		if label == "UNREAD" {
			message.Status = domain.StatusUnread
		}
	}

	for _, header := range messageGetResponse.Payload.Headers {
		switch header.Name {
		case "From":
			message.From = header.Value
		case "Subject":
			message.Subject = header.Value
		}
	}

	return message, nil
}

func (c *Client) DeleteLabelByID(messageID, labelID string) error {
	_, err := c.service.Users.Messages.Modify(c.user, messageID, &gmail.ModifyMessageRequest{
		RemoveLabelIds: []string{labelID},
	}).Do()
	if err != nil {
		return err
	}
	return nil
}
