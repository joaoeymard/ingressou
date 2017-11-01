package postgres

import (
	// _ Importanto apenas o init
	_ "github.com/lib/pq"
)

// ExecuteQuery recebir query para serem executadas no banco postgres
func ExecuteQuery(query string) error {
	res, err := postgres.Exec(query)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}
