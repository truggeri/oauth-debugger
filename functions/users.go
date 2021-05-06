package oauthdebugger

import (
	"context"
	"errors"
	"fmt"
	"time"

	firestore "cloud.google.com/go/firestore"
	"github.com/google/uuid"
)

type User struct {
	ClientId     string    `firestore:"client_id"`
	RefreshToken string    `firestore:"refresh_token"`
	Token        string    `firestore:"token"`
	TokenExpires time.Time `firestore:"token_expires"`
	Username     string    `firestore:"username"`
	Uuid         string    `firestore:"uuid"`
}

var emptyUser = User{}

// userFromCode Given a db code, create a db user and remove the code
func userFromCode(code Code) (User, error) {
	user := User{
		ClientId:     code.ClientId,
		RefreshToken: RandomString(32),
		Token:        RandomString(32),
		TokenExpires: time.Now().Add(24 * time.Hour),
		Username:     code.Username,
		Uuid:         uuid.New().String(),
	}

	db, err := connect()
	if err != nil {
		return emptyUser, err
	}
	userRef := userDoc(db, user.Token)
	codeRef := codeRef(db, code.Code)
	err = db.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		err := tx.Create(userRef, user)
		if err != nil {
			return err
		}
		return tx.Delete(codeRef)
	})

	return user, err
}

func userDoc(db *firestore.Client, token string) *firestore.DocumentRef {
	matcher := fmt.Sprintf("users/%s", token)
	return db.Doc(matcher)
}

// getDbUser Gets a user from the Db by token
func getDbUser(token, clientId string) (User, error) {
	db, err := connect()
	if err != nil {
		return emptyUser, err
	}

	userRef := userDoc(db, token)
	docsnap, err := userRef.Get(ctx)
	if err != nil {
		return emptyUser, err
	}

	var u User
	err = docsnap.DataTo(&u)
	if clientId != "" && u.ClientId != clientId {
		return emptyUser, errors.New("User not associated to client")
	}
	return u, err
}
