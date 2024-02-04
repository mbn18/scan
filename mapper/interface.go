package mapper

import (
	"context"
	"github.com/mbn18/scan/entity"
)

type Mapper interface {
	ListByKind(ctx context.Context, kind string) ([]*entity.Resource, error)
}
