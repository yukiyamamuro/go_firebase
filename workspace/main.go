package main

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
)

func main() {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, os.Getenv("GCP_PROJECT_ID"))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer client.Close()

	doc, _, err := client.Collection("users").Add(ctx, map[string]interface{}{
		"name": "yukiyama",
		"age":  123,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(doc.ID)

	q := client.Collection("users").Select("name", "age")
	docs, err := q.Documents(ctx).GetAll()
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, doc := range docs {
		log.Println(doc.Data())
	}

}
