package productresthandler

import (
	"github.com/afif0808/goechohelper"
	productrepo "github.com/afif0808/inventory-sales-service/product/repository"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
)

func InjectProductRESTHandlers(ee *echo.Echo, db *mongo.Database) {
	ee.GET("/api/products", GetMany(productmongorepo.GetMany(db)))
}

func GetMany(getMany productrepo.GetManyFunc) echo.HandlerFunc {
	return goechohelper.NewGetterHandler(getMany.Generic()).Handler()
}
