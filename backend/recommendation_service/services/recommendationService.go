package services

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/recommendation_service/model"
	"github.com/david-drvar/xws2021-nistagram/recommendation_service/repositories"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"sort"
)

type UserCommonFriends struct {
	User model.User
	PercentageInRec int
}

type RecommendationService struct {
	followersRepository repositories.FollowersRepository
}

func NewRecommendationService (driver neo4j.Driver) (*RecommendationService, error) {
	//pristup nasim prijateljima -> odabir odredjenog broja
	//pristup odredjenom broju njihovih prijatelja sa najvise zajednickih prijatelja sa nama
	//uzimanje njihovih interesovanja
	//random formula za racunanje procenta, i limitiranje broja datih usera
	repo, err := repositories.NewFollowersRepository(driver)
	if err != nil {
		return &RecommendationService{}, err
	}

	return &RecommendationService{followersRepository: repo},nil
}

func (recommendation *RecommendationService) RecommendationPattern(ctx context.Context, userId string) ([]UserCommonFriends, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "RecommendationPattern")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var retVal []UserCommonFriends
	var potentialUsers []model.User

	//Odredjeni broj prijatelja
	limitedFriends, err := recommendation.followersRepository.GetLimitedFriends(ctx, userId, 5)
	if err != nil {
		return nil, errors.New("Could not get limited friends!")
	}

	//Odredjeni broj usera koje ne prati, ali imaju zajednicke prijatelje
	for _, friend := range limitedFriends{
		users, err := recommendation.followersRepository.GetUsersWithCommonConnectionsLimited(ctx, friend.UserId, userId, 5)
		if err == nil {
			potentialUsers = append(potentialUsers, users ...)
		}
	}

	//Broj zajednickih prijatelja
	for _, friend := range potentialUsers {
		result, err := recommendation.followersRepository.GetNumberOfCommonFriends(ctx, friend.UserId, userId)
		if err == nil {
			retVal = append(retVal, UserCommonFriends{User: friend, PercentageInRec: result})
		}
	}

	if len(retVal) == 0 {
		result, err := recommendation.followersRepository.GetRandomUsers(ctx, 5)
		if err != nil {
			return nil, err
		}
		for _, f := range result{
			retVal = append(retVal, UserCommonFriends{User: f, PercentageInRec: 0})
		}
		return retVal, nil
	}

	//TODO zajednicki interesi!


	sort.SliceStable(retVal, func(i, j int) bool {
		return retVal[i].PercentageInRec < retVal[j].PercentageInRec
	})

	return retVal, nil
}


