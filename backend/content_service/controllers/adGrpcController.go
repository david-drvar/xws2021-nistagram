package controllers

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type AdGrpcController struct {
	service 	*services.AdService
	jwtManager  *common.JWTManager
}

func NewAdController(db *gorm.DB, jwtManager *common.JWTManager) (*AdGrpcController, error) {
	service, err := services.NewAdService(db)
	if err != nil {
		return nil, err
	}

	return &AdGrpcController{
		service,
		jwtManager,
	}, nil
}

// Updating and deleting Ads is not be allowed, only its Campaign data can be changed.

func (controller *AdGrpcController) GetAds(ctx context.Context, in *protopb.EmptyRequestContent) (*protopb.AdArray, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAds")
	defer span.Finish()
	// claims, _ := controller.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	response := []*protopb.Ad{}

	return &protopb.AdArray{
		Ads: response,
	}, nil
}

func (controller *AdGrpcController) GetAdsFromInfluencer(ctx context.Context, in *protopb.RequestId) (*protopb.AdArray, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAdsFromInfluencer")
	defer span.Finish()
	// claims, _ := controller.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	ads, err := controller.service.GetAdsFromInfluencer(ctx, in.Id)
	if err != nil { return &protopb.AdArray{}, err }

	response := []*protopb.Ad{}
	for _, ad := range ads{
		response = append(response, ad.ConvertToGrpc())
	}

	return &protopb.AdArray{
		Ads: response,
	}, nil
}

func (controller *AdGrpcController) CreateAd(ctx context.Context, in *protopb.Ad) (*protopb.EmptyResponseContent, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateAd")
	defer span.Finish()
	// claims, _ := controller.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var ad *domain.Ad
	ad = ad.ConvertFromGrpc(in)

	err := controller.service.CreateAd(ctx, *ad)
	if err != nil { return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "could not create an ad") }

	return &protopb.EmptyResponseContent{}, nil
}

func (controller *AdGrpcController) GetAdCategories(ctx context.Context, in *protopb.EmptyRequestContent) (*protopb.AdCategoryArray, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAdCategories")
	defer span.Finish()
	// claims, _ := controller.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	categories, err := controller.service.GetAdCategories(ctx)
	if err != nil { return &protopb.AdCategoryArray{}, err }

	responseCategories := []*protopb.AdCategory{}
	for _, category := range categories{
		responseCategories = append(responseCategories, category.ConvertToGrpc())
	}

	return &protopb.AdCategoryArray{
		Categories: responseCategories,
	}, nil
}

func (controller *AdGrpcController) GetAdCategory(ctx context.Context, in *protopb.RequestId) (*protopb.AdCategory, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAdCategory")
	defer span.Finish()
	// claims, _ := controller.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	category, err := controller.service.GetAdCategory(ctx, in.Id)
	if err != nil { return &protopb.AdCategory{}, err }

	return category.ConvertToGrpc(), nil
}

func (controller *AdGrpcController) CreateAdCategory(ctx context.Context, in *protopb.AdCategory) (*protopb.EmptyResponseContent, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateAdCategory")
	defer span.Finish()
	// claims, _ := controller.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var category domain.AdCategory
	category = category.ConvertFromGrpc(in)

	err := controller.service.CreateAdCategory(ctx, category)
	if err != nil { return &protopb.EmptyResponseContent{}, err }

	return &protopb.EmptyResponseContent{}, nil
}

func (controller *AdGrpcController) GetUsersAdCategories(ctx context.Context, in *protopb.EmptyRequestContent) (*protopb.AdCategoryArray, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUsersAdCategories")
	defer span.Finish()
	claims, _ := controller.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	categories, err := controller.service.GetUsersAdCategories(ctx, claims.UserId)
	if err != nil { return &protopb.AdCategoryArray{}, err }

	responseCategories := []*protopb.AdCategory{}
	for _, category := range categories{
		responseCategories = append(responseCategories, category.ConvertToGrpc())
	}

	return &protopb.AdCategoryArray{
		Categories: responseCategories,
	}, nil
}

func (controller *AdGrpcController) UpdateUsersAdCategories(ctx context.Context, in *protopb.AdCategoryArray) (*protopb.EmptyResponseContent, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateUsersAdCategories")
	defer span.Finish()
	claims, _ := controller.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	categories := []domain.AdCategory{}
	for _, category := range in.Categories{
		var domainCategory domain.AdCategory
		categories = append(categories, domainCategory.ConvertFromGrpc(category))
	}

	err := controller.service.UpdateUsersAdCategories(ctx, claims.UserId, categories)
	if err != nil { return &protopb.EmptyResponseContent{}, err }

	return &protopb.EmptyResponseContent{}, nil
}

func (controller *AdGrpcController) CreateUserAdCategories(ctx context.Context, in *protopb.RequestId) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUserAdCategories")
	defer span.Finish()
	// claims, _ := controller.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := controller.service.CreateUserAdCategories(ctx, in.Id)
	if err != nil { return &protopb.EmptyResponseContent{}, err }

	return &protopb.EmptyResponseContent{}, nil
}

func (controller *AdGrpcController) IncrementLinkClicks(ctx context.Context, in *protopb.RequestId) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "IncrementLinkClicks")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := controller.service.IncrementLinkClicks(ctx, in.Id)
	if err != nil { return &protopb.EmptyResponseContent{}, err }

	return &protopb.EmptyResponseContent{}, nil
}