package aegis

import (
	"fmt"
	"go.uber.org/zap"
)

type EvaluateRequest struct {
	Operation string
	Resource  Resource
}

func (instance *EvaluateRequest) Metadata() Metadata {
	return instance.Resource.Metadata()
}

type EvaluateResponse struct {
}

type Aegis struct {
	datasource Datasource
	logger     *zap.Logger
}

func (instance *Aegis) evaluate(request EvaluateRequest) (EvaluateResponse, error) {
	metadata := request.Metadata()
	logger := instance.logger.With(zap.String("resource_id", metadata.Id))
	page, numberOfElements, err := instance.datasource.GetPolicies(metadata.Type, metadata.Id)
	if err != nil {
		return EvaluateResponse{}, fmt.Errorf("an error occoured while attempting to determine policies. caused by:%v", err.Error())
	}
	logger.Debug(fmt.Sprintf("%d policies retrieved for the resource", numberOfElements))
	for _, policy := range page {

	}
	return EvaluateResponse{}, nil
}
