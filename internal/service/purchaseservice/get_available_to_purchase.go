package purchaseservice

import (
	"context"
	"fmt"

	"github.com/bhankey/pharmacy-automatization-purchase/internal/entities"
)

func (s *Service) GetAvailableToPurchaseProducts(ctx context.Context, pharmacyID int) ([]entities.PharmacyProductItem, error) {
	errBase := fmt.Sprintf("pharmacyservice.GetPharmacyProducts(%d)", pharmacyID)

	products, err := s.productRepo.GetAvailablePharmacyProducts(ctx, pharmacyID)
	if err != nil {
		return nil, fmt.Errorf("%s: Failed to get available products: %w", errBase, err)
	}

	return products, nil
}
