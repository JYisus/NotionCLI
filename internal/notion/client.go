package notion

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jomei/notionapi"
	"github.com/jyisus/notioncli/internal/config"
)

type Client struct {
	client     *notionapi.Client
	databaseId notionapi.DatabaseID
}

func NewClient(cfg config.Config) Client {
	client := notionapi.NewClient(notionapi.Token(cfg.NotionApiKey))
	return Client{
		client:     client,
		databaseId: notionapi.DatabaseID(cfg.DatabaseId),
	}
}

func (c Client) AddTask(task string) error {
	request := &notionapi.PageCreateRequest{
		Parent: notionapi.Parent{
			DatabaseID: c.databaseId,
		},
		Properties: notionapi.Properties{
			"Name": notionapi.TitleProperty{
				Title: []notionapi.RichText{
					{Text: notionapi.Text{Content: task}},
				},
			},
		},
	}
	_, err := c.client.Page.Create(context.TODO(), request) //.Update(context.TODO(), notionapi.DatabaseID(databaseId), request)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	return nil
}

func (c Client) ListTasks() ([]string, error) {
	request := &notionapi.DatabaseQueryRequest{}
	//res, err := client.Database.f(context.TODO(), databaseId)
	res, err := c.client.Database.Query(context.TODO(), c.databaseId, request)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	type Data struct {
		Title []struct {
			PlainText string `json:"plain_text"`
		} `json:"title"`
	}

	var tasks []string
	dt := Data{}
	for _, value := range res.Results {
		st, err := json.Marshal(value.Properties["Name"])
		if err != nil {
			fmt.Println(err)
		}
		json.Unmarshal(st, &dt)
		tasks = append(tasks, dt.Title[0].PlainText)
	}

	return tasks, nil
}
