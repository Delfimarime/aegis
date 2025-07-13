package aegis

type EvaluateResponse struct {
}

type Aegis struct {
	datasource Datasource
}

func (instance *Aegis) evaluate(resource Resource) (EvaluateResponse, error) {
	metadata := resource.Metadata()
	return EvaluateResponse{}, nil
}
