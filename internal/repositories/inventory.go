package repositories

import (
	"Orbit/internal/db"
	"Orbit/internal/models"
	"Orbit/internal/utils"
	"Orbit/internal/worker"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func BulkInsertInventory(ctx context.Context, inventories []models.Inventory) error {
	if len(inventories) == 0 {
		return nil
	}

	instance := db.GetInstance().Collection("inventories")
	ctx, cancel := context.WithTimeout(ctx, 7*time.Second)
	defer cancel()

	docs := make([]interface{}, len(inventories))
	for i, inv := range inventories {
		docs[i] = inv
	}

	_, err := instance.InsertMany(ctx, docs)
	if err != nil {
		return err
	}

	return nil
}

func GetAllProductsByEvent(ctx context.Context, eventIdStr string) ([]*models.Inventory, error) {
	eventId, err := utils.GetObjectFiedIdFromString(eventIdStr)
	if err != nil {
		return nil, err
	}

	instance := db.GetInstance().Collection("inventories")
	tCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cursor, err := instance.Find(tCtx, bson.M{"eventId": eventId})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(tCtx)

	var products []*models.Inventory
	for cursor.Next(tCtx) {
		var item models.Inventory
		if err := cursor.Decode(&item); err != nil {
			return nil, err
		}
		products = append(products, &item)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func GetProductByID(ctx context.Context, productIdStr string) (*models.Inventory, error) {
	productId, err := utils.GetObjectFiedIdFromString(productIdStr)
	if err != nil {
		return nil, err
	}

	instance := db.GetInstance().Collection("inventories")
	tCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var item models.Inventory
	err = instance.FindOne(tCtx, bson.M{"_id": productId}).Decode(&item)
	if err != nil {
		return nil, nil
	}

	return &item, nil
}

func UpdateProduct(ctx context.Context, productIdStr string, updateData bson.M) error {
	productId, err := utils.GetObjectFiedIdFromString(productIdStr)
	if err != nil {
		return err
	}

	instance := db.GetInstance().Collection("inventories")
	tCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	updateData["updatedAt"] = time.Now()

	_, err = instance.UpdateOne(tCtx, bson.M{"_id": productId}, bson.M{"$set": updateData})
	if err != nil {
		return err
	}

	return nil
}

func DeleteProduct(ctx context.Context, productIdStr string) error {
	productId, err := utils.GetObjectFiedIdFromString(productIdStr)
	if err != nil {
		return err
	}

	instance := db.GetInstance().Collection("inventories")
	tCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err = instance.DeleteOne(tCtx, bson.M{"_id": productId})
	if err != nil {
		return err
	}

	return nil
}

func PullProducts(sellerId bson.ObjectID, eventId bson.ObjectID) error {
	pool := worker.InitWorkerPool()

	collection := db.GetInstance().Collection("inventories")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	findOptions := options.Find().
		SetBatchSize(200).
		SetProjection(bson.M{
			"image":       0,
			"createdAt":   0,
			"updatedAt":   0,
			"title":       0,
			"description": 0,
			"sellerId":    0,
			"eventId":     0,
		})

	cursor, err := collection.Find(ctx,
		bson.M{"sellerId": sellerId, "eventId": eventId},
		findOptions,
	)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product worker.ProductPayload
		if err := cursor.Decode(&product); err != nil {
			log.Printf("decode error: %v", err)
			continue
		}
		pool.Send(product)
	}

	defer pool.Close()

	if err := pool.Wait(); err != nil {
		return fmt.Errorf("redis flush error: %w", err)
	}

	return cursor.Err()
}
