package services

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/agent_application/model/domain"
	"github.com/david-drvar/xws2021-nistagram/agent_application/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/agent_application/repositories"
	"github.com/david-drvar/xws2021-nistagram/agent_application/util/images"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"gorm.io/gorm"
)

type ProductService struct {
	productRepository repositories.ProductRepository
}

func NewProductService(db *gorm.DB) (*ProductService, error) {
	productRepository, err := repositories.NewProductRepo(db)

	return &ProductService{
		productRepository,
	}, err
}

func (service *ProductService) CreateProduct(ctx context.Context, product persistence.Product) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUserInAgentApp")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := service.productRepository.CreateProduct(ctx, product)
	if err != nil {
		return err
	}

	return nil
}

func (service *ProductService) GetAllProductsByAgentId(ctx context.Context, id string) ([]persistence.Product, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllProductsByAgentId")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	products, err := service.productRepository.GetAllProductsByAgentId(ctx, id)
	if err != nil {
		return nil, err
	}

	var finalProducts []persistence.Product

	for _, product := range products {
		base64Image, err := images.LoadImageToBase64(product.Photo)
		if err != nil {
			return nil, err
		}
		product.Photo = base64Image
		finalProducts = append(finalProducts, product)
	}

	return finalProducts, nil
}

func (service *ProductService) GetAllProducts(ctx context.Context) ([]persistence.Product, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllProducts")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	products, err := service.productRepository.GetAllProducts(ctx)
	if err != nil {
		return nil, err
	}

	var finalProducts []persistence.Product

	for _, product := range products {
		base64Image, err := images.LoadImageToBase64(product.Photo)
		if err != nil {
			return nil, err
		}
		product.Photo = base64Image
		finalProducts = append(finalProducts, product)
	}

	return finalProducts, nil
}

func (service *ProductService) GetProductById(ctx context.Context, id string) (persistence.Product, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetProductById")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	product, err := service.productRepository.GetProductById(ctx, id)
	if err != nil {
		return persistence.Product{}, err
	}

	base64Image, err := images.LoadImageToBase64(product.Photo)
	if err != nil {
		return persistence.Product{}, err
	}
	product.Photo = base64Image

	return product, nil
}

func (service *ProductService) DeleteProduct(ctx context.Context, id string) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "DeleteProduct")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.productRepository.DeleteProduct(ctx, id)
}

func (service *ProductService) UpdateProduct(ctx context.Context, product *persistence.Product) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateProduct")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.productRepository.UpdateProduct(ctx, product)
}

func (service *ProductService) OrderProduct(ctx context.Context, order *persistence.Order) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "OrderProduct")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.productRepository.OrderProduct(ctx, order)
}

func (service *ProductService) GetOrdersByUser(ctx context.Context, userId string) ([]domain.Order, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "OrderProduct")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.productRepository.GetOrdersByUser(ctx, userId)
}

func (service *ProductService) GetOrdersByAgent(ctx context.Context, agentId string) ([]domain.Order, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetOrdersByAgent")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.productRepository.GetOrdersByAgent(ctx, agentId)
}
