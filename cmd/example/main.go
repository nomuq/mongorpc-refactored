package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/mongorpc/mongorpc"
	"github.com/mongorpc/mongorpc/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	// gRPC server address
	address = "localhost:50051"
)

// Example MongoRPC Client
type ExampleClient struct {
	ctx      context.Context
	mongorpc proto.MongoRPCClient
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logrus.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a new client
	c := proto.NewMongoRPCClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	// Create a new client
	e := ExampleClient{
		ctx:      ctx,
		mongorpc: c,
	}

	e.ListCollections()
	e.ListDocuments()
	e.DocumentByID()
	e.CreateDocument()

	// update document
	// updateDocument(movie, data, err, result, c, ctx, insertResp)
}

// list all collections
func (c *ExampleClient) ListCollections() {
	r, err := c.mongorpc.ListCollections(c.ctx, &proto.ListCollectionsRequest{
		Database: "sample_mflix",
	})
	if err != nil {
		logrus.Fatalf("could not get collection: %v", err)
	}
	logrus.Printf("Collection: %s", r.Collections)
}

// get document
func (c *ExampleClient) DocumentByID() {
	doc, err := c.mongorpc.GetDocument(c.ctx, &proto.GetDocumentRequest{
		Database:   "sample_mflix",
		Collection: "movies",
		DocumentId: "573a1390f29313caabcd4135",
	})
	if err != nil {
		logrus.Fatalf("could not get document: %v", err)
	}
	logrus.Printf("Document: %s", doc.Document)
}

// list all documents
func (c *ExampleClient) ListDocuments() {
	documents, err := c.mongorpc.ListDocuments(c.ctx, &proto.ListDocumentsRequest{
		Database:   "sample_mflix",
		Collection: "movies",
		Limit:      1,
	}, grpc.MaxCallRecvMsgSize(1024*1024*1024))
	if err != nil {
		logrus.Fatalf("could not get documents: %v", err)
	}
	logrus.Printf("Documents: %s", documents.Documents)
}

// create document
func (c *ExampleClient) CreateDocument() {

	// Movie document
	movie := Movie{
		Title: "The Shawshank Redemption",
		Year:  "1994",
	}

	// Encode the movie to JSON
	var result map[string]interface{}
	data, err := json.Marshal(movie)
	if err != nil {
		logrus.Fatalf("could not marshal movie: %v", err)
	}

	// Decode the JSON to a map
	err = json.Unmarshal(data, &result)
	if err != nil {
		logrus.Fatalf("could not unmarshal movie: %v", err)
	}

	// create movie document
	insertResp, err := c.mongorpc.CreateDocument(c.ctx, &proto.CreateDocumentRequest{
		Database:   "sample_mflix",
		Collection: "sample_movies",
		Document: &proto.Value{
			Type: &proto.Value_MapValue{
				MapValue: mongorpc.EncodeMap(result),
			},
		},
	})
	if err != nil {
		logrus.Fatalf("could not create document: %v", err)
	}
	logrus.Printf("Document: %s", insertResp.DocumentId)
}

// Movie struct
type Movie struct {
	// Title of the movie
	Title string `json:"title"`
	// Year of the movie
	Year string `json:"year"`
}