package services

import (
	"context"
	"errors"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/user_service/repositories"
	"github.com/david-drvar/xws2021-nistagram/user_service/saga"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
)

type PrivacyService struct {
	repository  repositories.PrivacyRepository
	userService *UserService
}

func NewPrivacyService(db *gorm.DB, redis *saga.RedisServer) (*PrivacyService, error) {
	repository, err := repositories.NewPrivacyRepo(db)
	service, err := NewUserService(db, redis)
	return &PrivacyService{
		repository:  repository,
		userService: service,
	}, err
}

func (service *PrivacyService) CreatePrivacy(ctx context.Context, privacy *persistence.Privacy) (persistence.Privacy, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreatePrivacy")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.repository.CreatePrivacy(ctx, privacy)
}

func (service *PrivacyService) UpdatePrivacy(ctx context.Context, privacy *persistence.Privacy) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdatePrivacy")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.repository.UpdatePrivacy(ctx, privacy)
}

func (service *PrivacyService) BlockUser(ctx context.Context, block *persistence.BlockedUsers) (bool, error) {
	//TODO Proveri da li ti useri postoje i posalji zahtev da im se obrise prijateljstvo
	span := tracer.StartSpanFromContextMetadata(ctx, "BlockUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	res, err := service.repository.BlockUser(ctx, block)
	if !res || err != nil {
		return false, errors.New("Could not block user in privacy service")
	}

	var conn *grpc.ClientConn
	conn, err = grpc.Dial(":8095", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := protopb.NewFollowersClient(conn)

	createFollowerRequest := protopb.CreateFollowerRequest{
		Follower: &protopb.Follower{
			UserId:     block.UserId,
			FollowerId: block.BlockedUserId,
		},
	}

	_, err = c.DeleteBiDirectedConnection(context.Background(), &createFollowerRequest)
	if err != nil {
		log.Fatalf("Could not block user in followers service")
	}

	return true, nil
}

func (service *PrivacyService) UnBlockUser(ctx context.Context, block *persistence.BlockedUsers) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UnBlockUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.repository.UnBlockUser(ctx, block)
}

func (service *PrivacyService) CheckUserProfilePublic(ctx context.Context, userId string) bool {
	span := tracer.StartSpanFromContextMetadata(ctx, "CheckUserProfilePublic")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	privacy, _ := service.repository.GetUserPrivacy(ctx, userId)
	return privacy.IsProfilePublic == true
}

func (service *PrivacyService) GetAllPublicUsers(ctx context.Context) []string {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllPublicUsers")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	privacies, _ := service.repository.GetAlLPublicUsers(ctx)

	publicUsers := []string{}
	for _, privacy := range privacies {
		if privacy.IsProfilePublic {
			publicUsers = append(publicUsers, privacy.UserId)
		}
	}

	return publicUsers
}

func (service *PrivacyService) CheckIfBlocked(ctx context.Context, userId string, blockedUserId string) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllPublicUsers")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.repository.CheckIfBlocked(ctx, userId, blockedUserId)
}

func (service *PrivacyService) GetBlockedUsers(ctx context.Context, userId string) ([]persistence.BlockedUsers, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllPublicUsers")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.repository.GetBlockedUsers(ctx, userId)
}

func (service *PrivacyService) GetUserPrivacy(ctx context.Context, userId string) (*persistence.Privacy, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUserPrivacy")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.repository.GetUserPrivacy(ctx, userId)
}
