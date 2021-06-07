package productmongorepo

import (
	"context"
	"errors"
	"log"

	"github.com/afif0808/gomongohelper"
	"github.com/afif0808/inventory-sales-service/model"
	productrepo "github.com/afif0808/inventory-sales-service/product/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetMany(db *mongo.Database, filterFuncs ...gomongohelper.FilterFunc) productrepo.GetManyFunc {
	return func(ctx context.Context, parameters map[string]interface{}) ([]model.Product, error) {
		ps := []model.Product{}

		filter := bson.M{}
		var err error
		if parameters != nil {
			for _, f := range filterFuncs {
				filter, err = f(filter, parameters)
				if err != nil {
					log.Println(err)
					return nil, err
				}
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

func Create(db *mongo.Database) productrepo.CreateFunc {
	return func(ctx context.Context, p model.Product) (model.Product, error) {
		res, err := db.Collection("product").InsertOne(ctx, p)
		if err != nil {
			log.Println(err)
			return p, err
		}
		p.ID = res.InsertedID.(primitive.ObjectID)

		return p, nil
	}
}

func DeleteMany(db *mongo.Database, filterFuncs ...gomongohelper.FilterFunc) productrepo.DeleteManyFunc {
	return func(ctx context.Context, parameters map[string]interface{}) error {
		filter := bson.M{}
		if parameters != nil {
			for _, f := range filterFuncs {
				var err error
				filter, err = f(filter, parameters)
				if err != nil {
					log.Println(err)
					return err
				}
			}
		}
		res, err := db.Collection("product").DeleteMany(ctx, filter)
		if err != nil {
			log.Println(err)
			return err
		}
		if res.DeletedCount < 1 {
			log.Println("Error : No products were successfuly deleted ")
			return errors.New("Error : No products were successfuly deleted")
		}
		return nil
	}
}
func UpdateMany(db *mongo.Database, filterFuncs ...gomongohelper.FilterFunc) productrepo.UpdateManyFunc {
	return func(ctx context.Context, p model.Product, parameters map[string]interface{}) error {
		filter := bson.M{}
		var err error
		if parameters != nil {
			for _, f := range filterFuncs {
				filter, err = f(filter, parameters)
				if err != nil {
					log.Println(err)
					return err
				}
			}
		}
		res, err := db.Collection("product").UpdateMany(ctx, filter, p)
		if err != nil {
			log.Println(err)
			return err
		}
		if res.ModifiedCount < 1 {
			log.Println("Error : No products were successfuly updated ")
			return errors.New("Error : No products were successfuly updated")
		}
		return nil
	}
}
