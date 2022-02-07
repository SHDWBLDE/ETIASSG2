package db

import (
	"context"
	"fmt"
	"log"

	"github.com/MalukiMuthusi/edufi/server/graph/model"
)

// CreateModule saves a new module to the database
func (P *Postgres) CreateModule(ctx context.Context, input model.NewModule) (*string, error) {
	sqlStmt := "INSERT INTO modules(name,synopsis) VALUES( $1,$2) RETURNING id"

	rows := Db.QueryRowContext(ctx, sqlStmt, input.Name, input.Synopsis)

	var i int64 = -1

	err := rows.Scan(&i)
	if err != nil {
		log.Printf("CreateModule: failed to create module: %v", err)
		return nil, fmt.Errorf("module not created")
	}

	id, err := P.createID("modules", i)
	if err != nil {
		log.Printf("CreateModule: failed to create from the table ID: %v", err)
		return nil, err
	}

	return id, nil
}
