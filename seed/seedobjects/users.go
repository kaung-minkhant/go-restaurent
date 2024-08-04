package seedobjects

type user struct {
	Name       string `json:"name"`
	Password   string `json:"encrypted_password"`
	Role       string `json:"role"`
	EmployeeID string `json:"employee_id"`
}

var UsersToCreate = []user{
	{
		Name:       "admin",
		Password:   "admin",
		Role:       "admin",
		EmployeeID: "admin",
	},
	{
		Name:       "shun le",
		Password:   "shunn",
		Role:       "manager",
		EmployeeID: "shunnleyee",
	},
	{
		Name:       "shun le staff",
		Password:   "shunnstaff",
		Role:       "staff",
		EmployeeID: "shunnleyeestaff",
	},
}
