package writing

type User struct {
	Id       string `json:"id",omitempty`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
