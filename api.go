package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
)

func main() {
	//new Firestore client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "project-resume-24")
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}

	// API HANDLER
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//allow cors
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// get current value
		doc, err := client.Collection("counters").Doc("counter").Get(ctx)
		if err != nil {
			log.Fatalf("Failed to get document: %v", err)
		}

		value := doc.Data()["value"].(int64)

		// incrementing value by 1
		value++

		// update value ==> new value
		_, err = client.Collection("counters").Doc("counter").Set(ctx, map[string]interface{}{
			"value": value,
		}, firestore.MergeAll)
		if err != nil {
			log.Fatalf("Failed to update document: %v", err)
		}
		//writing new value as JSON to response
		fmt.Fprintf(w, `{"value": %d}`, value)
	})

	//API SERVER ON
	port := "8080"
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
