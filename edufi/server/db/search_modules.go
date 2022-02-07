package db

import (
	"context"
	"fmt"
	"log"

	"github.com/MalukiMuthusi/edufi/server/graph/model"
	"github.com/MalukiMuthusi/edufi/server/util"
)

// SearchModules searches and returns any modules that have the text string in their name
func (P *Postgres) SearchModules(ctx context.Context, text string) ([]*model.Module, error) {
	sql := "SELECT id,name,synopsis FROM modules WHERE name_searchable @@ websearch_to_tsquery($1)"

	rows, err := Db.QueryContext(ctx, sql, text)
	if err != nil {
		log.Printf("SearchModules: %v", rows)
		return nil, fmt.Errorf("no module name matches")
	}

	var modules []*model.Module
	for rows.Next() {
		var (
			id       int64
			name     string
			synopsis string
		)

		err = rows.Scan(&id, &name, &synopsis)
		if err != nil {
			log.Printf("SearchModules: %v", err)
			return nil, fmt.Errorf("no module found")
		}

		idString, err := P.createID(util.ModuleTable, id)
		if err != nil {
			log.Printf("ListModules: %v", err)
			return nil, err
		}

		module := model.Module{
			ID:       *idString,
			Name:     name,
			Synopsis: synopsis,
		}

		modules = append(modules, &module)

	}
	return modules, nil
}
