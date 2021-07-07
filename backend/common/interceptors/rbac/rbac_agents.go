package rbac

import "gorm.io/gorm"

func SetupAgentsRBAC(db *gorm.DB) error {
	dropAgentTables(db)
	err := db.AutoMigrate(&UserRole{}, Permission{}, RolePermission{})
	if err != nil {
		return err
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		userRoles := []UserRole{basic, agent, nonregistered}
		result := db.Create(&userRoles)
		if result.Error != nil {
			return result.Error
		}

		permissions := []Permission{
			createProduct, getUserByUsernameAgent, getAllProductsByAgentId, getAllProducts,
			getProductById, deleteProduct, updateProduct, orderProduct, getOrdersByUser, getOrdersByAgent,
		}
		result = db.Create(&permissions)
		if result.Error != nil {
			return result.Error
		}

		rolePermissions := []RolePermission{
			agentCreateProduct,
			agentGetUserByUsernameAgent, basicGetUserByUsernameAgent,
			basicGetAllProductsByAgentId, agentGetAllProductsByAgentId,
			basicGetAllProducts, agentGetAllProducts,
			basicGetProductById, agentGetProductById,
			agentDeleteProduct,
			agentUpdateProduct,
			basicOrderProduct,
			basicGetOrdersByUser,
			agentGetOrdersByAgent,
		}
		result = db.Create(&rolePermissions)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	return err
}

func dropAgentTables(db *gorm.DB) {
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
	createProduct           = Permission{Id: "dc1e031f-c276-42de-87c6-786e358fc51e", Name: "CreateProduct"}
	getUserByUsernameAgent  = Permission{Id: "98ec230f-7995-4970-87be-a71e2efbcbd2", Name: "GetUserByUsername"}
	getAllProductsByAgentId = Permission{Id: "9f0095b3-0c0a-42c8-908a-dc05757a1453", Name: "GetAllProductsByAgentId"}
	getAllProducts          = Permission{Id: "d149a012-4b67-4525-82cb-66e9b2fccd54", Name: "GetAllProducts"}
	getProductById          = Permission{Id: "5ce5d04c-c936-4ace-bada-7ebb89255ffb", Name: "GetProductById"}
	deleteProduct           = Permission{Id: "e447ab3f-2a85-4fd2-a025-07c117674de5", Name: "DeleteProduct"}
	updateProduct           = Permission{Id: "331282b9-1572-4919-b506-94aaf8091994", Name: "UpdateProduct"}
	orderProduct            = Permission{Id: "466b2260-fafd-4384-a304-d60bb77838a3", Name: "OrderProduct"}
	getOrdersByUser         = Permission{Id: "e8e64a61-a63d-4cad-a039-ef417603b91d", Name: "GetOrdersByUser"}
	getOrdersByAgent        = Permission{Id: "ee1a39fc-a246-41fc-976b-5b12e2a1319a", Name: "GetOrdersByAgent"}
)

var (
	agentCreateProduct = RolePermission{RoleId: agent.Id, PermissionId: createProduct.Id}

	agentGetUserByUsernameAgent = RolePermission{RoleId: agent.Id, PermissionId: getUserByUsernameAgent.Id}
	basicGetUserByUsernameAgent = RolePermission{RoleId: basic.Id, PermissionId: getUserByUsernameAgent.Id}

	agentGetAllProductsByAgentId = RolePermission{RoleId: agent.Id, PermissionId: getAllProductsByAgentId.Id}
	basicGetAllProductsByAgentId = RolePermission{RoleId: basic.Id, PermissionId: getAllProductsByAgentId.Id}

	basicGetAllProducts = RolePermission{RoleId: basic.Id, PermissionId: getAllProducts.Id}
	agentGetAllProducts = RolePermission{RoleId: agent.Id, PermissionId: getAllProducts.Id}

	basicGetProductById = RolePermission{RoleId: basic.Id, PermissionId: getProductById.Id}
	agentGetProductById = RolePermission{RoleId: agent.Id, PermissionId: getProductById.Id}

	agentDeleteProduct = RolePermission{RoleId: agent.Id, PermissionId: deleteProduct.Id}

	agentUpdateProduct = RolePermission{RoleId: agent.Id, PermissionId: updateProduct.Id}

	basicOrderProduct = RolePermission{RoleId: basic.Id, PermissionId: orderProduct.Id}

	basicGetOrdersByUser = RolePermission{RoleId: basic.Id, PermissionId: getOrdersByUser.Id}

	agentGetOrdersByAgent = RolePermission{RoleId: agent.Id, PermissionId: getOrdersByAgent.Id}
)
