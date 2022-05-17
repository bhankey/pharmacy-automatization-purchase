package container

import (
	"github.com/bhankey/pharmacy-automatization-purchase/internal/adapter/repository/productrepo"
	"github.com/bhankey/pharmacy-automatization-purchase/internal/adapter/repository/receiptrepo"
	purchase "github.com/bhankey/pharmacy-automatization-purchase/internal/delivery/grpc/v1/products"
	"github.com/bhankey/pharmacy-automatization-purchase/internal/service/purchaseservice"
)

func (c *Container) GetPurchaseGRPCHandler() *purchase.GRPCHandler {
	const key = "PurchaseGRPCHandler"

	dependency, ok := c.dependencies[key]
	if ok {
		typedDependency, ok := dependency.(*purchase.GRPCHandler)
		if ok {
			return typedDependency
		}
	}

	typedDependency := purchase.NewPurchaseGRPCHandler(c.getPurchaseSrv())

	c.dependencies[key] = typedDependency

	return typedDependency
}

func (c *Container) getPurchaseSrv() *purchaseservice.Service {
	const key = "PurchaseSrv"

	dependency, ok := c.dependencies[key]
	if ok {
		typedDependency, ok := dependency.(*purchaseservice.Service)
		if ok {
			return typedDependency
		}
	}

	typedDependency := purchaseservice.NewPurchaseService(
		c.getProductStorage(),
		c.getReceiptStorage(),
	)

	c.dependencies[key] = typedDependency

	return typedDependency
}

func (c *Container) getProductStorage() *productrepo.Repository {
	const key = "ProductStorage"

	dependency, ok := c.dependencies[key]
	if ok {
		typedDependency, ok := dependency.(*productrepo.Repository)
		if ok {
			return typedDependency
		}
	}

	typedDependency := productrepo.NewProductRepo(c.masterPostgresDB, c.slavePostgresDB)

	c.dependencies[key] = typedDependency

	return typedDependency
}

func (c *Container) getReceiptStorage() *receiptrepo.Repository {
	const key = "ReceiptStorage"

	dependency, ok := c.dependencies[key]
	if ok {
		typedDependency, ok := dependency.(*receiptrepo.Repository)
		if ok {
			return typedDependency
		}
	}

	typedDependency := receiptrepo.NewReceiptRepo(c.masterPostgresDB, c.slavePostgresDB)

	c.dependencies[key] = typedDependency

	return typedDependency
}
