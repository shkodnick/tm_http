package store

type SignUpParams struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
}
