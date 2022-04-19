package notion

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jomei/notionapi"
	"github.com/jyisus/notioncli/internal/config"
)

type Client struct {
	client          *notionapi.Client
	defaultDatabase database
	databases       []database
}

type database struct {
	name    string
	id      notionapi.DatabaseID
	key     string
	request notionapi.DatabaseQueryRequest
}

type data struct {
	Title []struct {
		PlainText string `json:"plain_text"`
	} `json:"title"`
}

func NewClient(cfg config.Config) Client {
	client := notionapi.NewClient(notionapi.Token(cfg.NotionApiKey))
	var databases []database
	var defaultDatabase database
	for _, db := range cfg.Databases {
		var request notionapi.DatabaseQueryRequest

		if db.Filter != "" {
			var filter notionapi.PropertyFilter
			err := json.Unmarshal([]byte(db.Filter), &filter)
			if err != nil {
				fmt.Println(err)
				return Client{}
			}
			request.PropertyFilter = &filter
		}

		newDatabase := database{
			name:    db.Name,
			id:      notionapi.DatabaseID(db.Id),
			key:     db.Key,
			request: request,
		}
		databases = append(databases, newDatabase)
		if db.Name == cfg.DefaultDatabase {
			defaultDatabase = newDatabase
		}
	}

	//var databasesIds []notionapi.DatabaseID
	return Client{
		client:          client,
		databases:       databases,
		defaultDatabase: defaultDatabase,
	}
}

func (c Client) AddTask(databaseName, task string) error {
	var database database
	var err error
	if databaseName == "default" {
		database = c.defaultDatabase
	} else {
		database, err = c.getDatabaseByName(databaseName)
		if err != nil {
			return err
		}
	}
	request := &notionapi.PageCreateRequest{
		Parent: notionapi.Parent{
			DatabaseID: database.id,
		},
		Properties: notionapi.Properties{
			database.key: notionapi.TitleProperty{
				Title: []notionapi.RichText{
					{Text: notionapi.Text{Content: task}},
				},
			},
		},
	}
	_, err = c.client.Page.Create(context.TODO(), request) //.Update(context.TODO(), notionapi.DatabaseID(databaseId), request)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	return nil
}

func (c Client) ListTasks(databaseName string) ([]string, error) {
	var database database
	var err error
	if databaseName == "default" {
		database = c.defaultDatabase
	} else {
		database, err = c.getDatabaseByName(databaseName)
		if err != nil {
			return nil, err
		}
	}

	//request := &notionapi.DatabaseQueryRequest{
	//	PropertyFilter: &notionapi.PropertyFilter{
	//		Property: "Done",
	//		Checkbox: &notionapi.CheckboxFilterCondition{
	//			//Equals:       false,
	//			DoesNotEqual: true,
	//		},
	//	},
	//}

	res, err := c.client.Database.Query(context.TODO(), database.id, &database.request)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	var tasks []string
	dt := data{}
	for _, value := range res.Results {
		st, err := json.Marshal(value.Properties[database.key])
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(st, &dt)
		if err != nil {
			return nil, err
		}
		if len(dt.Title) == 0 {
			tasks = append(tasks, "Untitled")
		} else {
			tasks = append(tasks, dt.Title[0].PlainText)
		}
	}

	return tasks, nil
}

func (c Client) getDatabaseByName(databaseName string) (database, error) {
	for _, db := range c.databases {
		if db.name == databaseName {
			return db, nil
		}
	}
	return database{}, errors.New("unable to find database in configuration file")
}
