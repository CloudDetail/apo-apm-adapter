package jaeger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/CloudDetail/apo-module/apm/model/v1"
)

type JaegerApi struct {
	Address string
	Timeout time.Duration

	client *http.Client
}

func NewJaegerApi(address string, timeout int64) *JaegerApi {
	var client *http.Client = &http.Client{
		Timeout:   time.Duration(timeout) * time.Second,
		Transport: http.DefaultTransport,
	}
	return &JaegerApi{
		Address: fmt.Sprintf("http://%s/api/traces", address),
		Timeout: time.Duration(timeout) * time.Second,
		client:  client,
	}
}

func (jaeger *JaegerApi) QueryList(ctx context.Context, traceId string, startTimeMs int64, attributes string) ([]*model.OtelServiceNode, error) {
	resp, err := jaeger.queryJson(ctx, fmt.Sprintf("%s/%s", jaeger.Address, traceId), jaeger.Timeout)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("[x Trace NotFound] Jaeger unexpected status code: %d, msg: %s", resp.StatusCode, string(bodyBytes))
	}

	var response JaegerResponse
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	if len(response.Data) == 0 {
		return nil, fmt.Errorf("[x Trace NotFound] Jaeger traceId: %s", traceId)
	}
	return ConvertToServiceNodes(&response.Data[0])
}

func (jaeger *JaegerApi) queryJson(ctx context.Context, url string, timeout time.Duration) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return jaeger.client.Do(req)
}
