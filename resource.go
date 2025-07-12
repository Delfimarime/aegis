package aegis

import "time"

type Resource interface {
	Metadata() Metadata
}

type Metadata struct {
	Id            string
	Type          string
	Policies      []Policy
	Operations    []string
	CreatedAt     *time.Time
	CreatedBy     *Principal
	LastUpdatedAt *time.Time
	LastUpdatedBy *Principal
	Labels        map[string]string
}

type Policy struct {
	Id            string
	Name          string
	Description   string
	Content       []byte
	Tags          []string
	CreatedAt     *time.Time
	CreatedBy     *Principal
	LastUpdatedAt *time.Time
	LastUpdatedBy *Principal
}

type Principal struct {
	Username string
	TenantId string
}
