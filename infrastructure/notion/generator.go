package notion

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jomei/notionapi"
	"github.com/jyisus/notioncli/entity"
)

type Client struct {
	client *notionapi.Client
}

type responseData struct {
	Title []struct {
		PlainText string `json:"plain_text"`
	} `json:"title"`
}

func NewClient(cfg entity.Config) Client {
	client := notionapi.NewClient(notionapi.Token(cfg.NotionApiKey))
	return Client{
		client: client,
	}
}

func (c Client) AddTask(ctx context.Context, database entity.Database, task string) error {
	request := &notionapi.PageCreateRequest{
		Parent: notionapi.Parent{
			DatabaseID: notionapi.DatabaseID(database.Id),
		},
		Properties: notionapi.Properties{
			database.Key: notionapi.TitleProperty{
				Title: []notionapi.RichText{
					{Text: notionapi.Text{Content: task}},
				},
			},
		},
	}
	_, err := c.client.Page.Create(ctx, request)
	if err != nil {
		return err
	}

	return nil
}

func (c Client) DeleteTask(ctx context.Context, task string) error {
	_, err := c.client.Block.Delete(ctx, notionapi.BlockID(task))
	if err != nil {
		return err
	}
	fmt.Printf("Deleted block with ID %s\n", task)
	return nil
}

func (c Client) ListTasks(ctx context.Context, database entity.Database) ([]entity.Task, error) {
	var request notionapi.DatabaseQueryRequest
	if database.Filter != "" {
		var filter notionapi.PropertyFilter

		err := json.Unmarshal([]byte(database.Filter), &filter)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		request.PropertyFilter = &filter
	}

	res, err := c.client.Database.Query(ctx, notionapi.DatabaseID(database.Id), &request)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	var tasks []entity.Task
	dt := responseData{}
	for _, value := range res.Results {
		st, err := json.Marshal(value.Properties[database.Key])
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(st, &dt)
		if err != nil {
			return nil, err
		}
		taskId := string(value.ID)
		if len(dt.Title) == 0 {
			tasks = append(
				tasks,
				entity.Task{
					Text: "Untitled",
					Id:   taskId,
				})
		} else {
			tasks = append(
				tasks,
				entity.Task{
					Text: dt.Title[0].PlainText,
					Id:   taskId,
				})
		}
	}

	return tasks, nil
}
