package firebase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"termtsk/ui/form"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type Config struct {
	KeyPath string
}
type Firestore struct {
	client *firestore.Client
}

func NewFirestore(config Config) *Firestore {
	firebaseConfig := &firebase.Config{}
	ctx := context.Background()
	opt := option.WithCredentialsFile(config.KeyPath)
	app, err := firebase.NewApp(ctx, firebaseConfig, opt)
	if err != nil {
		log.Fatalln(err.Error())
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return &Firestore{
		client: client,
	}
}

func (f *Firestore) Create(task *form.Task) error {
	doc := f.client.Collection("Task").Doc(task.ID)
	_, err := doc.Set(context.Background(), map[string]interface{}{
		"id":         task.ID,
		"title":      task.Title,
		"detail":     task.Detail,
		"importance": task.Importance,
		"Done":       task.Done,
		"CreatedAt":  time.Now(),
	}, firestore.MergeAll)
	if err != nil {
		return err
	}
	return nil
}

func (f *Firestore) GetAll() []*form.Task {
	ctx := context.Background()
	iter := f.client.Collection("Task").Documents(ctx)
	var taskList []*form.Task
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		jsonData, err := json.Marshal(doc.Data())
		if err != nil {
			fmt.Println(err)
		}
		var task *form.Task
		if err := json.Unmarshal(jsonData, &task); err != nil {
			fmt.Println(err)
		}
		taskList = append(taskList, task)
	}
	return taskList
}

func (f *Firestore) Update(task *form.Task) error {
	doc := f.client.Collection("Task").Doc(task.ID)
	_, err := doc.Set(context.Background(), map[string]interface{}{
		"id":         task.ID,
		"title":      task.Title,
		"detail":     task.Detail,
		"importance": task.Importance,
		"Done":       task.Done,
		"CreatedAt":  task.CreatedAt,
	}, firestore.MergeAll)
	if err != nil {
		return err
	}
	return nil
}
