package service

import (
	"context"
	"github.com/blocktock/kit-mdp/model"
)

// ReflectionService describes the service.
type ReflectionService interface {	GetSpec(ctx context.Context, id string) (status bool, errinfo string, data *model.Reflection)
	GetAll(ctx context.Context, page, offset, limit int, orderby string) (status bool, errinfo string, data []model.Reflection)
	GetUDF(ctx context.Context, name string, page, offset, limit int, orderby string) (status bool, errinfo string, data []model.Reflection)
	Post(ctx context.Context, Reflection model.Reflection) (status bool, errinfo string, data *model.Reflection)
	Put(ctx context.Context, id string, Reflection model.Reflection) (status bool, errinfo string, data *model.Reflection)
	Delete(ctx context.Context, id string) (status bool, errinfo string, data *model.Reflection)
}
