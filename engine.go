package aegis

import (
	"context"
	"fmt"
	"go.uber.org/zap"
)

const (
	obsNamespace = "github.com/raitonbl/aegis"
)

type Aegis struct {
	datasource Datasource
	logger     *zap.Logger
}

func (instance *Aegis) evaluate(ctx context.Context, principal SecurityPrincipal, request EvaluateRequest) (EvaluateResponse, error) {
	if principal == nil {
		return EvaluateResponse{Granted: false}, nil
	}
	metadata := request.Resource.Metadata()
	logger := instance.logger.With(
		zap.String(fmt.Sprintf("%s/resource.id", obsNamespace), metadata.Id),
		zap.String(fmt.Sprintf("%s/principal.id", obsNamespace), principal.Id()),
		zap.String(fmt.Sprintf("%s/principal.tenant", obsNamespace), principal.Tenant()),
	)
	page, numberOfElements, err := instance.datasource.GetPolicies(ctx, metadata.Type, metadata.Id)
	if err != nil {
		return EvaluateResponse{}, fmt.Errorf("an error occoured while attempting to determine policies. caused by:%v", err.Error())
	}
	logger.Debug(fmt.Sprintf("%d policies retrieved for the resource", numberOfElements))
	for _, policy := range page {
		canAccess, prob := instance.evaluateRego(ctx, principal, request, policy.Content)
		if prob != nil {
			return EvaluateResponse{}, fmt.Errorf("an error occoured while attempting to process policy. caused by:%v", prob.Error())
		}
		if !canAccess {
			logger.Debug(
				fmt.Sprintf("policy %s doesn't grant access to resource %s", policy.Name, metadata.Name),
				zap.String(fmt.Sprintf("%s/policy.id", obsNamespace), policy.Id),
			)
			continue
		}
		logger.Info(
			fmt.Sprintf("policy %s granted access to resource %s", policy.Name, metadata.Name),
			zap.String(fmt.Sprintf("%s/policy.id", obsNamespace), policy.Id),
		)
		return EvaluateResponse{Granted: true, Policy: policy.Id}, nil
	}
	return EvaluateResponse{Granted: false}, nil
}

func (instance *Aegis) evaluateRego(ctx context.Context, principal SecurityPrincipal, request EvaluateRequest, content []byte) (bool, error) {
	return false, nil
}
