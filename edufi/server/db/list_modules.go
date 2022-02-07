package db

import (
	"context"
	"log"

	"github.com/MalukiMuthusi/edufi/server/graph/model"
	"github.com/MalukiMuthusi/edufi/server/util"
)

// ListModules returns all the modules saved in the database
func (P *Postgres) ListModules(ctx context.Context) ([]*model.Module, error) {
	sqlStmt := "SELECT id,name,synopsis FROM modules"

	rows, err := Db.QueryContext(ctx, sqlStmt)
	if err != nil {
		log.Printf("ListModules: %v", err)
		return nil, err
	}
	defer rows.Close()

	var modules []*model.Module
	for rows.Next() {
		var (
			id       int64
			name     string
			synopsis string
		)
		err = rows.Scan(&id, &name, &synopsis)
		if err != nil {
			log.Printf("ListModules: %v", err)
			return nil, err
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
