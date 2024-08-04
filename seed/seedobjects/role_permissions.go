package seedobjects

type rolePermission struct {
	Role        string
	Permissions []string
}

var RolePermissionsToCreate = []rolePermission{
	// admin
	{
		Role: "admin",
		Permissions: []string{
			"sign_up_users",
			"create_roles",
			"create_menu_item",
			"update_menu_item",
			"delete_menu_item",
			"create_category",
			"update_category",
			"delete_category",
			"create_sub_category",
			"update_sub_category",
			"delete_sub_category",
		},
	},

	// manager
	{
		Role: "manager",
		Permissions: []string{
			"create_menu_item",
			"update_menu_item",
			"delete_menu_item",
			"create_category",
			"update_category",
			"delete_category",
			"create_sub_category",
			"update_sub_category",
			"delete_sub_category",
		},
	},

	// staff
}
