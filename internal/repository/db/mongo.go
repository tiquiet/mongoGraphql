package db

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/tiquiet/mongoGraphql/internal/custom_models"
	"github.com/tiquiet/mongoGraphql/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type db struct {
	collection *mongo.Collection
}

func NewStorage(storage *mongo.Database, collection string) repository.Repository {
	return &db{
		collection: storage.Collection(collection),
	}
}

func (d *db) CreateBook(ctx context.Context, input custom_models.NewBook) (string, error) {
	input.LastUpdate = time.Now()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	result, err := d.collection.InsertOne(ctx, input)
	if err != nil {
		return "", fmt.Errorf("failed to create book error:%v", err)
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("failed to convet objectid to hex")
	}

	return oid.Hex(), nil
}

func (d *db) GetAllBooks(ctx context.Context) ([]*custom_models.Book, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	cursor, err := d.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to get all items from book colection. error: %w", err)
	}

	var results []*custom_models.Book
	if err = cursor.All(ctx, &results); err != nil {
		return nil, fmt.Errorf("failed to convert books data from db to []*custom_models.Book. error: %w", err)
	}

	return results, nil
}

func (d *db) GetBook(ctx context.Context, id string) (*custom_models.Book, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("failed to convert hex to objectid. error: %w", err)
	}

	filter := bson.M{"_id": objectID}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	result := d.collection.FindOne(ctx, filter)
	if err = result.Err(); err != nil {
		return nil, fmt.Errorf("failed to find book from book collection with filter:\n%s\nerror: %w", filter, err)
	}

	book := new(custom_models.Book)
	if err = result.Decode(book); err != nil {
		return nil, fmt.Errorf("failed to decode result to custom_models.Book. error: %w", err)
	}

	return book, nil
}

func (d *db) UpdateBook(ctx context.Context, id string, input custom_models.UpdateBook) (bool, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, fmt.Errorf("failed to convert hex to objectid. error: %w", err)
	}

	input.LastUpdate = time.Now()

	bookByte, err := bson.Marshal(input)
	if err != nil {
		return false, fmt.Errorf("failed to marshal input to []byte. error: %w", err)
	}

	var updateObj bson.M
	err = bson.Unmarshal(bookByte, &updateObj)
	if err != nil {
		return false, fmt.Errorf("failed to unmarshal bookByte []byte to bson.M. error: %w", err)
	}

	update := bson.M{
		"$set": updateObj,
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	result, err := d.collection.UpdateByID(ctx, objectID, update)
	if err != nil {
		return false, fmt.Errorf("failed to update book error:%v", err)
	}

	if result.MatchedCount == 0 {
		return false, errors.New("result.MatchedCount: nothing to update")
	}
	if result.UpsertedCount != 0 {
		return false, fmt.Errorf("inserted a new document with ID %v\n", result.UpsertedID)
	}

	return true, nil
}

func (d *db) DeleteBook(ctx context.Context, id string) (bool, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, fmt.Errorf("failed to convert hex to objectid. error: %w", err)
	}

	filter := bson.M{"_id": objectID}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	result, err := d.collection.DeleteOne(ctx, filter)
	if err != nil {
		return false, fmt.Errorf("failed to delete object. error: %w", err)
	}
	if result.DeletedCount == 0 {
		return false, fmt.Errorf("delet zero items. error: %w", err)
	}

	return true, nil
}
