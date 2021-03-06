package infrastructure

import (
	"context"
	"github.com/pkg/errors"
	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
	"stonks/internal/config"
	"stonks/internal/interfaces/db_interfaces"
)

type MongoClient struct {
	Client *mongo.Client
}

func InitDbClient(logger *zap.SugaredLogger, cfg *config.Config, ctx context.Context) (db_interfaces.IDBHandler, error) {
	clientOptions := options.Client().ApplyURI(cfg.DbURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		logger.Error(err.Error())
		return &MongoClient{}, errors.Wrap(err, "mongo initialization err")
	}
	logger.Info("db client init ok")

	return &MongoClient{Client: client}, nil
}

func (c *MongoClient) AcquireDatabase(name string) *mongo.Database {
	return c.Client.Database(name)
}

func (c *MongoClient) AcquireClient() *mongo.Client {
	return c.Client
}
