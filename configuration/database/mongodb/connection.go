package mongodb

import (
	"context"
	"fullcycle-auction_go/configuration"
	"fullcycle-auction_go/configuration/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDBConnection(ctx context.Context, cfg *configuration.Conf) (*mongo.Database, error) {
	client, err := mongo.Connect(
		ctx, options.Client().ApplyURI(cfg.DatabaseURL))
	if err != nil {
		logger.Error("Error trying to connect to mongodb database", err)
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		logger.Error("Error trying to ping mongodb database", err)
		return nil, err
	}

	return client.Database(cfg.DatabaseName), nil
}
