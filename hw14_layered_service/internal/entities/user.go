package entities

type User struct {
	Login string `json:"login"`
}

type UserWithPassword struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
