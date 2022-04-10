package users

type User struct {
	ID             string `json:"id" bson:"_id, omitempty"`
	Email          string `json:"email" bson:"email"`
	Name           string `json:"name" bson:"name, omitempty"`
	HashedPassword string `json:"-" bson:"password"`
}

type UserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
