package postgres

import (
	"database/sql"

	"sync"

	"github.com/JoaoEymard/ingressou/api/utils/settings"
	// _ Importanto apenas o init
	_ "github.com/lib/pq"
)

var (
	postgres *sql.DB
	mutex    sync.Mutex
)

// Open realizar a conexão com o banco de dados postgres.
func Open() error {
	var (
		err error
	)

	mutex.Lock()
	criptPass := settings.GetSettings().Database.ConnectionRw.Pass
	postgres, err = sql.Open("postgres", "host="+settings.GetSettings().Database.ConnectionRw.Host+" user="+settings.GetSettings().Database.ConnectionRw.User+" password="+string(criptPass[0:2])+string(criptPass[len(criptPass)-2:])+" dbname="+settings.GetSettings().Database.ConnectionRw.Database+" sslmode=disable")
	mutex.Unlock()
	if err != nil {
		return err
	}

	if err := postgres.Ping(); err != nil {
		return err
	}
	postgres.SetMaxOpenConns(settings.GetSettings().Database.ConnectionRw.MaxOpenConn)
	return nil
}

// Close t
func Close() {
	// conexao.Commit()
	postgres.Close()
}

// GetPostgres Retornando o postgres
func GetPostgres() *sql.DB {
	return postgres
}
