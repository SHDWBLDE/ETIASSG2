package db

import (
	"context"
	"fmt"

	"github.com/MalukiMuthusi/edufi/server/graph/model"
	"github.com/MalukiMuthusi/edufi/server/util"
)

// GetModule returns the module specified by the id
func (P *Postgres) GetModule(ctx context.Context, id string) (*model.Module, error) {
	sql := "SELECT id, name, synopsis FROM modules WHERE id = $1"
	idString, err := P.getID(id, util.ModuleTable)
	if err != nil {
		return nil, fmt.Errorf("provided id not found")
	}

	row := Db.QueryRowContext(ctx, sql, idString)
	var (
		idReturned int64
		name       string
		synopsis   string
	)

	err = row.Scan(&idReturned, &name, &synopsis)
	if err != nil {
		return nil, fmt.Errorf("module not found")
	}

	newID, err := P.createID(util.ModuleTable, idReturned)
	if err != nil {
		return nil, err
	}

	module := model.Module{
		ID:       *newID,
		Name:     name,
		Synopsis: synopsis,
	}

	return &module, nil
}
