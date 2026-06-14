package repositories

import (
	"Orbit/internal/db"
	"Orbit/internal/models"
	"Orbit/internal/utils"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func CreateAnEvent(event *models.Event) error {
	instance := db.GetInstance().Collection("events")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := instance.InsertOne(ctx, event)
	if err != nil {
		return err
	}
	return nil
}

func GetAllEvents(sellerId string) ([]*models.Event, error) {
	ObjectifiedId, err := utils.GetObjectFiedIdFromString(sellerId)
	if err != nil {
		return nil, err
	}

	instance := db.GetInstance().Collection("events")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := instance.Find(ctx, bson.M{"sellerId": ObjectifiedId})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	events := []*models.Event{}
	for cursor.Next(ctx) {
		var event models.Event
		if err := cursor.Decode(&event); err != nil {
			return nil, err
		}
		events = append(events, &event)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return events, nil
}

func GetAnEvent(eventId string) (*models.Event, error) {
	ObjectifiedId, err := utils.GetObjectFiedIdFromString(eventId)
	if err != nil {
		return nil, err
	}

	instance := db.GetInstance().Collection("events")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var event models.Event
	err = instance.FindOne(ctx, bson.M{"_id": ObjectifiedId}).Decode(&event)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func UpdateAnEvent(eventId string, updateData bson.M) error {
	ObjectifiedId, err := utils.GetObjectFiedIdFromString(eventId)
	if err != nil {
		return err
	}

	instance := db.GetInstance().Collection("events")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	updateData["updatedAt"] = time.Now()

	_, err = instance.UpdateOne(ctx, bson.M{"_id": ObjectifiedId}, bson.M{"$set": updateData})
	if err != nil {
		return err
	}

	return nil
}

func DeleteAnEvent(eventId string) error {
	ObjectifiedId, err := utils.GetObjectFiedIdFromString(eventId)
	if err != nil {
		return err
	}

	instance := db.GetInstance().Collection("events")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = instance.DeleteOne(ctx, bson.M{"_id": ObjectifiedId})
	if err != nil {
		return err
	}

	return nil
}