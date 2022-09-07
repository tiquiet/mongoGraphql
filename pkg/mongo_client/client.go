package mongo_client

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func NewClient(ctx context.Context, host, port, username, password, database string) (*mongo.Database, error) {
	var mongoDBURL string
	var anonymous bool
	if username == "" || password == "" {
		anonymous = true
		mongoDBURL = fmt.Sprintf("mongodb://%s:%s", host, port)
	} else {
		mongoDBURL = fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)
	}
	reqCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	log.Print(mongoDBURL)
	clientOptions := options.Client().ApplyURI(mongoDBURL)
	if !anonymous {
		clientOptions.SetAuth(options.Credential{
			Username:    username,
			Password:    password,
			PasswordSet: true,
		})
	}
	
	client, err := mongo.Connect(reqCtx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to create client to mongodb due to error %w", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create client to mongodb due to error %w", err)
	}

	return client.Database(database), nil
}
