package details_interface

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type ICompanyDetailsRepo interface {
	GetCompanyDetails(*http.Request) (interface{}, error)
	GetDbCompanyDetails(string, string, *mongo.Database, bson.D) (interface{}, error)
	InsertCompanyDetails(string, *mongo.Database, interface{}) (interface{}, error)
}
