package product

import (
	"Orbit/internal/models"
	"Orbit/internal/repositories"
	"Orbit/internal/utils"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func RegisterProductsService(ctx context.Context, sellerIdStr string, eventIdStr string, items []ProductItemRequest) error {
	sellerId, err := utils.GetObjectFiedIdFromString(sellerIdStr)
	if err != nil {
		return errors.New("invalid seller ID")
	}

	eventId, err := utils.GetObjectFiedIdFromString(eventIdStr)
	if err != nil {
		return errors.New("invalid event ID")
	}

	event, err := repositories.GetAnEvent(eventIdStr)
	if err != nil || event == nil {
		return errors.New("event not found")
	}

	if event.SellerID != sellerId {
		return errors.New("unauthorized access to this event")
	}

	now := time.Now()
	inventories := make([]models.Inventory, len(items))

	for i, item := range items {
		inventories[i] = models.Inventory{
			ID:          bson.NewObjectID(),
			SellerID:    sellerId,
			EventID:     eventId,
			Title:       item.Title,
			Description: item.Description,
			Price:       item.Price,
			Frequency:   item.Frequency,
			Image:       item.Image,
			CreatedAt:   now,
			UpdatedAt:   now,
		}
	}

	return repositories.BulkInsertInventory(ctx, inventories)
}

func GetAllEventProductsService(ctx context.Context, eventIdStr string) ([]*models.Inventory, error) {
	return repositories.GetAllProductsByEvent(ctx, eventIdStr)
}

func GetAnEventProductService(ctx context.Context, productIdStr string) (*models.Inventory, error) {
	product, err := repositories.GetProductByID(ctx, productIdStr)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}
	return product, nil
}

func UpdateAnEventProductService(ctx context.Context, sellerIdStr, eventIdStr, productIdStr string, req UpdateProductRequestBody) error {
	sellerId, err := utils.GetObjectFiedIdFromString(sellerIdStr)
	if err != nil {
		return errors.New("invalid seller ID")
	}

	event, err := repositories.GetAnEvent(eventIdStr)
	if err != nil || event == nil {
		return errors.New("event not found")
	}

	if event.SellerID != sellerId {
		return errors.New("unauthorized access to this event")
	}

	product, err := repositories.GetProductByID(ctx, productIdStr)
	if err != nil || product == nil {
		return errors.New("product not found")
	}

	updateData := bson.M{}
	if req.Title != nil {
		updateData["title"] = *req.Title
	}
	if req.Description != nil {
		updateData["description"] = *req.Description
	}
	if req.Price != nil {
		updateData["price"] = *req.Price
	}
	if req.Frequency != nil {
		updateData["frequency"] = *req.Frequency
	}
	if req.Image != nil {
		updateData["image"] = *req.Image
	}

	if len(updateData) == 0 {
		return errors.New("no fields provided for update")
	}

	return repositories.UpdateProduct(ctx, productIdStr, updateData)
}

func DeleteAnEventProductService(ctx context.Context, sellerIdStr, eventIdStr, productIdStr string) error {
	sellerId, err := utils.GetObjectFiedIdFromString(sellerIdStr)
	if err != nil {
		return errors.New("invalid seller ID")
	}

	event, err := repositories.GetAnEvent(eventIdStr)
	if err != nil || event == nil {
		return errors.New("event not found")
	}

	if event.SellerID != sellerId {
		return errors.New("unauthorized access to this event")
	}

	product, err := repositories.GetProductByID(ctx, productIdStr)
	if err != nil || product == nil {
		return errors.New("product not found")
	}

	return repositories.DeleteProduct(ctx, productIdStr)
}