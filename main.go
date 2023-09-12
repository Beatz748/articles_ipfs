package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Article struct {
	ID      string `json:"_id" bson:"_id"`
	Content string `json:"content"`
}

var (
	mongoURI      = os.Getenv("MONGO_URI")         // MongoDB connection URI
	mongoDBName   = os.Getenv("MONGO_DB_NAME")     // MongoDB database name
	mongoCollName = os.Getenv("MONGO_COLL_NAME")   // MongoDB collection name
	port          = os.Getenv("HTTP_SERVER_PORT")  // HTTP server port
	apiKey        = os.Getenv("PINATA_API_KEY")    // API key
	apiSecret     = os.Getenv("PINATA_API_SECRET") // API secret
)

func main() {
	// Create a MongoDB client
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		fmt.Println("Error while creating a MongoDB client: ", err)
	}
	defer mongoClient.Disconnect(ctx)

	// Get the articles collection from MongoDB
	articlesColl := mongoClient.Database(mongoDBName).Collection(mongoCollName)

	// Create a Gorilla Mux router
	r := mux.NewRouter()

	r.HandleFunc("/paper/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var article Article
		if id == "" {
			http.Error(w, "Article ID is not specified", http.StatusBadRequest)
			log.Println("Article ID is not specified")
			return
		}

		log.Printf("Received request for ID: %s\n", id)

		// Check if the article exists in MongoDB
		if err := getArticleFromMongoDB(id, articlesColl, &article); err != nil {
			http.Error(w, "Article not found in MongoDB", http.StatusNotFound)
			log.Printf("Article with ID %s not found in MongoDB: %v\n", id, err)
			return
		}

		// Log that the article exists
		log.Printf("Article with ID %s exists in MongoDB\n", id)

		// Upload the article content to IPFS
		cid, err := uploadToIPFS(apiKey, apiSecret, article.Content)
		if err != nil {
			http.Error(w, "Error uploading the article to IPFS", http.StatusInternalServerError)
			log.Printf("Error uploading the article with ID %s to IPFS: %v\n", id, err)
			return
		}

		// Return the CID
		fmt.Fprintf(w, "Article CID in IPFS: %s", cid)
	})

	// Create an HTTP server
	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	r.HandleFunc("/add-paper", func(w http.ResponseWriter, r *http.Request) {
		// JSON-request decode
		var article Article
		if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			log.Printf("Failed to parse request body: %v\n", err)
			return
		}

		// Adding to MongoDB
		if err := addArticleToMongoDB(articlesColl, &article); err != nil {
			http.Error(w, "Failed to add article to MongoDB", http.StatusInternalServerError)
			log.Printf("Failed to add article to MongoDB: %v\n", err)
			return
		}

		// Returning success
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Article added to MongoDB with ID: %s", article.ID)
		log.Printf("Article added to MongoDB with ID: %s", article.ID)
	})

	fmt.Printf("Server is listening on port %s...\n", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Error starting the HTTP server:", err)
	}
}
