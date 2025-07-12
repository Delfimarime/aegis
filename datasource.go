package aegis

import "github.com/thoas/go-funk"

type SearchOpts func(*DatasourceSearchOpts)

type DatasourceSearchOpts struct {
	Criteria
	Offset *int
	Limit  *int
	Sort   []DatasourceSortSearchOpt
}

type DatasourceSortSearchOpt struct {
	Property  string
	Direction string
}

type Datasource interface {
	GetPolicies(resourceType, id string, opt ...SearchOpts) ([]Policy, int64, error)
}

func WithPagination(offset, limit int) SearchOpts {
	return func(opts *DatasourceSearchOpts) {
		opts.Limit = funk.PtrOf(limit)
		opts.Offset = funk.PtrOf(limit)
	}
}
