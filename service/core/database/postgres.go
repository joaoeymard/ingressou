package postgres

import (
	"database/sql"

	"sync"

	"fmt"

	_ "github.com/lib/pq"

	"errors"

	utils "github.com/JoaoEymard/ingressou/service/utils"
	"gopkg.in/kataras/iris.v6"
)

var (
	postgres *sql.DB
	mutex    sync.Mutex
)

// Open realizar a conexão com o banco de dados postgres.
func Open() error {
	return errors.New("Teste")

	var (
		err error
	)

	mutex.Lock()
	postgres, err = sql.Open("postgres", "host=192.168.111.11 port=5435 user=postgres password=1234567890abcdefghij dbname=site_brisa sslmode=disable")
	mutex.Unlock()

	if err != nil {
		return err
	}
	if err := postgres.Ping(); err != nil {
		return err
	}
	postgres.SetMaxOpenConns(utils.Get().Database.ConnectionRo.MaxIdleConn)
	return nil
}

// Close t
func Close() {
	// conexao.Commit()
	postgres.Close()
}

// ExecuteSQL recebir query para serem executadas no banco postgres
func ExecuteSQL(sql string) ([]iris.Map, error) {

	var dados []iris.Map

	rows, err := postgres.Query(sql)

	if err != nil {
		return nil, err
	}

	columns, err := rows.Columns()
	if err != nil {
		fmt.Println("[Erro]", err)
		return nil, err
	}

	values := make([]interface{}, len(columns))
	scanValue := make([]interface{}, len(columns))
	for i := range columns {
		scanValue[i] = &values[i]
	}

	for rows.Next() {
		err := rows.Scan(scanValue...)
		if err != nil {
			return nil, err
		}

		var dado = make(iris.Map)
		for i, v := range columns {
			dado[v] = values[i]
		}

		dados = append(dados, dado)
	}

	rows.Close()

	return dados, nil
}
