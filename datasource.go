package aegis

import (
	"github.com/thoas/go-funk"
)

type Direction string

const (
	AscDirection  Direction = "+"
	DescDirection Direction = "-"
)

type SearchOpts func(*DatasourceSearchOpts)

type DatasourceSearchOpts struct {
	Criteria
	Offset *int
	Limit  *int
	Sort   *DatasourceSortSearchOpt
}

type DatasourceSortSearchOpt struct {
	Name          *Direction
	CreatedAt     *Direction
	LastUpdatedAt *Direction
}

type Datasource interface {
	GetPolicies(resourceType, id string, opt ...SearchOpts) ([]Policy, int64, error)
}

func WithPagination(offset, limit int) SearchOpts {
	return func(opts *DatasourceSearchOpts) {
		opts.Limit = funk.PtrOf(limit).(*int)
		opts.Offset = funk.PtrOf(limit).(*int)
	}
}

func WithSort(v DatasourceSortSearchOpt) SearchOpts {
	return func(opts *DatasourceSearchOpts) {
		opts.Sort = funk.PtrOf(v).(*DatasourceSortSearchOpt)
	}
}
