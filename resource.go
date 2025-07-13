package aegis

import "time"

type Resource interface {
	Metadata() Metadata
}

type Metadata struct {
	Id            string
	Type          string
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
	Type          string
	Description   string
	Content       []byte
	Tags          []string
	InheritedFrom Named
	CreatedAt     *time.Time
	CreatedBy     *Principal
	LastUpdatedAt *time.Time
	LastUpdatedBy *Principal
}

type Named struct {
	Id   string
	Name string
	Type string
}

type Principal struct {
	Username string
	TenantId string
}
