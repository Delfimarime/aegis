package aegis

import (
	"context"
	"fmt"
	"go.uber.org/zap"
)

const (
	obsNamespace = "github.com/raitonbl/aegis"
)

type SecurityPrincipal interface {
	Id() string
	Tenant() string
	Roles() []string
}

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

func (instance *Aegis) evaluate(ctx context.Context, principal SecurityPrincipal, request EvaluateRequest) (EvaluateResponse, error) {
	metadata := request.Metadata()
	logger := instance.logger.With(zap.String(fmt.Sprintf("%s/resource.id", obsNamespace), metadata.Id))
	page, numberOfElements, err := instance.datasource.GetPolicies(ctx, metadata.Type, metadata.Id)
	if err != nil {
		return EvaluateResponse{}, fmt.Errorf("an error occoured while attempting to determine policies. caused by:%v", err.Error())
	}
	logger.Debug(fmt.Sprintf("%d policies retrieved for the resource", numberOfElements))
	for _, policy := range page {
		//policy.Content
	}
	return EvaluateResponse{}, nil
}
