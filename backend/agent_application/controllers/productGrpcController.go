package controllers

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/agent_application/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/agent_application/services"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/logger"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type ProductGrpcController struct {
	service    *services.ProductService
	jwtManager *common.JWTManager
	logger     *logger.Logger
}

func NewProductController(db *gorm.DB, jwtManager *common.JWTManager, logger *logger.Logger) (*ProductGrpcController, error) {
	service, err := services.NewProductService(db)
	if err != nil {
		return nil, err
	}

	return &ProductGrpcController{
		service,
		jwtManager,
		logger,
	}, nil
}

func (s *ProductGrpcController) CreateProduct(ctx context.Context, in *protopb.Product) (*protopb.EmptyResponseAgent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateProduct")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var product *persistence.Product
	product = product.ConvertFromGrpc(in)
	if product == nil {
		return &protopb.EmptyResponseAgent{}, status.Errorf(codes.Unknown, "cannot create product")
	}

	err := s.service.CreateProduct(ctx, *product)
	if err != nil {
		return &protopb.EmptyResponseAgent{}, status.Errorf(codes.Unknown, err.Error())
	}

	return &protopb.EmptyResponseAgent{}, nil
}

func (s *ProductGrpcController) GetAllProductsByAgentId(ctx context.Context, in *protopb.UserAgentApp) (*protopb.ProductsArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllProductsByAgentId")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	agentId := in.Id

	products, err := s.service.GetAllProductsByAgentId(ctx, agentId)
	if err != nil {
		return &protopb.ProductsArray{}, status.Errorf(codes.Unknown, err.Error())
	}

	responseProducts := []*protopb.Product{}
	for _, product := range products {
		responseProducts = append(responseProducts, product.ConvertToGrpc())
	}

	return &protopb.ProductsArray{Products: responseProducts}, nil
}

func (s *ProductGrpcController) GetAllProducts(ctx context.Context, in *protopb.EmptyRequestAgent) (*protopb.ProductsArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllProducts")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	products, err := s.service.GetAllProducts(ctx)
	if err != nil {
		return &protopb.ProductsArray{}, status.Errorf(codes.Unknown, err.Error())
	}

	responseProducts := []*protopb.Product{}
	for _, product := range products {
		responseProducts = append(responseProducts, product.ConvertToGrpc())
	}

	return &protopb.ProductsArray{Products: responseProducts}, nil
}

func (s *ProductGrpcController) GetProductById(ctx context.Context, in *protopb.Product) (*protopb.Product, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetProductById")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	productId := in.Id

	product, err := s.service.GetProductById(ctx, productId)
	if err != nil {
		return &protopb.Product{}, status.Errorf(codes.Unknown, err.Error())
	}

	var responseProduct *protopb.Product
	responseProduct = product.ConvertToGrpc()

	return responseProduct, nil
}

func (s *ProductGrpcController) DeleteProduct(ctx context.Context, in *protopb.Product) (*protopb.EmptyResponseAgent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "DeleteProduct")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	productId := in.Id

	err := s.service.DeleteProduct(ctx, productId)
	if err != nil {
		return &protopb.EmptyResponseAgent{}, status.Errorf(codes.Unknown, err.Error())
	}

	return &protopb.EmptyResponseAgent{}, nil
}

func (s *ProductGrpcController) UpdateProduct(ctx context.Context, in *protopb.Product) (*protopb.EmptyResponseAgent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateProduct")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var product *persistence.Product
	product = product.ConvertFromGrpc(in)
	if product == nil {
		return &protopb.EmptyResponseAgent{}, status.Errorf(codes.Unknown, "cannot create product")
	}

	err := s.service.UpdateProduct(ctx, product)
	if err != nil {
		return &protopb.EmptyResponseAgent{}, status.Errorf(codes.Unknown, err.Error())
	}

	return &protopb.EmptyResponseAgent{}, nil
}

func (s *ProductGrpcController) OrderProduct(ctx context.Context, in *protopb.Order) (*protopb.EmptyResponseAgent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateProduct")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var order *persistence.Order
	order = order.ConvertFromGrpc(in)
	if order == nil {
		return &protopb.EmptyResponseAgent{}, status.Errorf(codes.Unknown, "cannot create product")
	}

	err := s.service.OrderProduct(ctx, order)
	if err != nil {
		return &protopb.EmptyResponseAgent{}, status.Errorf(codes.Unknown, err.Error())
	}

	return &protopb.EmptyResponseAgent{}, nil
}

func (s *ProductGrpcController) GetOrdersByUser(ctx context.Context, in *protopb.UserAgentApp) (*protopb.OrdersArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetOrdersByUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	orders, err := s.service.GetOrdersByUser(ctx, in.Id)
	if err != nil {
		return &protopb.OrdersArray{}, status.Errorf(codes.Unknown, err.Error())
	}

	responseOrders := []*protopb.Order{}
	for _, order := range orders {
		responseOrders = append(responseOrders, order.ConvertToGrpc())
	}

	return &protopb.OrdersArray{Orders: responseOrders}, nil
}

func (s *ProductGrpcController) GetOrdersByAgent(ctx context.Context, in *protopb.UserAgentApp) (*protopb.OrdersArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetOrdersByAgent")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	orders, err := s.service.GetOrdersByAgent(ctx, in.Id)
	if err != nil {
		return &protopb.OrdersArray{}, status.Errorf(codes.Unknown, err.Error())
	}

	responseOrders := []*protopb.Order{}
	for _, order := range orders {
		responseOrders = append(responseOrders, order.ConvertToGrpc())
	}

	return &protopb.OrdersArray{Orders: responseOrders}, nil
}
