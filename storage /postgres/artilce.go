package postgres

import (
	"github.com/jmoiron/sqlx"
)

type articleRepo struct {
	db *sqlx.DB
}
