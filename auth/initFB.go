package auth

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)
  
var App *firebase.App
var Err error
  func InitFireBase() {
	opt := option.WithCredentialsFile("/Users/rinishsoekhlall/Documents/golang-crud/tmy-dev-firebase-adminsdk-n7566-bdf34e53c8.json")
	App, Err = firebase.NewApp(context.Background(), nil, opt)
	log.Println("success")
	if Err != nil {
		log.Printf("error initializing app: %v", Err)
	  }
  } 