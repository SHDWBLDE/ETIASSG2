package db

import (
	"context"
	"fmt"
	"log"

	"github.com/MalukiMuthusi/edufi/server/graph/model"
	"github.com/MalukiMuthusi/edufi/server/util"
)

// UpdateModule edits a module, saves and returns the edited module
func (P *Postgres) UpdateModule(ctx context.Context, id string, input model.NewModule) (*model.Module, error) {
	sql := "UPDATE modules SET name = $1, synopsis = $2 WHERE id = $3"

	dbID, err := P.getID(id, util.ModuleTable)
	if err != nil {
		return nil, fmt.Errorf("module not found with the provided ID")
	}

	res, err := Db.ExecContext(ctx, sql, input.Name, input.Synopsis, dbID)
	if err != nil {
		log.Printf("UpdateModule: %v", err)
		return nil, fmt.Errorf("module not updated")
	}

	i, err := res.RowsAffected()
	if err != nil {
		log.Printf("UpdateModule: %v", err)
		return nil, fmt.Errorf("module not updated")
	}

	if i != 1 {
		return nil, fmt.Errorf("system error when updating module")
	}

	module, err := P.GetModule(ctx, id)
	if err != nil {
		return nil, err
	}

	return module, nil
}
