package models

type Log struct {
	TokenID   string `db:"token_id"`
	Token     string `db:"token"`
	UserAgent string `db:"user_agent"`
	Url       string `db:"url"`
	Count     int    `db:"count"`
}
