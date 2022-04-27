package db

type Blog struct {
	ID            int    `db:"id"`
	PublishedDate string `db:"published_date"`
	Title         string `db:"title"`
	Content       string `db:"content"`
	Public        bool   `db:"public"`
	Groups        string `db:"groups"`
	User_ID       string `db:"user_id"`
}
