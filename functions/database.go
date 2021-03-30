package oauthdebugger

import (
	"context"
	"fmt"
	"os"
	"time"

	firestore "cloud.google.com/go/firestore"
)

// Client representation of a client (user)
type Client struct {
	ClientId     string     `firestore:"client_id"`
	ClientSecret string     `firestore:"client_secret"`
	Expires      time.Time  `firestore:"expires"`
	Name         string     `firestore:"name"`
	RedirectUri  string     `firestore:"redirect_uri"`
	Users        []AuthUser `firestore:"users"`
}

type AuthUser struct {
	Code         string    `firestore:"code"`
	RefreshToken string    `firestore:"refresh_token"`
	Token        string    `firestore:"token"`
	TokenExpires time.Time `firestore:"token_expires"`
	Username     string    `firestore:"username"`
	Uuid         string    `firestore:"uuid"`
}

func (c Client) empty() bool {
	return c.ClientId == ""
}

var emptyClient = Client{}
var ctx = context.Background()

type docAction func(*firestore.DocumentRef) error

// getDbClient Retrieves a client structure from the database
func getDbClient(clientId string) (Client, error) {
	var db *firestore.Client
	db, err := connect()
	if err != nil {
		return emptyClient, err
	}

	clientRef := doc(db, clientId)
	docsnap, err := clientRef.Get(ctx)
	if err != nil {
		return emptyClient, err
	}

	var c Client
	err = docsnap.DataTo(&c)
	c.ClientId = clientId
	return c, err
}

func connect() (*firestore.Client, error) {
	projectId := os.Getenv("OAD_PROJECT_ID")
	client, err := firestore.NewClient(ctx, projectId)
	return client, err
}

func doc(db *firestore.Client, clientId string) *firestore.DocumentRef {
	matcher := fmt.Sprintf("clients/%s", clientId)
	return db.Doc(matcher)
}

// Creates a new Client in the database
func createDbClient(c Client) error {
	createAction := func(doc *firestore.DocumentRef) error {
		_, err := doc.Create(ctx, c)
		return err
	}
	return withDbDoc(c, createAction)
}

func withDbDoc(c Client, action docAction) error {
	var db *firestore.Client
	db, err := connect()
	if err != nil {
		return err
	}

	doc := doc(db, c.ClientId)
	return action(doc)
}

// Updates an existing Client in the database
func _updateDbClient(c Client, updates []firestore.Update) error {
	updateAction := func(doc *firestore.DocumentRef) error {
		_, err := doc.Update(ctx, updates)
		return err
	}
	return withDbDoc(c, updateAction)
}

// Sets entire existing client to the database
func _setDbClient(c Client) error {
	saveAction := func(doc *firestore.DocumentRef) error {
		_, err := doc.Set(ctx, c)
		return err
	}
	return withDbDoc(c, saveAction)
}

// Merges an AuthUser to an array
func mergeDbUser(c Client, au AuthUser) error {
	updateObj := map[string]interface{}{
		"users": []AuthUser{au},
	}
	mergeAction := func(doc *firestore.DocumentRef) error {
		_, err := doc.Set(ctx, updateObj, firestore.MergeAll)
		return err
	}
	return withDbDoc(c, mergeAction)
}
