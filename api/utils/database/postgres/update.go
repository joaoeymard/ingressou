package postgres

import (
	"fmt"
	"strings"

	"github.com/JoaoEymard/ingressoscariri/api/v1/utils"
	// _ Importanto apenas o init
	_ "github.com/lib/pq"
)

// UpdateOne Insere um registro no Banco Postgres
func UpdateOne(tabela string, params map[string]interface{}, where string) (map[string]interface{}, error) {
	var (
		queryKeyValue []string
		value         []interface{}
		dados         map[string]interface{}
	)

	count := 0

	for k, v := range params {

		if v != nil {
			count++
			// Setando as colunas da tabela
			queryKeyValue = append(queryKeyValue, fmt.Sprintf("%v = $%d", k, count))

			// Setando os valores que serão inseridos
			value = append(value, v)
		}
	}

	query := fmt.Sprintf(`UPDATE %v
	SET %v
	WHERE %v
	RETURNING id;`, tabela, strings.Join(queryKeyValue, ", "), where)

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
		return nil, utils.BancoDadosMethod("atualizar")
	}

	return dados, nil
}
