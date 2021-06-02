package productmongorepo

import (
	"context"
	"log"

	"github.com/afif0808/gomongohelper"
	"github.com/afif0808/inventory-and-sales-service/models"
	productrepo "github.com/afif0808/inventory-and-sales-service/product/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetMany(db *mongo.Database, filterFuncs ...gomongohelper.FilterFunc) productrepo.GetManyFunc {
	return func(ctx context.Context, parameters map[string]interface{}) ([]models.Product, error) {
		ps := []models.Product{}

		filter := bson.M{}
		if filterFUncs != nil && parameters != nil {
			for _, f := range filterFuncs {
				filter = f(filter, parameters)
			}
		}
		cursor, err := db.Collection("product").Find(ctx, filter)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		err = cursor.All(ctx, &ps)
		if err != nil {
			return nil, err
		}
		return ps, err
	}
}
