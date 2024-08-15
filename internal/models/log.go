package models

type Log struct {
	ID          string `db:"id"`
	UserAgent   string `db:"user_agent"`
	RequestId   string `db:"request_id"`
	RandomValue string `db:"random_value"`
	Url         string `db:"url"`
	Count       int    `db:"count"`
}
