package oauthdebugger

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
)

// Client representation of a client (user)
type Client struct {
	ClientId     string `firestore:"client_id"`
	ClientSecret string `firestore:"client_secret"`
	Code         string `firestore:"code"`
	Name         string `firestore:"name"`
	RedirectUri  string `firestore:"redirect_uri"`
}

var ctx = context.Background()

// GetClient Retrieves a client structure from the database
func GetClient(clientId string) (Client, error) {
	var db *firestore.Client
	db, err := connect()
	if err != nil {
		return Client{}, err
	}

	clientRef := doc(db, clientId)
	docsnap, err := clientRef.Get(ctx)
	if err != nil {
		return Client{}, err
	}

	var c Client
	err = docsnap.DataTo(&c)
	c.ClientId = clientId
	return c, err
}

func doc(db *firestore.Client, clientId string) *firestore.DocumentRef {
	matcher := fmt.Sprintf("clients/%s", clientId)
	return db.Doc(matcher)
}

func connect() (*firestore.Client, error) {
	projectId := os.Getenv("OAD_PROJECT_ID")
	client, err := firestore.NewClient(ctx, projectId)
	return client, err
}

// Save puts record in db
func Save(c Client) error {
	var db *firestore.Client
	db, err := connect()
	if err != nil {
		return err
	}

	doc := doc(db, c.ClientId)
	_, err = doc.Create(ctx, c)
	if err != nil {
		return err
	}
	return nil
}
