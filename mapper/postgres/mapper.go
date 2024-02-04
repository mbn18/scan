package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mbn18/scan/entity"
)

// @TODO should we convert NULL data to primitives zero value? depend on the business logic.
// @TODO should we use upper to allow case insensitivity on kind (type) search
const sqlListByKind = "SELECT urn, type, name, data, COALESCE(generated_at, '1970-01-01'::timestamp) FROM resource WHERE upper(type)=upper($1)"

type Mapper struct {
	pool *pgxpool.Pool
}

func (m Mapper) ListByKind(ctx context.Context, kind string) ([]*entity.Resource, error) {
	rows, err := m.pool.Query(ctx, sqlListByKind, kind)
	if err != nil {
		return nil, err
	}

	list := make([]*entity.Resource, 0)
	for rows.Next() {
		r := new(entity.Resource)
		err = rows.Scan(&r.Urn, &r.Kind, &r.Name, &r.Data, &r.GeneratedAt)
		if err != nil {
			return nil, err
		}
		list = append(list, r)
	}

	return list, err
}

func NewMapper(conn *pgxpool.Pool) *Mapper {
	return &Mapper{pool: conn}
}
