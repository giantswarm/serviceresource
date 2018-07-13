package service

import (
	"context"
)

func (r *Resource) GetDesiredState(ctx context.Context, obj interface{}) (interface{}, error) {
	return r.desiredServicesFunc(ctx, obj)
}
