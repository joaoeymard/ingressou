package postgres

import (
	"fmt"
	"strings"

	"github.com/JoaoEymard/ingressou/api/v1/utils"
	// _ Importanto apenas o init
	_ "github.com/lib/pq"
)

// InsertOne Insere um registro no Banco Postgres
func InsertOne(tabela string, params map[string]interface{}) (map[string]interface{}, error) {
	var (
		key, queryValue []string
		value           []interface{}
		dados           map[string]interface{}
	)

	count := 0

	for k, v := range params {
		// Setando as colunas da tabela
		key = append(key, k)
		// Setando os valores que serão inseridos
		value = append(value, v)

		count++
		queryValue = append(queryValue, fmt.Sprintf("$%d", count))
	}

	query := fmt.Sprintf(`INSERT INTO %v
	( %v )
	VALUES( %v )
	RETURNING id;`, tabela, strings.Join(key, ", "), strings.Join(queryValue, ", "))

	stmt, err := postgres.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(value...)
	if err != nil {
		return nil, err
	}

	columns, _ := rows.Columns()

	for rows.Next() {

		var (
			rowsValues = make(map[string]interface{}, len(columns))
			refs       = make([]interface{}, 0, len(columns))
		)

		for _, column := range columns {
			var ref interface{}
			rowsValues[column] = &ref
			refs = append(refs, &ref)
		}

		rows.Scan(refs...)

		dados = rowsValues

	}

	if dados == nil {
		return nil, utils.BancoDadosMethod("inserir")
	}

	return dados, nil
}
