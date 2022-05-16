package model

type Method interface {
	// sql(SELECT version FROM users GROUP BY version ORDER BY version DESC LIMIT 1)
	GetMaxVersionCount() (int, error)

}
