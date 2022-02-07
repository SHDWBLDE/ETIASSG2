package store

import (
	"context"

	"github.com/MalukiMuthusi/edufi/server/graph/model"
)

type Store interface {
	CreateModule(ctx context.Context, input model.NewModule) (*string, error)
	ListModules(ctx context.Context) ([]*model.Module, error)
	GetModule(ctx context.Context, id string) (*model.Module, error)
	DeleteModule(ctx context.Context, id string) (*string, error)
	UpdateModule(ctx context.Context, id string, input model.NewModule) (*model.Module, error)
	SearchModules(ctx context.Context, text string) ([]*model.Module, error)
}

func init() {}
