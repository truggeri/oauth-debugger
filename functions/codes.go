package oauthdebugger

import (
	"fmt"
	"time"

	firestore "cloud.google.com/go/firestore"
)

// Code representation of a code
type Code struct {
	Code     string    `firestore:"code"`
	ClientId string    `firestore:"client_id"`
	Expires  time.Time `firestore:"expires"`
	Username string    `firestore:"username"`
}

var emptyCode = Code{}

// getDbCode Retrieves a code structure from the database
func getDbCode(code string) (Code, error) {
	db, err := connect()
	if err != nil {
		return emptyCode, err
	}

	docRef := codeRef(db, code)
	docsnap, err := docRef.Get(ctx)
	if err != nil {
		return emptyCode, err
	}

	var c Code
	err = docsnap.DataTo(&c)
	c.Code = code
	return c, err
}

func codeRef(db *firestore.Client, code string) *firestore.DocumentRef {
	matcher := fmt.Sprintf("codes/%s", code)
	return db.Doc(matcher)
}

// Creates a new Code in the database
func createDbCode(c Code) error {
	createAction := func(doc *firestore.DocumentRef) error {
		_, err := doc.Create(ctx, c)
		return err
	}
	return withDbCode(c, createAction)
}

func withDbCode(c Code, action docAction) error {
	var db *firestore.Client
	db, err := connect()
	if err != nil {
		return err
	}

	doc := codeRef(db, c.Code)
	return action(doc)
}
