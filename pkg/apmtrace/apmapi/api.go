package apmapi

import (
	"context"

	"github.com/CloudDetail/apo-module/apm/model/v1"
)

type QueryByApmApi interface {
	QueryList(ctx context.Context, traceId string, startTimeMs int64, attributes string) ([]*model.OtelServiceNode, error)
}
