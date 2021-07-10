package rbac

import (
	"gorm.io/gorm"
)

func SetupContentRBAC(db *gorm.DB) error {
	dropContentTables(db)
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
			createPost, getAllPosts, getPostsForUser, removePost, getPostById, searchContentByLocation, getPostsByHashtag,
			createStory, getAllStories, getStoriesForUser, getMyStories, removeStory, getStoryById,
			createComment, getCommentsForPost,
			createLike, getLikesForPost, getDislikesForPost,
			getAllCollections, getCollection, createCollection, removeCollection, getUserFavorites, getUserFavoritesOptimized, createFavorite, removeFavorite,
			createHashtag, getAllHashtags,
			getAllHighlights, getHighlight, createHighlight, removeHighlight, createHighlightStory, removeHighlightStory, getUserLikedOrDislikedPosts,
			createContentComplaint, getAllContentComplaints, rejectById, deleteComplaintByUserId,
			getAds, getAdsFromInfluencer, createAd, incrementLinkClicks,
			getCampaigns, getCampaign, getCampaignStats, createCampaign, updateCampaign, deleteCampaign,
			getAdCategories, getAdCategory, createAdCategory, createUserAdCategories, getUsersAdCategories, updateUsersAdCategories,
			createCampaignRequest,getCampaignRequestsByAgent,
			updateCampaignRequest,

		}
		result = db.Create(&permissions)
		if result.Error != nil {
			return result.Error
		}

		rolePermissions := []RolePermission{
			basicCreatePost, agentCreatePost, verifiedCreatePost,
			basicGetAllPosts, adminGetAllPosts, agentGetAllPosts, nonregisteredGetAllPosts, verifiedGetAllPosts,
			basicGetPostsForUser, adminGetPostsForUser, agentGetPostsForUser, nonregisteredGetPostsForUser, verifiedGetPostsForUser,
			basicRemovePost, adminRemovePost, agentRemovePost, verifiedRemovePost,
			basicGetPostById, adminGetPostById, agentGetPostById, nonregisteredGetPostById, verifiedGetPostById,
			basicSearchContentByLocation, agentSearchContentByLocation, adminSearchContentByLocation, verifiedSearchContentByLocation, nonregisteredSearchContentByLocation,
			basicGetPostsByHashtag, adminGetPostsByHashtag, agentGetPostsByHashtag, nonregisteredGetPostsByHashtag, verifiedGetPostsByHashtag,
			basicCreateStory, agentCreateStory, verifiedCreateStory,
			basicGetAllStories, adminGetAllStories, agentGetAllStories, nonregisteredGetAllStories, verifiedGetAllStories,
			basicGetStoriesForUser, adminGetStoriesForUser, agentGetStoriesForUser, nonregisteredGetStoriesForUser, verifiedGetStoriesForUser,
			basicGetMyStories, agentGetMyStories, verifiedGetMyStories,
			basicRemoveStory, agentRemoveStory, adminRemoveStory, verifiedRemoveStory,
			basicGetStoryById, adminGetStoryById, agentGetStoryById, verifiedGetStoryById, nonregisteredGetStoryById,
			basicCreateComment, agentCreateComment, verifiedCreateComment,
			basicGetCommentsForPost, agentGetCommentsForPost, adminGetCommentsForPost, verifiedGetCommentsForPost, nonregisteredGetCommentsForPost,
			basicCreateLike, agentCreateLike, verifiedCreateLike,
			basicGetLikesForPost, agentGetLikesForPost, adminGetLikesForPost, nonregisteredGetLikesForPost, verifiedGetLikesForPost,
			basicGetDislikesForPost, agentGetDislikesForPost, adminGetDislikesForPost, nonregisteredGetDislikesForPost, verifiedGetDislikesForPost,
			basicGetAllCollections, agentGetAllCollections, verifiedGetAllCollections,
			basicGetUserFavoritesOptimized, agentGetUserFavoritesOptimized, verifiedGetUserFavoritesOptimized,
			basicGetCollection, agentGetCollection, verifiedGetCollection,
			basicCreateCollection, agentCreateCollection, verifiedCreateCollection,
			basicRemoveCollection, agentRemoveCollection, verifiedRemoveCollection,
			basicGetUserFavorites, agentGetUserFavorites, verifiedGetUserFavorites,
			basicCreateFavorite, agentCreateFavorite, verifiedCreateFavorite,
			basicRemoveFavorite, agentRemoveFavorite, verifiedRemoveFavorite,
			basicCreateHashtag, agentCreateHashtag, verifiedCreateHashtag,
			basicGetAllHighlights, agentGetAllHighlights, adminGetAllHighlights, verifiedGetAllHighlights, nonregisteredGetAllHighlights,
			basicGetHighlight, agentGetHighlight, adminGetHighlight, verifiedGetHighlight, nonregisteredGetHighlight,
			basicCreateHighlight, verifiedCreateHighlight, agentCreateHighlight,
			basicRemoveHighlight, agentRemoveHighlight, verifiedRemoveHighlight,
			basicCreateHighlightStory, agentCreateHighlightStory, verifiedCreateHighlightStory,
			basicRemoveHighlightStory, agentRemoveHighlightStory, verifiedRemoveHighlightStory,
			basicGetAllHashtags, adminGetAllHashtags, agentGetAllHashtags, verifiedGetAllHashtags,
			basicGetUserLikedOrDislikedPosts, verifiedGetUserLikedOrDislikedPosts,
			basicCreateContentComplaint, verifiedCreateContentComplaint, agentCreateContentComplaint,
			adminGetAllContentComplaints,
			adminRejectById,
			adminDeleteComplaintByUserId,nonregisteredDeleteComplaintByUserId,basicDeleteComplaintByUserId, verifiedDeleteComplaintByUserId, agentDeleteComplaintByUserId,
			basicGetAds, agentGetAds, nonregisteredGetAds, verifiedGetAds, adminGetAds,
			basicGetAdsFromInfluencer, agentGetAdsFromInfluencer, nonregisteredGetAdsFromInfluencer, verifiedGetAdsFromInfluencer, adminGetAdsFromInfluencer,
			agentCreateAd,
			agentCreateCampaign, agentGetCampaigns, agentGetCampaign, agentGetCampaignStats, agentUpdateCampaign, agentDeleteCampaign,
			basicGetAdCategories, nonregisteredGetAdCategories, verifiedGetAdCategories, agentGetAdCategories, adminGetAdCategories,
			basicGetAdCategory, nonregisteredGetAdCategory, verifiedGetAdCategory, agentGetAdCategory, adminGetAdCategory,
			adminCreateAdCategory,
			basicCreateUserAdCategories, nonregisteredCreateUserAdCategories, verifiedCreateUserAdCategories, agentCreateUserAdCategories, adminCreateUserAdCategories,
			basicIncrementLinkClicks, nonregisteredIncrementLinkClicks, verifiedIncrementLinkClicks, agentIncrementLinkClicks, adminIncrementLinkClicks,
			basicGetUsersAdCategories, verifiedGetUsersAdCategories, agentGetUsersAdCategories, adminGetUsersAdCategories,
			basicUpdateUsersAdCategories, verifiedUpdateUsersAdCategories, agentUpdateUsersAdCategories, adminUpdateUsersAdCategories,
			agentCreateCampaignRequest,
			agentGetCampaignRequestsByAgent,
			basicUpdateCampaignRequest, agentUpdateCampaignRequest, verifiedUpdateCampaignRequest,
		}

		result = db.Create(&rolePermissions)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	return err
}

func dropContentTables(db *gorm.DB) {
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

// Content RBAC
var (
	createPost              = Permission{Id: "a684386e-b5ab-4c58-9642-f930e5ab9937", Name: "CreatePost"}
	getAllPosts             = Permission{Id: "3d327644-43a4-4084-996e-b92a6469e353", Name: "GetAllPosts"}
	getPostsForUser         = Permission{Id: "79e74592-79e5-4853-a272-6fa9ab51426a", Name: "GetPostsForUser"}
	removePost              = Permission{Id: "172a8599-6f7d-4dbf-b9d2-8e97e48e1140", Name: "RemovePost"}
	getPostById             = Permission{Id: "d3bc0110-1445-403c-b352-3182f33ae575", Name: "GetPostById"}
	searchContentByLocation = Permission{Id: "117a5f7e-a7a7-4d37-99b7-b998d2b5d972", Name: "SearchContentByLocation"}
	getPostsByHashtag       = Permission{Id: "fede824e-ee6b-4fc8-b3fb-fbb9baace4e2", Name: "GetPostsByHashtag"}

	createStory       = Permission{Id: "9bf49450-4f25-46bf-9691-428f112868b5", Name: "CreateStory"}
	getAllStories     = Permission{Id: "7440d209-98f9-4dad-898c-8ec5daa2d71d", Name: "GetAllStories"}
	getStoriesForUser = Permission{Id: "01c4aef8-b6fa-48cd-98d8-02af401c83e2", Name: "GetStoriesForUser"}
	getMyStories      = Permission{Id: "7172c042-6d9a-4cdb-8c59-ab65427df96b", Name: "GetMyStories"}
	removeStory       = Permission{Id: "62e8ef56-5096-46e5-a57a-5e2025240d86", Name: "RemoveStory"}
	getStoryById      = Permission{Id: "9e8ba8ed-f14a-4ceb-9006-90f24f487db8", Name: "GetStoryById"}

	createComment      = Permission{Id: "4c43c1ae-fd17-4bc4-b9f9-bbab9208ad94", Name: "CreateComment"}
	getCommentsForPost = Permission{Id: "f53a6be2-bd22-43d4-a1f5-b24483343f38", Name: "GetCommentsForPost"}

	createLike                  = Permission{Id: "99553bd8-9cac-4357-be26-3a1fd1f220a2", Name: "CreateLike"}
	getLikesForPost             = Permission{Id: "9f0b1c74-0910-43e8-83f0-77436d245f3f", Name: "GetLikesForPost"}
	getDislikesForPost          = Permission{Id: "48067fbd-f078-40d0-8dee-a7f5ece74cb3", Name: "GetDislikesForPost"}
	getUserLikedOrDislikedPosts = Permission{Id: "94ec116c-92fe-4cad-b262-a566d88c2041", Name: "GetUserLikedOrDislikedPosts"}

	getAllCollections         = Permission{Id: "f7ce029b-1d08-40d6-bf16-17a2e4b26c43", Name: "GetAllCollections"}
	getUserFavoritesOptimized = Permission{Id: "f99726f8-f73b-49eb-806b-b0d45c0ae4f6", Name: "GetUserFavoritesOptimized"}
	getCollection             = Permission{Id: "1c0d7507-4e50-49cf-ae3c-9d330583acdf", Name: "GetCollection"}
	createCollection          = Permission{Id: "ebd1ebf8-07fb-4062-a5ee-cedb08a8236a", Name: "CreateCollection"}
	removeCollection          = Permission{Id: "672eb20a-26e5-42b7-a666-708b80f983ee", Name: "RemoveCollection"}
	getUserFavorites          = Permission{Id: "964d53bc-cde0-4274-9b2f-59795189550e", Name: "GetUserFavorites"}
	createFavorite            = Permission{Id: "2aff2df4-9a89-4cb8-846b-8a43a3f08c27", Name: "CreateFavorite"}
	removeFavorite            = Permission{Id: "2ff55e61-a11c-46ae-80c3-c4b5caab9da0", Name: "RemoveFavorite"}

	createHashtag = Permission{Id: "21202557-fcf5-43da-99f2-78f51b4d601e", Name: "CreateHashtag"}

	getAllHighlights     = Permission{Id: "a97df5dc-9fed-4209-97ee-cd756b1d926d", Name: "GetAllHighlights"}
	getHighlight         = Permission{Id: "4e470495-4547-4933-a336-e1051e2ecf8e", Name: "GetHighlight"}
	createHighlight      = Permission{Id: "c3a0660d-9919-4a5c-86c3-9c86a48434c4", Name: "CreateHighlight"}
	removeHighlight      = Permission{Id: "f48e1500-208f-4087-b67a-c9ef8b6ec884", Name: "RemoveHighlight"}
	createHighlightStory = Permission{Id: "fd684575-a6ac-4aa2-b4ad-d967a65cb808", Name: "CreateHighlightStory"}
	removeHighlightStory = Permission{Id: "46f3c5f7-1979-410c-96ed-d9957ef58bac", Name: "RemoveHighlightStory"}

	getAllHashtags          = Permission{Id: "1932ca5d-24af-4dcf-8f42-4ca32c799815", Name: "GetAllHashtags"}
	createContentComplaint  = Permission{Id: "334e6ffd-a9d7-4e83-bfb5-9f5f053b0069", Name: "CreateContentComplaint"}
	getAllContentComplaints = Permission{Id: "7f2315ec-909e-4856-9ccf-01c30ae76263", Name: "GetAllContentComplaints"}
	rejectById 				= Permission{Id: "fe5fecc4-dc44-11eb-ba80-0242ac130004", Name: "RejectById"}
	deleteComplaintByUserId = Permission{Id: "3dc694f6-dcf1-11eb-ba80-0242ac130004", Name: "DeleteComplaintByUserId"}

	getAds						= Permission{Id: "e1341af3-1c36-4cb8-bb26-88bd5ab2af3e", Name: "GetAds"}
	getAdsFromInfluencer		= Permission{Id: "4c8a48c2-f065-44b4-b9a5-fc7d84f6dc0d", Name: "GetAdsFromInfluencer"}
	createAd					= Permission{Id: "1cf92a55-a41a-425f-960d-8a94e7a3e997", Name: "CreateAd"}
	incrementLinkClicks 		= Permission{Id: "6f915dce-dc97-486b-b912-921407e78723", Name: "IncrementLinkClicks"}

	getCampaigns 				= Permission{Id: "cb810786-dea4-4774-940d-d4e4d693d6f4", Name: "GetCampaigns"}
	getCampaign 				= Permission{Id: "8e3a5753-3dac-4f69-a1ad-124143906070", Name: "GetCampaign"}
	getCampaignStats 			= Permission{Id: "dc0e54e9-9cf7-4195-b4d7-4de4383328c4", Name: "GetCampaignStats"}
	createCampaign 				= Permission{Id: "5b194118-4624-4725-b891-cba62eb5b506", Name: "CreateCampaign"}
	updateCampaign 				= Permission{Id: "3588ce92-6da6-4401-95e2-0af71b2e11c9", Name: "UpdateCampaign"}
	deleteCampaign 				= Permission{Id: "f54196d5-81e6-4364-86dd-45f8d864d3c1", Name: "DeleteCampaign"}

	getAdCategories 			= Permission{Id: "05cd31bd-1260-4871-ad8b-182759a95781", Name: "GetAdCategories"}
	getAdCategory 				= Permission{Id: "c3e205c5-d111-4486-847c-b3ac5e3869da", Name: "GetAdCategory"}
	createAdCategory 			= Permission{Id: "e1bb9ce5-01e9-4562-b841-ffebffef161d", Name: "CreateAdCategory"}
	createUserAdCategories 		= Permission{Id: "10b8383f-cbba-4a50-ace6-ca514a1ec5b0", Name: "CreateUserAdCategories"}
	getUsersAdCategories 		= Permission{Id: "e642da5f-fb24-4eab-8e9f-3bb9f84da712", Name: "GetUsersAdCategories"}
	updateUsersAdCategories 	= Permission{Id: "f3af637d-e25d-401d-96c8-dff295869aa3", Name: "UpdateUsersAdCategories"}

	createCampaignRequest      = Permission{Id: "666112e5-395e-4c04-a278-5d5d319bd1e8", Name: "CreateCampaignRequest"}
	getCampaignRequestsByAgent = Permission{Id: "f0d7798f-188b-4936-ba2c-0d5cf7c7539a", Name: "GetCampaignRequestsByAgent"}
	updateCampaignRequest      = Permission{Id: "797df38e-4862-47af-a87a-0194faed3c15", Name: "UpdateCampaignRequest"}
)

var (
	// Posts
	basicCreatePost    = RolePermission{RoleId: basic.Id, PermissionId: createPost.Id}
	verifiedCreatePost = RolePermission{RoleId: verified.Id, PermissionId: createPost.Id}
	agentCreatePost    = RolePermission{RoleId: agent.Id, PermissionId: createPost.Id}

	basicGetAllPosts         = RolePermission{RoleId: basic.Id, PermissionId: getAllPosts.Id}
	verifiedGetAllPosts      = RolePermission{RoleId: verified.Id, PermissionId: getAllPosts.Id}
	adminGetAllPosts         = RolePermission{RoleId: admin.Id, PermissionId: getAllPosts.Id}
	agentGetAllPosts         = RolePermission{RoleId: agent.Id, PermissionId: getAllPosts.Id}
	nonregisteredGetAllPosts = RolePermission{RoleId: nonregistered.Id, PermissionId: getAllPosts.Id}

	basicGetPostsForUser         = RolePermission{RoleId: basic.Id, PermissionId: getPostsForUser.Id}
	verifiedGetPostsForUser      = RolePermission{RoleId: verified.Id, PermissionId: getPostsForUser.Id}
	adminGetPostsForUser         = RolePermission{RoleId: admin.Id, PermissionId: getPostsForUser.Id}
	agentGetPostsForUser         = RolePermission{RoleId: agent.Id, PermissionId: getPostsForUser.Id}
	nonregisteredGetPostsForUser = RolePermission{RoleId: nonregistered.Id, PermissionId: getPostsForUser.Id}

	basicRemovePost    = RolePermission{RoleId: basic.Id, PermissionId: removePost.Id}
	verifiedRemovePost = RolePermission{RoleId: verified.Id, PermissionId: removePost.Id}
	agentRemovePost    = RolePermission{RoleId: agent.Id, PermissionId: removePost.Id}
	adminRemovePost    = RolePermission{RoleId: admin.Id, PermissionId: removePost.Id}

	basicGetPostById         = RolePermission{RoleId: basic.Id, PermissionId: getPostById.Id}
	verifiedGetPostById      = RolePermission{RoleId: verified.Id, PermissionId: getPostById.Id}
	adminGetPostById         = RolePermission{RoleId: admin.Id, PermissionId: getPostById.Id}
	agentGetPostById         = RolePermission{RoleId: agent.Id, PermissionId: getPostById.Id}
	nonregisteredGetPostById = RolePermission{RoleId: nonregistered.Id, PermissionId: getPostById.Id}

	basicSearchContentByLocation         = RolePermission{RoleId: basic.Id, PermissionId: searchContentByLocation.Id}
	verifiedSearchContentByLocation      = RolePermission{RoleId: verified.Id, PermissionId: searchContentByLocation.Id}
	adminSearchContentByLocation         = RolePermission{RoleId: admin.Id, PermissionId: searchContentByLocation.Id}
	agentSearchContentByLocation         = RolePermission{RoleId: agent.Id, PermissionId: searchContentByLocation.Id}
	nonregisteredSearchContentByLocation = RolePermission{RoleId: nonregistered.Id, PermissionId: searchContentByLocation.Id}

	basicGetPostsByHashtag         = RolePermission{RoleId: basic.Id, PermissionId: getPostsByHashtag.Id}
	verifiedGetPostsByHashtag      = RolePermission{RoleId: verified.Id, PermissionId: getPostsByHashtag.Id}
	adminGetPostsByHashtag         = RolePermission{RoleId: admin.Id, PermissionId: getPostsByHashtag.Id}
	agentGetPostsByHashtag         = RolePermission{RoleId: agent.Id, PermissionId: getPostsByHashtag.Id}
	nonregisteredGetPostsByHashtag = RolePermission{RoleId: nonregistered.Id, PermissionId: getPostsByHashtag.Id}
	// - - - - - - - - - -

	// Stories
	basicCreateStory    = RolePermission{RoleId: basic.Id, PermissionId: createStory.Id}
	verifiedCreateStory = RolePermission{RoleId: verified.Id, PermissionId: createStory.Id}
	agentCreateStory    = RolePermission{RoleId: agent.Id, PermissionId: createStory.Id}

	basicGetAllStories         = RolePermission{RoleId: basic.Id, PermissionId: getAllStories.Id}
	verifiedGetAllStories      = RolePermission{RoleId: verified.Id, PermissionId: getAllStories.Id}
	adminGetAllStories         = RolePermission{RoleId: admin.Id, PermissionId: getAllStories.Id}
	agentGetAllStories         = RolePermission{RoleId: agent.Id, PermissionId: getAllStories.Id}
	nonregisteredGetAllStories = RolePermission{RoleId: nonregistered.Id, PermissionId: getAllStories.Id}

	basicGetStoriesForUser         = RolePermission{RoleId: basic.Id, PermissionId: getStoriesForUser.Id}
	verifiedGetStoriesForUser      = RolePermission{RoleId: verified.Id, PermissionId: getStoriesForUser.Id}
	adminGetStoriesForUser         = RolePermission{RoleId: admin.Id, PermissionId: getStoriesForUser.Id}
	agentGetStoriesForUser         = RolePermission{RoleId: agent.Id, PermissionId: getStoriesForUser.Id}
	nonregisteredGetStoriesForUser = RolePermission{RoleId: nonregistered.Id, PermissionId: getStoriesForUser.Id}

	basicGetMyStories    = RolePermission{RoleId: basic.Id, PermissionId: getMyStories.Id}
	verifiedGetMyStories = RolePermission{RoleId: verified.Id, PermissionId: getMyStories.Id}
	agentGetMyStories    = RolePermission{RoleId: agent.Id, PermissionId: getMyStories.Id}

	basicRemoveStory    = RolePermission{RoleId: basic.Id, PermissionId: removeStory.Id}
	verifiedRemoveStory = RolePermission{RoleId: verified.Id, PermissionId: removeStory.Id}
	agentRemoveStory    = RolePermission{RoleId: agent.Id, PermissionId: removeStory.Id}
	adminRemoveStory    = RolePermission{RoleId: admin.Id, PermissionId: removeStory.Id}

	basicGetStoryById         = RolePermission{RoleId: basic.Id, PermissionId: getStoryById.Id}
	verifiedGetStoryById      = RolePermission{RoleId: verified.Id, PermissionId: getStoryById.Id}
	adminGetStoryById         = RolePermission{RoleId: admin.Id, PermissionId: getStoryById.Id}
	agentGetStoryById         = RolePermission{RoleId: agent.Id, PermissionId: getStoryById.Id}
	nonregisteredGetStoryById = RolePermission{RoleId: nonregistered.Id, PermissionId: getStoryById.Id}
	// - - - - - - - - - -

	// Comments
	basicCreateComment    = RolePermission{RoleId: basic.Id, PermissionId: createComment.Id}
	verifiedCreateComment = RolePermission{RoleId: verified.Id, PermissionId: createComment.Id}
	agentCreateComment    = RolePermission{RoleId: agent.Id, PermissionId: createComment.Id}

	basicGetCommentsForPost         = RolePermission{RoleId: basic.Id, PermissionId: getCommentsForPost.Id}
	verifiedGetCommentsForPost      = RolePermission{RoleId: verified.Id, PermissionId: getCommentsForPost.Id}
	adminGetCommentsForPost         = RolePermission{RoleId: admin.Id, PermissionId: getCommentsForPost.Id}
	agentGetCommentsForPost         = RolePermission{RoleId: agent.Id, PermissionId: getCommentsForPost.Id}
	nonregisteredGetCommentsForPost = RolePermission{RoleId: nonregistered.Id, PermissionId: getCommentsForPost.Id}
	// - - - - - - - - - -

	// Likes & Dislikes
	basicCreateLike    = RolePermission{RoleId: basic.Id, PermissionId: createLike.Id}
	verifiedCreateLike = RolePermission{RoleId: verified.Id, PermissionId: createLike.Id}
	agentCreateLike    = RolePermission{RoleId: agent.Id, PermissionId: createLike.Id}

	basicGetLikesForPost         = RolePermission{RoleId: basic.Id, PermissionId: getLikesForPost.Id}
	verifiedGetLikesForPost      = RolePermission{RoleId: verified.Id, PermissionId: getLikesForPost.Id}
	adminGetLikesForPost         = RolePermission{RoleId: admin.Id, PermissionId: getLikesForPost.Id}
	agentGetLikesForPost         = RolePermission{RoleId: agent.Id, PermissionId: getLikesForPost.Id}
	nonregisteredGetLikesForPost = RolePermission{RoleId: nonregistered.Id, PermissionId: getLikesForPost.Id}

	basicGetDislikesForPost         = RolePermission{RoleId: basic.Id, PermissionId: getDislikesForPost.Id}
	verifiedGetDislikesForPost      = RolePermission{RoleId: verified.Id, PermissionId: getDislikesForPost.Id}
	adminGetDislikesForPost         = RolePermission{RoleId: admin.Id, PermissionId: getDislikesForPost.Id}
	agentGetDislikesForPost         = RolePermission{RoleId: agent.Id, PermissionId: getDislikesForPost.Id}
	nonregisteredGetDislikesForPost = RolePermission{RoleId: nonregistered.Id, PermissionId: getDislikesForPost.Id}

	basicGetUserLikedOrDislikedPosts    = RolePermission{RoleId: basic.Id, PermissionId: getUserLikedOrDislikedPosts.Id}
	verifiedGetUserLikedOrDislikedPosts = RolePermission{RoleId: verified.Id, PermissionId: getUserLikedOrDislikedPosts.Id}
	// - - - - - - - - - -

	// Collections & Favorites
	basicGetAllCollections    = RolePermission{RoleId: basic.Id, PermissionId: getAllCollections.Id}
	verifiedGetAllCollections = RolePermission{RoleId: verified.Id, PermissionId: getAllCollections.Id}
	agentGetAllCollections    = RolePermission{RoleId: agent.Id, PermissionId: getAllCollections.Id}

	basicGetUserFavoritesOptimized    = RolePermission{RoleId: basic.Id, PermissionId: getUserFavoritesOptimized.Id}
	verifiedGetUserFavoritesOptimized = RolePermission{RoleId: verified.Id, PermissionId: getUserFavoritesOptimized.Id}
	agentGetUserFavoritesOptimized    = RolePermission{RoleId: agent.Id, PermissionId: getUserFavoritesOptimized.Id}

	basicGetCollection    = RolePermission{RoleId: basic.Id, PermissionId: getCollection.Id}
	verifiedGetCollection = RolePermission{RoleId: verified.Id, PermissionId: getCollection.Id}
	agentGetCollection    = RolePermission{RoleId: agent.Id, PermissionId: getCollection.Id}

	basicCreateCollection    = RolePermission{RoleId: basic.Id, PermissionId: createCollection.Id}
	verifiedCreateCollection = RolePermission{RoleId: verified.Id, PermissionId: createCollection.Id}
	agentCreateCollection    = RolePermission{RoleId: agent.Id, PermissionId: createCollection.Id}

	basicRemoveCollection    = RolePermission{RoleId: basic.Id, PermissionId: removeCollection.Id}
	verifiedRemoveCollection = RolePermission{RoleId: verified.Id, PermissionId: removeCollection.Id}
	agentRemoveCollection    = RolePermission{RoleId: agent.Id, PermissionId: removeCollection.Id}

	basicGetUserFavorites    = RolePermission{RoleId: basic.Id, PermissionId: getUserFavorites.Id}
	verifiedGetUserFavorites = RolePermission{RoleId: verified.Id, PermissionId: getUserFavorites.Id}
	agentGetUserFavorites    = RolePermission{RoleId: agent.Id, PermissionId: getUserFavorites.Id}

	basicCreateFavorite    = RolePermission{RoleId: basic.Id, PermissionId: createFavorite.Id}
	verifiedCreateFavorite = RolePermission{RoleId: verified.Id, PermissionId: createFavorite.Id}
	agentCreateFavorite    = RolePermission{RoleId: agent.Id, PermissionId: createFavorite.Id}

	basicRemoveFavorite    = RolePermission{RoleId: basic.Id, PermissionId: removeFavorite.Id}
	verifiedRemoveFavorite = RolePermission{RoleId: verified.Id, PermissionId: removeFavorite.Id}
	agentRemoveFavorite    = RolePermission{RoleId: agent.Id, PermissionId: removeFavorite.Id}
	// - - - - - - - - - -

	// Hashtags
	basicCreateHashtag    = RolePermission{RoleId: basic.Id, PermissionId: createHashtag.Id}
	verifiedCreateHashtag = RolePermission{RoleId: verified.Id, PermissionId: createHashtag.Id}
	agentCreateHashtag    = RolePermission{RoleId: agent.Id, PermissionId: createHashtag.Id}

	basicGetAllHashtags    = RolePermission{RoleId: basic.Id, PermissionId: getAllHashtags.Id}
	verifiedGetAllHashtags = RolePermission{RoleId: verified.Id, PermissionId: getAllHashtags.Id}
	agentGetAllHashtags    = RolePermission{RoleId: agent.Id, PermissionId: getAllHashtags.Id}
	adminGetAllHashtags    = RolePermission{RoleId: admin.Id, PermissionId: getAllHashtags.Id}
	// - - - - - - - - - -

	// Highlights
	basicGetAllHighlights         = RolePermission{RoleId: basic.Id, PermissionId: getAllHighlights.Id}
	verifiedGetAllHighlights      = RolePermission{RoleId: verified.Id, PermissionId: getAllHighlights.Id}
	adminGetAllHighlights         = RolePermission{RoleId: admin.Id, PermissionId: getAllHighlights.Id}
	agentGetAllHighlights         = RolePermission{RoleId: agent.Id, PermissionId: getAllHighlights.Id}
	nonregisteredGetAllHighlights = RolePermission{RoleId: nonregistered.Id, PermissionId: getAllHighlights.Id}

	basicGetHighlight         = RolePermission{RoleId: basic.Id, PermissionId: getHighlight.Id}
	verifiedGetHighlight      = RolePermission{RoleId: verified.Id, PermissionId: getHighlight.Id}
	adminGetHighlight         = RolePermission{RoleId: admin.Id, PermissionId: getHighlight.Id}
	agentGetHighlight         = RolePermission{RoleId: agent.Id, PermissionId: getHighlight.Id}
	nonregisteredGetHighlight = RolePermission{RoleId: nonregistered.Id, PermissionId: getHighlight.Id}

	basicCreateHighlight    = RolePermission{RoleId: basic.Id, PermissionId: createHighlight.Id}
	verifiedCreateHighlight = RolePermission{RoleId: verified.Id, PermissionId: createHighlight.Id}
	agentCreateHighlight    = RolePermission{RoleId: agent.Id, PermissionId: createHighlight.Id}

	basicRemoveHighlight    = RolePermission{RoleId: basic.Id, PermissionId: removeHighlight.Id}
	verifiedRemoveHighlight = RolePermission{RoleId: verified.Id, PermissionId: removeHighlight.Id}
	agentRemoveHighlight    = RolePermission{RoleId: agent.Id, PermissionId: removeHighlight.Id}

	basicCreateHighlightStory    = RolePermission{RoleId: basic.Id, PermissionId: createHighlightStory.Id}
	verifiedCreateHighlightStory = RolePermission{RoleId: verified.Id, PermissionId: createHighlightStory.Id}
	agentCreateHighlightStory    = RolePermission{RoleId: agent.Id, PermissionId: createHighlightStory.Id}

	basicRemoveHighlightStory    = RolePermission{RoleId: basic.Id, PermissionId: removeHighlightStory.Id}
	verifiedRemoveHighlightStory = RolePermission{RoleId: verified.Id, PermissionId: removeHighlightStory.Id}
	agentRemoveHighlightStory    = RolePermission{RoleId: agent.Id, PermissionId: removeHighlightStory.Id}
	// - - - - - - - - - -

	// Content Complaint
	basicCreateContentComplaint    = RolePermission{RoleId: basic.Id, PermissionId: createContentComplaint.Id}
	verifiedCreateContentComplaint = RolePermission{RoleId: verified.Id, PermissionId: createContentComplaint.Id}
	agentCreateContentComplaint    = RolePermission{RoleId: agent.Id, PermissionId: createContentComplaint.Id}

	adminGetAllContentComplaints = RolePermission{RoleId: admin.Id, PermissionId: getAllContentComplaints.Id}
	adminRejectById 			 = RolePermission{RoleId: admin.Id, PermissionId: rejectById.Id}

	adminDeleteComplaintByUserId = RolePermission{RoleId: admin.Id, PermissionId: deleteComplaintByUserId.Id}
	nonregisteredDeleteComplaintByUserId = RolePermission{RoleId: nonregistered.Id, PermissionId: deleteComplaintByUserId.Id}
	basicDeleteComplaintByUserId = RolePermission{RoleId: basic.Id, PermissionId: deleteComplaintByUserId.Id}
	verifiedDeleteComplaintByUserId = RolePermission{RoleId: verified.Id, PermissionId: deleteComplaintByUserId.Id}
	agentDeleteComplaintByUserId = RolePermission{RoleId: agent.Id, PermissionId: deleteComplaintByUserId.Id}
	// - - - - - - - - -

	// Ads
	basicGetAds         		  = RolePermission{RoleId: basic.Id, PermissionId: getAds.Id}
	verifiedGetAds			      = RolePermission{RoleId: verified.Id, PermissionId: getAds.Id}
	adminGetAds			          = RolePermission{RoleId: admin.Id, PermissionId: getAds.Id}
	agentGetAds			          = RolePermission{RoleId: agent.Id, PermissionId: getAds.Id}
	nonregisteredGetAds			  = RolePermission{RoleId: nonregistered.Id, PermissionId: getAds.Id}

	basicGetAdsFromInfluencer         		  = RolePermission{RoleId: basic.Id, PermissionId: getAdsFromInfluencer.Id}
	verifiedGetAdsFromInfluencer			  = RolePermission{RoleId: verified.Id, PermissionId: getAdsFromInfluencer.Id}
	adminGetAdsFromInfluencer			      = RolePermission{RoleId: admin.Id, PermissionId: getAdsFromInfluencer.Id}
	agentGetAdsFromInfluencer			      = RolePermission{RoleId: agent.Id, PermissionId: getAdsFromInfluencer.Id}
	nonregisteredGetAdsFromInfluencer		  = RolePermission{RoleId: nonregistered.Id, PermissionId: getAdsFromInfluencer.Id}

	agentCreateAd 				  = RolePermission{RoleId: agent.Id, PermissionId: createAd.Id}

	basicIncrementLinkClicks         		  = RolePermission{RoleId: basic.Id, PermissionId: incrementLinkClicks.Id}
	verifiedIncrementLinkClicks			      = RolePermission{RoleId: verified.Id, PermissionId: incrementLinkClicks.Id}
	adminIncrementLinkClicks			      = RolePermission{RoleId: admin.Id, PermissionId: incrementLinkClicks.Id}
	agentIncrementLinkClicks			      = RolePermission{RoleId: agent.Id, PermissionId: incrementLinkClicks.Id}
	nonregisteredIncrementLinkClicks		  = RolePermission{RoleId: nonregistered.Id, PermissionId: incrementLinkClicks.Id}
	// - - - - - - - - -

	// Campaigns
	agentGetCampaign			  = RolePermission{RoleId: agent.Id, PermissionId: getCampaign.Id}
	agentGetCampaignStats		  = RolePermission{RoleId: agent.Id, PermissionId: getCampaignStats.Id}
	agentGetCampaigns			  = RolePermission{RoleId: agent.Id, PermissionId: getCampaigns.Id}
	agentCreateCampaign			  = RolePermission{RoleId: agent.Id, PermissionId: createCampaign.Id}
	agentUpdateCampaign			  = RolePermission{RoleId: agent.Id, PermissionId: updateCampaign.Id}
	agentDeleteCampaign			  = RolePermission{RoleId: agent.Id, PermissionId: deleteCampaign.Id}
	// - - - - - - - - -

	// Ad Categories
	basicGetAdCategories          = RolePermission{RoleId: basic.Id, PermissionId: getAdCategories.Id}
	verifiedGetAdCategories	      = RolePermission{RoleId: verified.Id, PermissionId: getAdCategories.Id}
	adminGetAdCategories		  = RolePermission{RoleId: admin.Id, PermissionId: getAdCategories.Id}
	agentGetAdCategories		  = RolePermission{RoleId: agent.Id, PermissionId: getAdCategories.Id}
	nonregisteredGetAdCategories  = RolePermission{RoleId: nonregistered.Id, PermissionId: getAdCategories.Id}

	basicGetAdCategory          = RolePermission{RoleId: basic.Id, PermissionId: getAdCategory.Id}
	verifiedGetAdCategory	    = RolePermission{RoleId: verified.Id, PermissionId: getAdCategory.Id}
	adminGetAdCategory		    = RolePermission{RoleId: admin.Id, PermissionId: getAdCategory.Id}
	agentGetAdCategory		    = RolePermission{RoleId: agent.Id, PermissionId: getAdCategory.Id}
	nonregisteredGetAdCategory  = RolePermission{RoleId: nonregistered.Id, PermissionId: getAdCategory.Id}

	basicCreateUserAdCategories          = RolePermission{RoleId: basic.Id, PermissionId: createUserAdCategories.Id}
	verifiedCreateUserAdCategories	    = RolePermission{RoleId: verified.Id, PermissionId: createUserAdCategories.Id}
	adminCreateUserAdCategories		    = RolePermission{RoleId: admin.Id, PermissionId: createUserAdCategories.Id}
	agentCreateUserAdCategories		    = RolePermission{RoleId: agent.Id, PermissionId: createUserAdCategories.Id}
	nonregisteredCreateUserAdCategories  = RolePermission{RoleId: nonregistered.Id, PermissionId: createUserAdCategories.Id}

	basicGetUsersAdCategories          = RolePermission{RoleId: basic.Id, PermissionId: getUsersAdCategories.Id}
	verifiedGetUsersAdCategories	      = RolePermission{RoleId: verified.Id, PermissionId: getUsersAdCategories.Id}
	adminGetUsersAdCategories		  = RolePermission{RoleId: admin.Id, PermissionId: getUsersAdCategories.Id}
	agentGetUsersAdCategories		  = RolePermission{RoleId: agent.Id, PermissionId: getUsersAdCategories.Id}

	basicUpdateUsersAdCategories          = RolePermission{RoleId: basic.Id, PermissionId: updateUsersAdCategories.Id}
	verifiedUpdateUsersAdCategories	      = RolePermission{RoleId: verified.Id, PermissionId: updateUsersAdCategories.Id}
	adminUpdateUsersAdCategories		  = RolePermission{RoleId: admin.Id, PermissionId: updateUsersAdCategories.Id}
	agentUpdateUsersAdCategories		  = RolePermission{RoleId: agent.Id, PermissionId: updateUsersAdCategories.Id}

	adminCreateAdCategory 	    = RolePermission{RoleId: agent.Id, PermissionId: createAdCategory.Id}
	// - - - - - - - - -

	agentCreateCampaignRequest = RolePermission{RoleId: agent.Id, PermissionId: createCampaignRequest.Id}


	agentGetCampaignRequestsByAgent = RolePermission{RoleId: agent.Id, PermissionId: getCampaignRequestsByAgent.Id}

	basicUpdateCampaignRequest    = RolePermission{RoleId: basic.Id, PermissionId: updateCampaignRequest.Id}
	agentUpdateCampaignRequest    = RolePermission{RoleId: agent.Id, PermissionId: updateCampaignRequest.Id}
	verifiedUpdateCampaignRequest = RolePermission{RoleId: verified.Id, PermissionId: updateCampaignRequest.Id}
)
