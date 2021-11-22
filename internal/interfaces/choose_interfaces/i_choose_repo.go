package choose_interfaces

import "go.mongodb.org/mongo-driver/mongo"

type IChooseRepo interface {
	GetChooseByPrice(*mongo.Database, string, interface{}) (interface{}, error)
}