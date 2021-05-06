package oauthdebugger

import (
	"os"

	firestore "cloud.google.com/go/firestore"
)

type docAction func(*firestore.DocumentRef) error

func connect() (*firestore.Client, error) {
	projectId := os.Getenv("OAD_PROJECT_ID")
	client, err := firestore.NewClient(ctx, projectId)
	return client, err
}
