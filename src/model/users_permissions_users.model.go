package model

type UsersPermissionsUsers struct {
	Data []UsersPermissionsUsersData `json:"data"`
}

type UsersPermissionsUsersData struct {
	Id         int64                           `json:"id"`
	Attributes UsersPermissionsUsersAttributes `json:"attributes"`
}

type UsersPermissionsUsersAttributes struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Provider   string `json:"provider"`
	Confirmed  bool   `json:"confirmed"`
	Blocked    bool   `json:"blocked"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
	Firstname  string `json:"firstname"`
	Middlename string `json:"middlename"`
	Lastname   string `json:"lastname"`
	Dob        string `json:"dob"`
	Gender     string `json:"gender"`
	Telephone  string `json:"telephone"`
	Agreement  bool   `json:"agreement"`
}
