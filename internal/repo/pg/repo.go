package pg

import (
	"database/sql"
)

type PgRepo struct {
	conn *sql.DB
}

func New(conn *sql.DB) *PgRepo {
	return &PgRepo{conn: conn}
}

func (p *PgRepo) handle_AS_ADDR_OBJ(ch chan map[string]string) {

}

func (p *PgRepo) handle_AS_ADDR_OBJ_DIVISION() {

}

func (p *PgRepo) handle_AS_ADDR_OBJ_PARAMS() {

}

func (p *PgRepo) handle_AS_ADM_HIERARCHY() {

}

func (p *PgRepo) handle_AS_APARTMENTS_PARAMS() {

}

func (p *PgRepo) handle_AS_APARTMENTS() {

}

func (p *PgRepo) handle_AS_CARPLACES() {

}

func (p *PgRepo) handle_AS_CARPLACES_PARAMS() {

}

func (p *PgRepo) handle_AS_CHANGE_HISTORY() {

}

func (p *PgRepo) handle_AS_HOUSES() {

}

func (p *PgRepo) handle_AS_HOUSES_PARAMS() {

}

func (p *PgRepo) handle_AS_MUN_HIERARCHY() {

}

func (p *PgRepo) handle_AS_NORMATIVE_DOCS() {

}

func (p *PgRepo) handle_AS_REESTR_OBJECTS() {

}

func (p *PgRepo) handle_AS_ROOMS() {

}

func (p *PgRepo) handle_AS_ROOMS_PARAMS() {

}

func (p *PgRepo) handle_AS_STEADS() {

}

func (p *PgRepo) handle_AS_STEADS_PARAMS() {

}
