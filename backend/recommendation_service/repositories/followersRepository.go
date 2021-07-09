package repositories

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/recommendation_service/model"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"strconv"

)

type FollowersRepository interface {
	CreateUserConnection(context.Context, model.Follower) (bool, error)
	GetAllFollowers(context.Context, string) ([]model.User, error)
	GetAllFollowing(context.Context, string) ([]model.User, error)
	GetAllFollowingsForHomepage(context.Context, string) ([]model.User, error)
	CreateUser(context.Context, model.User) (bool, error)
	DeleteDirectedConnection(context.Context, model.Follower) (bool, error)
	DeleteBiDirectedConnection (context.Context, model.Follower) (bool, error)
	UpdateUserConnection(context.Context, model.Follower) (*model.Follower,error)
	GetFollowersConnection(context.Context, model.Follower) (*model.Follower, error)
	GetCloseFriends(context.Context, string) ([]model.User, error)
    GetUsersForNotificationEnabled( context.Context,  string, string) ([]model.User, error)
	GetLimitedFriends( context.Context,  string, int) ([]model.User, error)
	GetUsersWithCommonConnectionsLimited(context.Context, string , string, int) ([]model.User, error)
	GetNumberOfCommonFriends( context.Context,  string,  string) (int, error)
	GetRandomUsers(ctx context.Context, limit int) ([]model.User, error)

}

type followersRepository struct {
	driver neo4j.Driver
}

func NewFollowersRepository(driver neo4j.Driver) (*followersRepository, error) {
	if driver == nil {
		panic("FollowersRepository not created, driver is nil")
	}

	return &followersRepository{driver: driver}, nil
}

func (repository *followersRepository) GetUsers(ctx context.Context, userId string, query string) ([]model.User, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetUsers")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	session := repository.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()
	var users []model.User
	_ , err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			query,
			map[string]interface{}{
				"UserId" : userId,
			})

		if err != nil {
			return nil, err
		}
		for result.Next() {
			user := result.Record().Values[0].(string)
			users = append(users, model.User{
				UserId: user,
			})

		}
		return users, nil
	})

	if err != nil {
		return nil, err
	}
	return users, nil
}

func (repository *followersRepository) GetRandomUsers(ctx context.Context, limit int) ([]model.User, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetRandomUsers")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	session := repository.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()
	var users []model.User

	result , err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (a:User) RETURN a.id LIMIT " + strconv.Itoa(limit) ,
			map[string]interface{}{
			})

		if err != nil {
			return nil, err
		}
		for result.Next() {
			user := result.Record().Values[0].(string)
			users = append(users, model.User{
				UserId: user,
			})

		}
		return false, errors.New("error: can not create user ")
	})
	if err != nil || result == nil {
		return nil, err
	}
	return users, nil
}

func (repository *followersRepository) GetLimitedFriends(ctx context.Context, userId  string, limit int) ([]model.User, error) {
	query := "MATCH (a:User {id : $UserId})-[r:Follows]->(b:User) WHERE r.IsApprovedRequest = true RETURN b.id LIMIT " + strconv.Itoa(limit)

	return repository.GetUsers(ctx, userId, query)
}

func (repository *followersRepository) GetUsersWithCommonConnectionsLimited(ctx context.Context, followerId string ,userId string, limit int) ([]model.User, error){
	query :=  "MATCH (a:User {id : "+"\""+followerId+"\""+ "})-[r:Follows]->(b:User) WHERE NOT (:User {id : $UserId})-[:Follows]->(b:User) RETURN b.id LIMIT "+ strconv.Itoa(limit)

	return repository.GetUsers(ctx, userId, query)
}

func (repository *followersRepository) GetNumberOfCommonFriends(ctx context.Context, randomFriendId string, userId string) (int, error) {
	query := "MATCH (a:User {id : " +"\""+randomFriendId +"\""+ "})-[r:Follows]-(b:User) WHERE (:User {id : $UserId})-[:Follows]->(b:User) RETURN b.id"

	users, err := repository.GetUsers(ctx, userId, query)
	return len(users), err
}




func (repository *followersRepository) GetAllFollowing(ctx context.Context, userId string) ([]model.User, error){
	query := "MATCH (a:User {id : $UserId})-[r:Follows]->(b:User) WHERE r.IsApprovedRequest = true RETURN b.id"
	return repository.GetUsers(ctx, userId, query)
}

func (repository *followersRepository) GetAllFollowers(ctx context.Context, userId string) ([]model.User, error){
	query := "MATCH (b:User)-[r:Follows]->(a:User {id : $UserId}) WHERE r.IsApprovedRequest = true RETURN b.id"
	return repository.GetUsers(ctx, userId, query)
}

func (repository *followersRepository) GetUsersForNotificationEnabled(ctx context.Context, userId string,notification string) ([]model.User, error) {
	query := "MATCH (a:User)-[r:Follows]->(b:User {id : $UserId}) WHERE r." + notification + " = true AND r.IsMuted = false RETURN a.id"
	return repository.GetUsers(ctx, userId, query)
}

func (repository *followersRepository) GetAllFollowingsForHomepage(ctx context.Context, userId string) ([]model.User, error){
	query := "MATCH (a:User {id : $UserId})-[r:Follows]->(b:User) WHERE r.IsApprovedRequest = true AND r.IsMuted = false RETURN b.id"
	return repository.GetUsers(ctx, userId, query)
}

func (repository *followersRepository) GetCloseFriends(ctx context.Context, id string) ([]model.User, error){
	query := "MATCH (a:User {id : $UserId})-[r:Follows]->(b:User) WHERE r.IsApprovedRequest = true AND r.IsMuted = false AND r.IsCloseFriend = true RETURN b.id"
	return repository.GetUsers(ctx, id, query)
}


func (repository *followersRepository) CreateUserConnection(ctx context.Context, f model.Follower) (bool, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateFollowersConnection")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	session := repository.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	result , err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (a:User {id : $UserId}), (b:User {id : $FollowerId}) MERGE" +
				"(a)-[:Follows {UserId : $UserId, FollowerId : $FollowerId ,IsMuted : $IsMuted," +
				" IsCloseFriend : $IsCloseFriend, IsApprovedRequest : $IsApprovedRequest, " +
				" RequestIsPending : $RequestIsPending, " +
				"IsMessageNotificationEnabled : $IsMessageNotificationEnabled," +
				"IsCommentNotificationEnabled : $IsCommentNotificationEnabled," +
				"IsPostNotificationEnabled : $IsPostNotificationEnabled," +
				"IsStoryNotificationEnabled : $IsStoryNotificationEnabled}]->(b)" +
				"RETURN a",
				map[string]interface{}{
				"UserId" : f.UserId,
				"FollowerId" : f.FollowerId,
				"IsMuted" : f.IsMuted,
				"IsCloseFriend" : f.IsCloseFriends,
				"IsApprovedRequest" : f.IsApprovedRequest,
				"IsMessageNotificationEnabled" : f.IsMessageNotificationEnabled,
				"IsPostNotificationEnabled" : f.IsPostNotificationEnabled,
				"IsCommentNotificationEnabled" : f.IsCommentNotificationEnabled,
				"IsStoryNotificationEnabled" : f.IsStoryNotificationEnabled,
				"RequestIsPending" : f.RequestIsPending,
			})
		if err != nil {
			return nil, err
		}
		if result.Next()  {
			return result.Record().Values[0], nil
		}
		return false, errors.New("error: can not create follower connection")
	})
	if err != nil || result == nil {
		return false, err
	}
	return true, nil
}

func (repository *followersRepository) UpdateUserConnection(ctx context.Context, f model.Follower) (*model.Follower,error){
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateUserConnection")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	session := repository.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	var follower = model.Follower{}
	_ , err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (a:User {id : $UserId})-[r:Follows]->(b:User {id : $FollowerId})" +
				" SET r.IsStoryNotificationEnabled = $IsStoryNotificationEnabled, r.IsCommentNotificationEnabled = $IsCommentNotificationEnabled, r.IsPostNotificationEnabled = $IsPostNotificationEnabled, r.IsMessageNotificationEnabled = $IsMessageNotificationEnabled, r.IsMuted = $IsMuted, r.IsCloseFriend = $IsCloseFriend, r.IsApprovedRequest = $IsApprovedRequest, r.RequestIsPending =  $RequestIsPending" +
				" RETURN r.UserId, r.FollowerId, r.IsMuted, r.IsCloseFriend, r.IsApprovedRequest, r.RequestIsPending, r.IsStoryNotificationEnabled, r.IsCommentNotificationEnabled, r.IsPostNotificationEnabled, r.IsMessageNotificationEnabled ",
			map[string]interface{}{
				"UserId" : f.UserId,
				"FollowerId" : f.FollowerId,
				"IsMuted" : f.IsMuted,
				"IsCloseFriend" : f.IsCloseFriends,
				"IsApprovedRequest" : f.IsApprovedRequest,
				"IsMessageNotificationEnabled" : f.IsMessageNotificationEnabled,
				"IsPostNotificationEnabled" : f.IsPostNotificationEnabled,
				"IsCommentNotificationEnabled" : f.IsCommentNotificationEnabled,
				"IsStoryNotificationEnabled" : f.IsStoryNotificationEnabled,
				"RequestIsPending" : f.RequestIsPending,
			})

		if err != nil {
			return nil, err
		}
		if result.Next()  {
			follower = model.Follower{
				UserId: result.Record().Values[0].(string),
				FollowerId: result.Record().Values[1].(string),
				IsMuted: result.Record().Values[2].(bool),
				IsCloseFriends: result.Record().Values[3].(bool),
				IsApprovedRequest: result.Record().Values[4].(bool),
				RequestIsPending: result.Record().Values[5].(bool),
				IsStoryNotificationEnabled: result.Record().Values[6].(bool),
				IsCommentNotificationEnabled: result.Record().Values[7].(bool),
				IsPostNotificationEnabled: result.Record().Values[8].(bool),
				IsMessageNotificationEnabled: result.Record().Values[9].(bool),
			}
			return nil, nil
		}
		return nil, errors.New("error: can not update users connection")
	})
	if err != nil{
		return nil, err
	}
	return &follower, nil
}

func (repository *followersRepository) GetFollowersConnection(ctx context.Context, f model.Follower) (*model.Follower, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetFollowersConnection")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	session := repository.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	var follower = model.Follower{}
	_ , err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (a:User {id : $UserId})-[r:Follows]->(b:User {id : $FollowerId})" +
				" RETURN r.UserId, r.FollowerId, r.IsMuted, r.IsCloseFriend, r.IsApprovedRequest, r.RequestIsPending, r.IsStoryNotificationEnabled, r.IsCommentNotificationEnabled, r.IsPostNotificationEnabled, r.IsMessageNotificationEnabled",
			map[string]interface{}{
				"UserId" : f.UserId,
				"FollowerId" : f.FollowerId,
			})

		if err != nil {
			return nil, err
		}
		if result.Next()  {
			follower = model.Follower{
				UserId: result.Record().Values[0].(string),
				FollowerId: result.Record().Values[1].(string),
				IsMuted: result.Record().Values[2].(bool),
				IsCloseFriends: result.Record().Values[3].(bool),
				IsApprovedRequest: result.Record().Values[4].(bool),
				RequestIsPending: result.Record().Values[5].(bool),
				IsStoryNotificationEnabled: result.Record().Values[6].(bool),
				IsCommentNotificationEnabled: result.Record().Values[7].(bool),
				IsPostNotificationEnabled: result.Record().Values[8].(bool),
				IsMessageNotificationEnabled: result.Record().Values[9].(bool),
			}
			return nil, nil
		}
		return &model.Follower{}, nil
	})
	if err != nil{
		return nil, err
	}
	return &follower, nil

}

func (repository *followersRepository) CreateUser(ctx context.Context, u model.User) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateFollowersConnection")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	session := repository.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	result , err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"CREATE (n:User {id : $UserId}) RETURN n",
			map[string]interface{}{
				"UserId" : u.UserId,
			})

		if err != nil {
			return nil, err
		}
		if result.Next()  {
			return true, nil
		}
		return false, errors.New("error: can not create user ")
	})
	if err != nil || result == nil {
		return false, err
	}
	return true, nil
}

func (repository *followersRepository) DeleteDirectedConnection(ctx context.Context, f model.Follower) (bool, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "DeleteConnection")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	session := repository.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	_ , err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		_, err := transaction.Run(
			"MATCH (a:User {id : $UserId})-[r:Follows]->(b:User {id : $FollowerId}) DELETE r",
			map[string]interface{}{
				"UserId" : f.UserId,
				"FollowerId" : f.FollowerId,
			})

		if err != nil {
			return false, err
		}
		return true, nil
	})
	if err != nil {
		return false, err
	}
	return true, nil


}

//Kada se useri blokiraju - u ovom slucaju ne posmatramo rezultat operacija, jer ukoliko konekcije izmedju usera ne postoje,
//svakako treba dozvoliti useru da blokira
func (repository *followersRepository) DeleteBiDirectedConnection(ctx context.Context, f model.Follower) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "DeleteBiDirectedConnection")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	session := repository.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		transaction.Run(
			"MATCH (a:User {id : $UserId})-[r:Follows]->(b:User {id : $FollowerId}) DELETE r",
			map[string]interface{}{
				"UserId" : f.UserId,
				"FollowerId" : f.FollowerId,
			})

		transaction.Run(
			"MATCH (b:User {id : $FollowerId})-[r:Follows]->(a:User {id : $UserId}) DELETE r",
			map[string]interface{}{
				"UserId" : f.UserId,
				"FollowerId" : f.FollowerId,
			})

		return true, nil
	})
	return true, nil
}