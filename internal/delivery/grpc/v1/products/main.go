package purchase

import (
	"context"

	"github.com/bhankey/pharmacy-automatization-purchase/internal/entities"
	"github.com/bhankey/pharmacy-automatization-purchase/pkg/api/purchaseproto"
)

type GRPCHandler struct {
	purchaseproto.UnimplementedPurchaseServiceServer // Must be

	purchaseSrv Srv
}

type Srv interface {
	AddToPurchase(ctx context.Context, productName string, position string, purchaseUUID string) error
	ConfirmPurchase(ctx context.Context, userID, pharmacyID int, purchaseUUID string, isSocialCardUsed bool) error
	GetPurchase(ctx context.Context, pharmacyID int, purchaseUUID string, isSocialCard bool) (entities.Purchase, error)
	GetAvailableToPurchaseProducts(ctx context.Context, pharmacyID int) ([]entities.PharmacyProductItem, error)
	DeleteFromPurchase(ctx context.Context, productName string, position string, purchaseUUID string) error
}

func NewPurchaseGRPCHandler(srv Srv) *GRPCHandler {
	return &GRPCHandler{
		UnimplementedPurchaseServiceServer: purchaseproto.UnimplementedPurchaseServiceServer{},
		purchaseSrv:                        srv,
	}
}
