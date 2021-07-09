package controllers

import (
	"github.com/david-drvar/xws2021-nistagram/common/logger"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/recommendation_service/services"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"golang.org/x/net/context"
)

type RecommendationGrpcController struct {
	service *services.RecommendationService
	logger  *logger.Logger
}

func NewRecommendationGrpcController (driver neo4j.Driver, logger *logger.Logger) (RecommendationGrpcController, error) {
	service, err := services.NewRecommendationService(driver)
	if err != nil {
		return RecommendationGrpcController{}, err
	}
	return RecommendationGrpcController{
		logger: logger,
		service: service,
	}, nil
}

func (controller RecommendationGrpcController) RecommendationPattern(ctx context.Context, in *protopb.RequestIdFollowers) (*protopb.RecommendationResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "RecommendationPattern")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var retVal []*protopb.Recommendation
	var result []services.UserCommonFriends
	result, err := controller.service.RecommendationPattern(ctx, in.Id)
	if err != nil {
		return &protopb.RecommendationResponse{}, err
	}
	for _, r := range result{
		retVal = append(retVal, &protopb.Recommendation{UserId: r.User.UserId, Percentage: string(r.PercentageInRec)})
	}
	return &protopb.RecommendationResponse{Recommendations: retVal}, nil
}
