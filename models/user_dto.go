package models

// User model
type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

var users map[string]User = map[string]User{
	"2fa0b8ef-cfeb-4055-8270-299cfdef8934": {
		ID:        "2fa0b8ef-cfeb-4055-8270-299cfdef8934",
		FirstName: "Oscar",
		LastName:  "Oram",
		Email:     "oscarm-b@tutamail.com",
	},
}
