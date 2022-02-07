package db

import (
	"context"
	"fmt"

	"github.com/MalukiMuthusi/edufi/server/util"
)

// DeleteModule deletes the module with the specified id
func (P *Postgres) DeleteModule(ctx context.Context, id string) (*string, error) {
	sql := "DELETE FROM modules WHERE id = $1"

	idInt, err := P.getID(id, util.ModuleTable)
	if err != nil {
		return nil, fmt.Errorf("module not deleted, id provided is invalid")
	}

	res, err := Db.ExecContext(ctx, sql, *idInt)
	if err != nil {
		return nil, fmt.Errorf("module not deleted, error in the query")
	}

	i, err := res.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("module not deleted, rows affected not returned")
	}

	if i != 1 {
		return nil, fmt.Errorf("module not deleted, zero rows affected")
	}

	return &id, nil
}
