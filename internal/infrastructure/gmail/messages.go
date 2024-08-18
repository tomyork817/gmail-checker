package gmail

import (
	"errors"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/googleapi"
	"route256-gmail-checker/internal/domain"
	"strconv"
)

var (
	NotPositiveCountErr = errors.New("count can't be <= 0")
)

func (c *Client) GetLastNMessageIDs(searchQuery string, count int) ([]string, error) {
	if count <= 0 {
		return nil, NotPositiveCountErr
	}

	messageListResponse, err := c.service.Users.Messages.List(c.user).Do(
		googleapi.QueryParameter("q", searchQuery),
		googleapi.QueryParameter("maxResults", strconv.Itoa(count)),
	)
	if err != nil {
		return nil, err
	}

	ids := make([]string, 0, count)
	for _, message := range messageListResponse.Messages {
		ids = append(ids, message.Id)
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
