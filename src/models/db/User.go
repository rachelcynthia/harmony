package db

type User struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password"`
}
