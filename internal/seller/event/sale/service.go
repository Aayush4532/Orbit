package sale

import (
	"Orbit/internal/repositories"
	"Orbit/internal/utils"
)

func LiveSaleService(sellerId string, eventId string) error {
	sellerObjectifiedId, err := utils.GetObjectFiedIdFromString(sellerId);
	if err != nil {
		return err;
	}
	eventObjectifiedId, err := utils.GetObjectFiedIdFromString(eventId);
	if err != nil {
		return  err;
	}

	if err := repositories.PullProducts(sellerObjectifiedId, eventObjectifiedId); err != nil {
		return err;
	}
	return  nil;
}