package db

type User struct {
	ID       string `db:"ID"`
	Name     string `db:"NAME"`
	Username string `db:"USERNAME"`
	Email    string `db:"EMAIL"`
	Password string `db:"PASSWORD"`
}
