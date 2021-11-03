package query_map

import (
	"errors"
	"strings"

	"github.com/jmoiron/sqlx"
)

type DB struct {
	*sqlx.DB
}

func SqlInjection(sql string) error {
	injection := []string{
		"insert",
		"update",
		"delete",
		"truncate",
		"alter",
		"drop",
		"exec",
		"create",
		"grant",
		"revolk",
	}
	for _, v := range injection {
		if strings.Contains(sql, v) {
			response := "error : prohibited command"
			println(response)
			return errors.New("error : sql contains prohibit command")
		}
	}
	return nil
}

func (db *DB) SqlQuery(sql string) (string, error) {
	err := SqlInjection(sql)
	if err != nil {
		return "", err
	}
	db.DB.Ping()
	return "pss", nil
}
