package cidades

import (
	sql "github.com/JoaoEymard/ingressou/service/core/database"
	"gopkg.in/kataras/iris.v6"
)

// FindAll Retorna todos os registro de cidades
func FindAll() ([]iris.Map, error) {

	values, err := sql.ExecuteSQL(`SELECT id, uf, name, has_fiber, has_radio, has_tv, has_phone, meets
	FROM public.s_cidades;`)

	if err != nil {
		return nil, err
	}

	return values, nil

}
