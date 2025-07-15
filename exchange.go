package aegis

type SecurityPrincipal interface {
	Id() string
	Tenant() string
	Roles() []string
}

type EvaluateRequest struct {
	Operation string
	Resource  Resource
}

type EvaluateResponse struct {
	Granted bool
	Policy  string
}
