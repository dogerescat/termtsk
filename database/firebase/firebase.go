package firebase

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type Config struct {
	KeyPath string
}
type FireStore struct {
	client *firestore.Client
}

func NewFireStore(config Config) *FireStore {
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
	return &FireStore{
		client: client,
	}
}
