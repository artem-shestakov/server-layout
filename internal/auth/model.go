package auth

type Role struct {
	Id   int
	Name string
}

// User model
type User struct {
	Id                 int    `json:"id" db:"id"`
	Firstname          string `json:"firstname" db:"firstname"`
	Lastname           string `json:"lastname" db:"lastname"`
	Email              string `json:"email" db:"email"`
	Role               string `json:"role" db:"role"`
	Password           string `json:"-" db:"password"`
	CreatedAt          string `json:"created_at" db:"created_at"`
	LastLogin          string `json:"last_login" db:"last_login"`
	LastPasswordChange string `json:"last_password_change" db:"last_password_change"`
	IsActive           bool   `json:"is_active" db:"is_active"`
}

type CreateUser struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Role      string `json:"-"`
	Password  string `json:"password"`
}

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWTResponse struct {
	Token string `json:"token"`
}
