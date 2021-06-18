package rbac

import "gorm.io/gorm"

func SetupUsersRBAC(db *gorm.DB) error {
	dropUsersTables(db)
	err := db.AutoMigrate(&UserRole{}, Permission{}, RolePermission{})
	if err != nil {
		return err
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		userRoles := []UserRole{basic, admin, verified, agent, nonregistered}
		result := db.Create(&userRoles)
		if result.Error != nil {
			return result.Error
		}

		permissions := []Permission{
			getUserById, getUsernameById, getPhotoById, getAllUsers,
			updateUserPassword, updateUserProfile, searchUser, checkIsApproved,
			getUserByUsername, submitVerificationRequest,
			updatePrivacy, blockUser, unBlockUser, checkIfBlocked, checkUserProfilePublic, getAllPublicUsers, approveAccount, getUserPrivacy,
			createNotification,
			getPendingVerificationRequests, changeVerificationRequestStatus, getVerificationRequestsByUserId, getAllVerificationRequests,
			updateUserPhoto, getUserNotifications,getBlockedUsers,
		}
		result = db.Create(&permissions)
		if result.Error != nil {
			return result.Error
		}

		rolePermissions := []RolePermission{
			basicGetUserById, agentGetUserById, adminGetUserById, verifiedGetUserById, nonregisteredGetUserById,
			basicGetUsernameById, adminGetUsernameById, agentGetUsernameById, verifiedGetUsernameById, nonregisteredGetUsernameById,
			basicGetPhotoById, adminGetPhotoById, agentGetPhotoById, verifiedGetPhotoById, nonregisteredGetPhotoById,
			basicGetAllUsers, agentGetAllUsers, adminGetAllUsers, verifiedGetAllUsers, nonregisteredGetAllUsers,
			basicUpdateUserPassword, agentUpdateUserPassword, verifiedUpdateUserPassword, adminUpdateUserPassword,
			basicUpdateUserProfile, agentUpdateUserProfile, verifiedUpdateUserProfile, adminUpdateUserProfile,
			basicSearchUser, agentSearchUser, adminSearchUser, verifiedSearchUser, nonregisteredSearchUser,
			basicUpdatePrivacy, agentUpdatePrivacy, verifiedUpdatePrivacy,
			basicBlockUser, agentBlockUser, verifiedBlockUser,
			basicUnBlockUser, agentUnBlockUser, verifiedUnBlockUser,
			basicGetUserPrivacy, agentGetUserPrivacy, verifiedGetUserPrivacy, adminGetUserPrivacy, nonregisteredGetUserPrivacy,
			basicCheckIfBlocked, agentCheckIfBlocked, adminCheckIfBlocked, verifiedCheckIfBlocked, nonregisteredCheckIfBlocked,
			basicCheckUserProfilePublic, agentCheckUserProfilePublic, adminCheckUserProfilePublic, verifiedCheckUserProfilePublic, nonregisteredCheckUserProfilePublic,
			basicGetAllPublicUsers, agentGetAllPublicUsers, adminGetAllPublicUsers, verifiedGetAllPublicUsers, nonregisteredGetAllPublicUsers,
			basicCreateNotification, adminCreateNotification, agentCreateNotification, verifiedCreateNotification, nonregisteredCreateNotification,
			basicApproveAccount, adminApproveAccount, agentApproveAccount, basicCheckIsApproved, adminCheckIsApproved, agentCheckIsApproved,
			basicGetUserByUsername, agentGetUserByUsername, adminGetUserByUsername, verifiedGetUserByUsername, nonregisteredGetUserByUsername,
			basicSubmitVerificationRequest, agentSubmitVerificationRequest, verifiedSubmitVerificationRequest,
			adminGetPendingVerificationRequests,
			adminChangeVerificationRequestStatus,
			basicGetVerificationRequestsByUserId, verifiedGetVerificationRequestsByUserId, adminGetVerificationRequestsByUserId, agentGetVerificationRequestsByUserId,
			adminGetAllVerificationRequests,
			basicUpdateUserPhoto, agentUpdateUserPhoto, adminUpdateUserPhoto, verifiedUpdateUserPhoto,
			basicGetBlockedUsers, agentGetBlockedUsers, adminGetBlockedUsers, verifiedGetBlockedUsers,
			basicGetUserNotifications, adminGetUserNotifications, verifiedGetUserNotifications, nonregisteredGetUserNotifications,
		}
		result = db.Create(&rolePermissions)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	return err
}

func dropUsersTables(db *gorm.DB) {
	if db.Migrator().HasTable(&UserRole{}) {
		db.Migrator().DropTable(&UserRole{})
	}
	if db.Migrator().HasTable(&Permission{}) {
		db.Migrator().DropTable(&Permission{})
	}
	if db.Migrator().HasTable(&RolePermission{}) {
		db.Migrator().DropTable(&RolePermission{})
	}
}

var (
	checkIsApproved   = Permission{Id: "c7f5bfa5-9749-4be3-a6bb-451a5acbd1b0", Name: "CheckIsApproved"}
	getUserByUsername = Permission{Id: "aa3305b0-0b68-490f-b38e-5a0c1cf97a9e", Name: "GetUserByUsername"}

	getUserById        = Permission{Id: "992d5bf5-3e7f-4c8e-a76a-ad8444c0b32e", Name: "GetUserById"}
	getUsernameById    = Permission{Id: "c9295278-8fe8-40e6-9c9d-75543d48fa94", Name: "GetUsernameById"}
	getPhotoById	   = Permission{Id: "b6fc6b92-f2b4-471c-b5bc-b6f0a442759e", Name: "GetPhotoById"}
	getAllUsers        = Permission{Id: "26e41719-e309-4591-bb7e-3291b59c6dd4", Name: "GetAllUsers"}
	updateUserProfile  = Permission{Id: "48719e11-38de-4935-a93a-a61886c9303e", Name: "UpdateUserProfile"}
	updateUserPassword = Permission{Id: "50db6a87-483e-4d97-b320-9ff68235280a", Name: "UpdateUserPassword"}
	searchUser         = Permission{Id: "afbbf68f-ad1d-45db-8345-37ab619eea33", Name: "SearchUser"}
	approveAccount     = Permission{Id: "15e9a950-8581-4aa6-81c1-ae722c527051", Name: "ApproveAccount"}
	updatePrivacy                   = Permission{Id: "3ce13d71-30e2-4cca-8a48-8a5ee1b6a42e", Name: "UpdatePrivacy"}
	blockUser                       = Permission{Id: "9ec3fb89-28d6-4789-82b8-f247706cb6a0", Name: "BlockUser"}
	unBlockUser                     = Permission{Id: "bf4632b1-e3ae-41d5-a04a-4bac73b7a2ef", Name: "UnBlockUser"}
	checkIfBlocked                  = Permission{Id: "ce7b4f42-02ce-4c92-bcc4-529972173a4b", Name: "CheckIfBlocked"}
	checkUserProfilePublic          = Permission{Id: "f2d282a9-c171-47f4-935d-32875fa61c8a", Name: "CheckUserProfilePublic"}
	getAllPublicUsers               = Permission{Id: "a30c3350-4b8f-4773-bd90-32c17e88d221", Name: "GetAllPublicUsers"}
	submitVerificationRequest       = Permission{Id: "1d867c15-595b-4a69-b8ad-7135457bc402", Name: "SubmitVerificationRequest"}
	getPendingVerificationRequests  = Permission{Id: "56a15e9b-3522-4d32-a11a-2fd869a41489", Name: "GetPendingVerificationRequests"}
	getAllVerificationRequests      = Permission{Id: "201569e4-b294-4b20-93f9-cd9d41e433bf", Name: "GetAllVerificationRequests"}
	changeVerificationRequestStatus = Permission{Id: "63a03b3a-46d6-4780-8517-fa61fbf8feba", Name: "ChangeVerificationRequestStatus"}
	getVerificationRequestsByUserId = Permission{Id: "735e6566-cf91-11eb-b8bc-0242ac130003", Name: "GetVerificationRequestsByUserId"}
	createNotification              = Permission{Id: "c6b63d7c-8344-43f4-b7c0-fb5e353aa2ae", Name: "CreateNotification"}
	updateUserPhoto                 = Permission{Id: "042cef39-9acb-49d9-8088-1a583623bfa0", Name: "UpdateUserPhoto"}
	getUserNotifications            = Permission{Id: "2687d1e4-cf89-11eb-b8bc-0242ac130003", Name: "GetUserNotifications"}
	getUserPrivacy                  = Permission{Id: "221ee966-d025-11eb-b8bc-0242ac130003", Name: "GetUserPrivacy"}
	getBlockedUsers                  = Permission{Id: "bb400be1-7dcb-439c-9aba-235b566ec1fd", Name: "GetBlockedUsers"}
)

var (
	basicGetUserById         = RolePermission{RoleId: basic.Id, PermissionId: getUserById.Id}
	agentGetUserById         = RolePermission{RoleId: agent.Id, PermissionId: getUserById.Id}
	adminGetUserById         = RolePermission{RoleId: admin.Id, PermissionId: getUserById.Id}
	verifiedGetUserById      = RolePermission{RoleId: verified.Id, PermissionId: getUserById.Id}
	nonregisteredGetUserById = RolePermission{RoleId: nonregistered.Id, PermissionId: getUserById.Id}

	basicGetUserByUsername         = RolePermission{RoleId: basic.Id, PermissionId: getUserByUsername.Id}
	agentGetUserByUsername         = RolePermission{RoleId: agent.Id, PermissionId: getUserByUsername.Id}
	adminGetUserByUsername         = RolePermission{RoleId: admin.Id, PermissionId: getUserByUsername.Id}
	verifiedGetUserByUsername      = RolePermission{RoleId: verified.Id, PermissionId: getUserByUsername.Id}
	nonregisteredGetUserByUsername = RolePermission{RoleId: nonregistered.Id, PermissionId: getUserByUsername.Id}

	basicGetUsernameById         = RolePermission{RoleId: basic.Id, PermissionId: getUsernameById.Id}
	agentGetUsernameById         = RolePermission{RoleId: agent.Id, PermissionId: getUsernameById.Id}
	adminGetUsernameById         = RolePermission{RoleId: admin.Id, PermissionId: getUsernameById.Id}
	verifiedGetUsernameById      = RolePermission{RoleId: verified.Id, PermissionId: getUsernameById.Id}
	nonregisteredGetUsernameById = RolePermission{RoleId: nonregistered.Id, PermissionId: getUsernameById.Id}

	basicGetPhotoById         = RolePermission{RoleId: basic.Id, PermissionId: getPhotoById.Id}
	agentGetPhotoById         = RolePermission{RoleId: agent.Id, PermissionId: getPhotoById.Id}
	adminGetPhotoById         = RolePermission{RoleId: admin.Id, PermissionId: getPhotoById.Id}
	verifiedGetPhotoById      = RolePermission{RoleId: verified.Id, PermissionId: getPhotoById.Id}
	nonregisteredGetPhotoById = RolePermission{RoleId: nonregistered.Id, PermissionId: getPhotoById.Id}

	basicApproveAccount = RolePermission{RoleId: basic.Id, PermissionId: approveAccount.Id}
	agentApproveAccount = RolePermission{RoleId: agent.Id, PermissionId: approveAccount.Id}
	adminApproveAccount = RolePermission{RoleId: admin.Id, PermissionId: approveAccount.Id}

	basicCheckIsApproved = RolePermission{RoleId: basic.Id, PermissionId: checkIsApproved.Id}
	agentCheckIsApproved = RolePermission{RoleId: agent.Id, PermissionId: checkIsApproved.Id}
	adminCheckIsApproved = RolePermission{RoleId: admin.Id, PermissionId: checkIsApproved.Id}

	basicUpdateUserProfile    = RolePermission{RoleId: basic.Id, PermissionId: updateUserProfile.Id}
	adminUpdateUserProfile    = RolePermission{RoleId: admin.Id, PermissionId: updateUserProfile.Id}
	agentUpdateUserProfile    = RolePermission{RoleId: agent.Id, PermissionId: updateUserProfile.Id}
	verifiedUpdateUserProfile = RolePermission{RoleId: verified.Id, PermissionId: updateUserProfile.Id}

	basicUpdateUserPassword    = RolePermission{RoleId: basic.Id, PermissionId: updateUserPassword.Id}
	adminUpdateUserPassword    = RolePermission{RoleId: admin.Id, PermissionId: updateUserPassword.Id}
	agentUpdateUserPassword    = RolePermission{RoleId: agent.Id, PermissionId: updateUserPassword.Id}
	verifiedUpdateUserPassword = RolePermission{RoleId: verified.Id, PermissionId: updateUserPassword.Id}

	basicGetAllUsers         = RolePermission{RoleId: basic.Id, PermissionId: getAllUsers.Id}
	agentGetAllUsers         = RolePermission{RoleId: agent.Id, PermissionId: getAllUsers.Id}
	adminGetAllUsers         = RolePermission{RoleId: admin.Id, PermissionId: getAllUsers.Id}
	verifiedGetAllUsers      = RolePermission{RoleId: verified.Id, PermissionId: getAllUsers.Id}
	nonregisteredGetAllUsers = RolePermission{RoleId: nonregistered.Id, PermissionId: getAllUsers.Id}

	basicSearchUser         = RolePermission{RoleId: basic.Id, PermissionId: searchUser.Id}
	agentSearchUser         = RolePermission{RoleId: agent.Id, PermissionId: searchUser.Id}
	adminSearchUser         = RolePermission{RoleId: admin.Id, PermissionId: searchUser.Id}
	verifiedSearchUser      = RolePermission{RoleId: verified.Id, PermissionId: searchUser.Id}
	nonregisteredSearchUser = RolePermission{RoleId: nonregistered.Id, PermissionId: searchUser.Id}

	basicUpdatePrivacy    = RolePermission{RoleId: basic.Id, PermissionId: updatePrivacy.Id}
	agentUpdatePrivacy    = RolePermission{RoleId: agent.Id, PermissionId: updatePrivacy.Id}
	verifiedUpdatePrivacy = RolePermission{RoleId: verified.Id, PermissionId: updatePrivacy.Id}

	basicBlockUser    = RolePermission{RoleId: basic.Id, PermissionId: blockUser.Id}
	agentBlockUser    = RolePermission{RoleId: agent.Id, PermissionId: blockUser.Id}
	verifiedBlockUser = RolePermission{RoleId: verified.Id, PermissionId: blockUser.Id}

	basicUnBlockUser    = RolePermission{RoleId: basic.Id, PermissionId: unBlockUser.Id}
	agentUnBlockUser    = RolePermission{RoleId: agent.Id, PermissionId: unBlockUser.Id}
	verifiedUnBlockUser = RolePermission{RoleId: verified.Id, PermissionId: unBlockUser.Id}

	basicCheckIfBlocked         = RolePermission{RoleId: basic.Id, PermissionId: checkIfBlocked.Id}
	agentCheckIfBlocked         = RolePermission{RoleId: agent.Id, PermissionId: checkIfBlocked.Id}
	adminCheckIfBlocked         = RolePermission{RoleId: admin.Id, PermissionId: checkIfBlocked.Id}
	verifiedCheckIfBlocked      = RolePermission{RoleId: verified.Id, PermissionId: checkIfBlocked.Id}
	nonregisteredCheckIfBlocked = RolePermission{RoleId: nonregistered.Id, PermissionId: checkIfBlocked.Id}

	basicCheckUserProfilePublic         = RolePermission{RoleId: basic.Id, PermissionId: checkUserProfilePublic.Id}
	agentCheckUserProfilePublic         = RolePermission{RoleId: agent.Id, PermissionId: checkUserProfilePublic.Id}
	adminCheckUserProfilePublic         = RolePermission{RoleId: admin.Id, PermissionId: checkUserProfilePublic.Id}
	verifiedCheckUserProfilePublic      = RolePermission{RoleId: verified.Id, PermissionId: checkUserProfilePublic.Id}
	nonregisteredCheckUserProfilePublic = RolePermission{RoleId: nonregistered.Id, PermissionId: checkUserProfilePublic.Id}

	basicGetAllPublicUsers         = RolePermission{RoleId: basic.Id, PermissionId: getAllPublicUsers.Id}
	agentGetAllPublicUsers         = RolePermission{RoleId: agent.Id, PermissionId: getAllPublicUsers.Id}
	adminGetAllPublicUsers         = RolePermission{RoleId: admin.Id, PermissionId: getAllPublicUsers.Id}
	verifiedGetAllPublicUsers      = RolePermission{RoleId: verified.Id, PermissionId: getAllPublicUsers.Id}
	nonregisteredGetAllPublicUsers = RolePermission{RoleId: nonregistered.Id, PermissionId: getAllPublicUsers.Id}

	basicSubmitVerificationRequest    = RolePermission{RoleId: basic.Id, PermissionId: submitVerificationRequest.Id}
	agentSubmitVerificationRequest    = RolePermission{RoleId: agent.Id, PermissionId: submitVerificationRequest.Id}
	verifiedSubmitVerificationRequest = RolePermission{RoleId: verified.Id, PermissionId: submitVerificationRequest.Id}

	adminGetPendingVerificationRequests = RolePermission{RoleId: admin.Id, PermissionId: getPendingVerificationRequests.Id}

	adminChangeVerificationRequestStatus = RolePermission{RoleId: admin.Id, PermissionId: changeVerificationRequestStatus.Id}

	basicGetVerificationRequestsByUserId    = RolePermission{RoleId: basic.Id, PermissionId: getVerificationRequestsByUserId.Id}
	verifiedGetVerificationRequestsByUserId = RolePermission{RoleId: verified.Id, PermissionId: getVerificationRequestsByUserId.Id}
	adminGetVerificationRequestsByUserId    = RolePermission{RoleId: admin.Id, PermissionId: getVerificationRequestsByUserId.Id}
	agentGetVerificationRequestsByUserId    = RolePermission{RoleId: agent.Id, PermissionId: getVerificationRequestsByUserId.Id}

	adminGetAllVerificationRequests = RolePermission{RoleId: admin.Id, PermissionId: getAllVerificationRequests.Id}

	basicCreateNotification         = RolePermission{RoleId: basic.Id, PermissionId: createNotification.Id}
	adminCreateNotification         = RolePermission{RoleId: admin.Id, PermissionId: createNotification.Id}
	agentCreateNotification         = RolePermission{RoleId: agent.Id, PermissionId: createNotification.Id}
	verifiedCreateNotification      = RolePermission{RoleId: verified.Id, PermissionId: createNotification.Id}
	nonregisteredCreateNotification = RolePermission{RoleId: nonregistered.Id, PermissionId: createNotification.Id}

	basicUpdateUserPhoto    = RolePermission{RoleId: basic.Id, PermissionId: updateUserPhoto.Id}
	agentUpdateUserPhoto    = RolePermission{RoleId: agent.Id, PermissionId: updateUserPhoto.Id}
	adminUpdateUserPhoto    = RolePermission{RoleId: admin.Id, PermissionId: updateUserPhoto.Id}
	verifiedUpdateUserPhoto = RolePermission{RoleId: verified.Id, PermissionId: updateUserPhoto.Id}

	basicGetUserNotifications         = RolePermission{RoleId: basic.Id, PermissionId: getUserNotifications.Id}
	adminGetUserNotifications         = RolePermission{RoleId: admin.Id, PermissionId: getUserNotifications.Id}
	verifiedGetUserNotifications      = RolePermission{RoleId: verified.Id, PermissionId: getUserNotifications.Id}
	nonregisteredGetUserNotifications = RolePermission{RoleId: nonregistered.Id, PermissionId: getUserNotifications.Id}

	basicGetUserPrivacy         = RolePermission{RoleId: basic.Id, PermissionId: getUserPrivacy.Id}
	adminGetUserPrivacy         = RolePermission{RoleId: admin.Id, PermissionId: getUserPrivacy.Id}
	verifiedGetUserPrivacy      = RolePermission{RoleId: verified.Id, PermissionId: getUserPrivacy.Id}
	nonregisteredGetUserPrivacy = RolePermission{RoleId: nonregistered.Id, PermissionId: getUserPrivacy.Id}
	agentGetUserPrivacy         = RolePermission{RoleId: agent.Id, PermissionId: getUserPrivacy.Id}

	basicGetBlockedUsers      = RolePermission{RoleId: basic.Id, PermissionId: getBlockedUsers.Id}
	adminGetBlockedUsers         = RolePermission{RoleId: admin.Id, PermissionId: getBlockedUsers.Id}
	verifiedGetBlockedUsers      = RolePermission{RoleId: verified.Id, PermissionId: getBlockedUsers.Id}
	agentGetBlockedUsers        = RolePermission{RoleId: agent.Id, PermissionId: getBlockedUsers.Id}
)
