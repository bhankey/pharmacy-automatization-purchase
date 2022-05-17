package purchase

import (
	"context"
	"fmt"

	"github.com/bhankey/go-utils/pkg/apperror"
	"github.com/bhankey/pharmacy-automatization-purchase/pkg/api/purchaseproto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *GRPCHandler) ReserveItem(ctx context.Context, req *purchaseproto.ReserveItemRequest) (*emptypb.Empty, error) {
	errBase := fmt.Sprintf("products.ReserveItem(%v)", req)

	if err := req.ValidateAll(); err != nil {
		return nil, apperror.NewClientError(apperror.WrongRequest, err)
	}

	if err := h.purchaseSrv.AddToPurchase(ctx, req.ItemName, req.Position, req.PurchaseUuid); err != nil {
		return nil, fmt.Errorf("%s: failed to add to purchase: %w", errBase, err)
	}

	return &emptypb.Empty{}, nil
}
func (h *GRPCHandler) CancelReserveItem(ctx context.Context, req *purchaseproto.ReserveItemRequest) (*emptypb.Empty, error) {
	errBase := fmt.Sprintf("products.CancelReserveItem(%v)", req)

	if err := req.ValidateAll(); err != nil {
		return nil, apperror.NewClientError(apperror.WrongRequest, err)
	}

	if err := h.purchaseSrv.DeleteFromPurchase(ctx, req.ItemName, req.Position, req.PurchaseUuid); err != nil {
		return nil, fmt.Errorf("%s: failed to add to purchase: %w", errBase, err)
	}

	return &emptypb.Empty{}, nil
}
func (h *GRPCHandler) GetReservation(ctx context.Context, req *purchaseproto.GetReservationRequest) (*purchaseproto.Reservation, error) {
	errBase := fmt.Sprintf("products.GetReservation(%v)", req)

	if err := req.ValidateAll(); err != nil {
		return nil, apperror.NewClientError(apperror.WrongRequest, err)
	}

	reservation, err := h.purchaseSrv.GetPurchase(ctx, int(req.PharmacyId), req.PurchaseUuid, req.IsSocialCard)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to add to purchase: %w", errBase, err)
	}

	reservationItems := make([]*purchaseproto.ReservationProductInfo, 0, len(reservation.Products))
	for _, product := range reservation.Products {
		reservationItems = append(reservationItems, &purchaseproto.ReservationProductInfo{
			Name:  product.Name,
			Count: int64(product.Count),
			Price: int64(product.Price),
		})
	}

	return &purchaseproto.Reservation{
		Items:        reservationItems,
		SummaryPrice: int64(reservation.Price),
	}, nil
}

func (h *GRPCHandler) BuyReservation(ctx context.Context, req *purchaseproto.BuyReservationRequest) (*emptypb.Empty, error) {
	errBase := fmt.Sprintf("products.BuyReservation(%v)", req)

	if err := req.ValidateAll(); err != nil {
		return nil, apperror.NewClientError(apperror.WrongRequest, err)
	}

	if err := h.purchaseSrv.ConfirmPurchase(ctx, int(req.UserId), int(req.PharmacyId), req.PurchaseUuid, req.IsSocialCard); err != nil {
		return nil, fmt.Errorf("%s: failed to add to purchase: %w", errBase, err)
	}

	return &emptypb.Empty{}, nil
}
func (h *GRPCHandler) GetAvailableProductsToReservation(ctx context.Context, req *purchaseproto.PharmacyID) (*purchaseproto.Products, error) {
	errBase := fmt.Sprintf("products.GetAvailableProductsToReservation(%v)", req)

	if err := req.ValidateAll(); err != nil {
		return nil, apperror.NewClientError(apperror.WrongRequest, err)
	}

	products, err := h.purchaseSrv.GetAvailableToPurchaseProducts(ctx, int(req.PharmacyId))
	if err != nil {
		return nil, fmt.Errorf("%s: failed to add to purchase: %w", errBase, err)
	}

	res := make([]*purchaseproto.Product, 0, len(products))
	for _, product := range products {
		res = append(res, &purchaseproto.Product{
			Name:                        product.Name,
			Price:                       int64(product.Price),
			Count:                       int64(product.Count),
			Position:                    product.Position,
			NeedPrescriptionForMedicine: product.RecipeOnly,
		})
	}

	return &purchaseproto.Products{
		Products: res,
	}, nil
}
