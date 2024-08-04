package seedobjects

type permission struct {
	Permission string
	Method     string
	Route      string
}

var PermissionsToCreate = []permission{
	// menu items
	{Permission: "create_menu_item", Method: "POST", Route: "/menu_items"},
	{Permission: "update_menu_item", Method: "PATCH", Route: "/menu_items/{menuItemId}"},
	{Permission: "delete_menu_item", Method: "DELETE", Route: "/menu_items/{menuItemId}"},

	// category
	{Permission: "create_category", Method: "POST", Route: "/categories"},
	{Permission: "update_category", Method: "PATCH", Route: "/categories/{categoryId}"},
	{Permission: "delete_category", Method: "DELETE", Route: "/categories/{categoryId}"},

	// subcategory
	{Permission: "create_sub_category", Method: "POST", Route: "/sub_categories"},
	{Permission: "update_sub_category", Method: "PATCH", Route: "/sub_categories/{subCategoryId}"},
	{Permission: "delete_sub_category", Method: "DELETE", Route: "/sub_categories/{subCategoryId}"},

	// users
	{Permission: "sign_up_users", Method: "POST", Route: "/auth/signup"},

	// roles
	{Permission: "create_roles", Method: "POST", Route: "/roles"},

	// permissions

	// role permissions
}
