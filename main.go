package main

import (
	"context"
	"fmt"
	"log"
	"time"

	productmongorepo "github.com/afif0808/inventory-and-sales-service/product/repository/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	db := client.Database("afentory")
	ps, err := productmongorepo.GetMany(db)(context.Background(), nil)
	fmt.Println(ps, err)
}
