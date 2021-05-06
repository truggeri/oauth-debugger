package oauthdebugger

import (
	"context"
	"fmt"
	"time"

	firestore "cloud.google.com/go/firestore"
)

// Client representation of a client
type Client struct {
	ClientId     string    `firestore:"client_id"`
	ClientSecret string    `firestore:"client_secret"`
	Expires      time.Time `firestore:"expires"`
	Name         string    `firestore:"name"`
	RedirectUri  string    `firestore:"redirect_uri"`
}

func (c Client) empty() bool {
	return c.ClientId == ""
}

var emptyClient = Client{}
var ctx = context.Background()

// getDbClient Retrieves a client structure from the database
func getDbClient(clientId string) (Client, error) {
	db, err := connect()
	if err != nil {
		return emptyClient, err
	}

	clientRef := clientDoc(db, clientId)
	docsnap, err := clientRef.Get(ctx)
	if err != nil {
		return emptyClient, err
	}

	var c Client
	err = docsnap.DataTo(&c)
	c.ClientId = clientId
	return c, err
}

func clientDoc(db *firestore.Client, clientId string) *firestore.DocumentRef {
	matcher := fmt.Sprintf("clients/%s", clientId)
	return db.Doc(matcher)
}

// Creates a new Client in the database
func createDbClient(c Client) error {
	createAction := func(doc *firestore.DocumentRef) error {
		_, err := doc.Create(ctx, c)
		return err
	}
	return withDbClient(c, createAction)
}

func withDbClient(c Client, action docAction) error {
	var db *firestore.Client
	db, err := connect()
	if err != nil {
		return err
	}

	doc := clientDoc(db, c.ClientId)
	return action(doc)
}
