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
	Select *DatasourceAttributeSelectionSearchOpt
}

type DatasourceSortSearchOpt struct {
	Name          *Direction
	CreatedAt     *Direction
	LastUpdatedAt *Direction
	Type          *Direction
}

type DatasourceAttributeSelectionSearchOpt struct {
	Name          bool
	Description   bool
	Content       bool
	Tags          bool
	InheritedFrom bool
	CreatedAt     bool
	CreatedBy     bool
	LastUpdatedAt bool
	LastUpdatedBy bool
}

type Datasource interface {
	GetPolicy(resourceType, resourceId, id string) (*Policy, error)
	GetPolicies(resourceType, resourceId string, opt ...SearchOpts) ([]Policy, int64, error)
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

func WithSelect(opt *DatasourceAttributeSelectionSearchOpt) SearchOpts {
	return func(opts *DatasourceSearchOpts) {
		opts.Select = opt
	}
}
