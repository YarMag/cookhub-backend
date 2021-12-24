package gofirebase

import (
  "fmt"
  "context"

  firebase "firebase.google.com/go"
  "firebase.google.com/go/auth"

  "google.golang.org/api/option"
)


func SetupAuth() (*auth.Client, error) {
	opt := option.WithCredentialsFile("/tmp/serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
  		return nil, fmt.Errorf("error initializing app: %v", err)
	}
  auth, err := app.Auth(context.Background())
  if err != nil {
      return nil, fmt.Errorf("error initializing auth: %v", err)
  }
  return auth, nil
}