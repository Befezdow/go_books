package core

import (
	"context"
	"fmt"
	"github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

import (
	"github.com/befezdow/go-books-rest-api/internal/shared"
)

var books = []shared.Book{
	{Id: uuid.NewV4().String(), Title: "Война и Мир", Author: &shared.Author{Firstname: "Лев", Lastname: "Толстой"}},
	{Id: uuid.NewV4().String(), Title: "Преступление и наказание", Author: &shared.Author{Firstname: "Фёдор", Lastname: "Достоевский"}},
}

var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
var client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://Befezdow:kjifhf33893@ds125945.mlab.com:25945"))

func getCollection(collectionName string) *mongo.Collection {
	if err != nil {
		log.Fatal(err)
	}
	return client.Database("library").Collection(collectionName)
}

func test() {
	var cur, err = getCollection("books").Find(ctx, bson.D{})
	defer cur.Close(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}

func AddBook(book shared.Book) shared.Book {
	book.Id = uuid.NewV4().String()
	books = append(books, book)
	return book
}

func GetBooks() []shared.Book {
	return books
}
