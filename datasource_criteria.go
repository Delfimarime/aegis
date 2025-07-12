package aegis

import "time"

type Criteria struct {
	Search        *string
	CreatedAt     *TimeCriteria
	CreatedBy     *PrincipalCriteria
	LastUpdatedAt *TimeCriteria
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
