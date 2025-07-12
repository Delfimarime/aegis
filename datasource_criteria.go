package aegis

import (
	"github.com/thoas/go-funk"
	"time"
)

type Criteria struct {
	Search        *string
	CreatedAt     *TimeCriteria
	LastUpdatedAt *TimeCriteria
	CreatedBy     *PrincipalCriteria
	LastUpdatedBy *PrincipalCriteria
	Tags          *ArrayCriteria[string]
}

type EqCriteria[T any] struct {
	Eq T
}

type ArrayCriteria[T any] struct {
	Contains []T
	*EqCriteria[[]T]
}

type PrincipalCriteria struct {
	Username EqCriteria[string]
	Tenant   *EqCriteria[string]
}

type TimeCriteria struct {
	Eq          *time.Time
	GreaterThan *time.Time
	LesserThan  *time.Time
}

func WithSearch(value string) SearchOpts {
	return func(opts *DatasourceSearchOpts) {
		if value == "" {
			opts.Search = nil
			return
		}
		opts.Search = funk.PtrOf(value).(*string)
	}
}

func WithTagsEq(v ...string) SearchOpts {
	return func(opts *DatasourceSearchOpts) {
		if len(v) == 0 {
			opts.Tags = nil
			return
		}
		opts.Tags = &ArrayCriteria[string]{
			EqCriteria: &EqCriteria[[]string]{Eq: v},
		}
	}
}

func WithTagsContains(v ...string) SearchOpts {
	return func(opts *DatasourceSearchOpts) {
		if len(v) == 0 {
			opts.Tags = nil
			return
		}
		opts.Tags = &ArrayCriteria[string]{
			Contains: v,
		}
	}
}

func WithCreatedByEq(value, tenantId string) SearchOpts {
	return func(opts *DatasourceSearchOpts) {
		opts.CreatedBy = &PrincipalCriteria{
			Username: EqCriteria[string]{Eq: value},
		}
		if tenantId != "" {
			opts.CreatedBy.Tenant = &EqCriteria[string]{Eq: tenantId}
		}

	}
}

func WithLastUpdatedByEq(value, tenantId string) SearchOpts {
	return func(opts *DatasourceSearchOpts) {
		opts.LastUpdatedBy = &PrincipalCriteria{
			Username: EqCriteria[string]{Eq: value},
		}
		if tenantId != "" {
			opts.LastUpdatedBy.Tenant = &EqCriteria[string]{Eq: tenantId}
		}
	}
}
func setTimeCriteria(tc **TimeCriteria, setter func(*TimeCriteria)) {
	if *tc == nil {
		*tc = &TimeCriteria{}
	}
	setter(*tc)
}

func WithCreatedAtEq(value time.Time) SearchOpts {
	return func(opts *DatasourceSearchOpts) {
		setTimeCriteria(&opts.CreatedAt, func(tc *TimeCriteria) {
			tc.Eq = funk.PtrOf(value).(*time.Time)
		})
	}
}

func WithCreatedAtGt(value time.Time) SearchOpts {
	return func(opts *DatasourceSearchOpts) {
		setTimeCriteria(&opts.CreatedAt, func(tc *TimeCriteria) {
			tc.GreaterThan = funk.PtrOf(value).(*time.Time)
		})
	}
}

func WithCreatedAtGte(value time.Time) SearchOpts {
	return func(opts *DatasourceSearchOpts) {
		setTimeCriteria(&opts.CreatedAt, func(tc *TimeCriteria) {
			ptr := funk.PtrOf(value).(*time.Time)
			tc.Eq = ptr
			tc.GreaterThan = ptr
		})
	}
}

func WithCreatedAtLt(value time.Time) SearchOpts {
	return func(opts *DatasourceSearchOpts) {
		setTimeCriteria(&opts.CreatedAt, func(tc *TimeCriteria) {
			tc.LesserThan = funk.PtrOf(value).(*time.Time)
		})
	}
}

func WithCreatedAtLte(value time.Time) SearchOpts {
	return func(opts *DatasourceSearchOpts) {
		setTimeCriteria(&opts.CreatedAt, func(tc *TimeCriteria) {
			ptr := funk.PtrOf(value).(*time.Time)
			tc.Eq = ptr
			tc.LesserThan = ptr
		})
	}
}

func WithLastUpdatedAtEq(value time.Time) SearchOpts {
	return func(opts *DatasourceSearchOpts) {
		setTimeCriteria(&opts.LastUpdatedAt, func(tc *TimeCriteria) {
			tc.Eq = funk.PtrOf(value).(*time.Time)
		})
	}
}

func WithLastUpdatedAtGt(value time.Time) SearchOpts {
	return func(opts *DatasourceSearchOpts) {
		setTimeCriteria(&opts.LastUpdatedAt, func(tc *TimeCriteria) {
			tc.GreaterThan = funk.PtrOf(value).(*time.Time)
		})
	}
}

func WithLastUpdatedAtGte(value time.Time) SearchOpts {
	return func(opts *DatasourceSearchOpts) {
		setTimeCriteria(&opts.LastUpdatedAt, func(tc *TimeCriteria) {
			ptr := funk.PtrOf(value).(*time.Time)
			tc.Eq = ptr
			tc.GreaterThan = ptr
		})
	}
}

func WithLastUpdatedAtLt(value time.Time) SearchOpts {
	return func(opts *DatasourceSearchOpts) {
		setTimeCriteria(&opts.LastUpdatedAt, func(tc *TimeCriteria) {
			tc.LesserThan = funk.PtrOf(value).(*time.Time)
		})
	}
}

func WithLastUpdatedAtLte(value time.Time) SearchOpts {
	return func(opts *DatasourceSearchOpts) {
		setTimeCriteria(&opts.LastUpdatedAt, func(tc *TimeCriteria) {
			ptr := funk.PtrOf(value).(*time.Time)
			tc.Eq = ptr
			tc.LesserThan = ptr
		})
	}
}
