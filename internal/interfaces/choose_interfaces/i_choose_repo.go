package choose_interfaces

import "go.mongodb.org/mongo-driver/mongo"

type IChooseRepo interface {
	ChooseByPrice(*mongo.Database, string, interface{}) (interface{}, error)
	ChooseByPE(*mongo.Database, interface{}) (interface{}, error)
}
