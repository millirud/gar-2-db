package pg

import "database/sql"

type PgRepo struct {
	conn *sql.DB
}

func New(conn *sql.DB) *PgRepo {
	return &PgRepo{conn: conn}
}
