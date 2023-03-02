package httpClient

import (
	"KAIWorkerViz/beckend/Data"
	"context"
	"encoding/json"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"net/http"
	"time"
	_ "time"
)

type Client struct {
	client http.Client
	ctx    context.Context
}

func NewClient(ctx context.Context) Client {
	return Client{
		client: http.Client{Timeout: 10 * time.Second},
		ctx:    ctx,
	}

}

func (c Client) GetJson(url string) (*[]Data.WorkerInfo, error) {
	r, err := c.client.Get(url)
	if err != nil {

		runtime.LogDebugf(c.ctx, "Failed to get url: %q", err.Error())
		return nil, err
	}

	defer r.Body.Close()

	var outInfo []Data.WorkerInfo

	err = json.NewDecoder(r.Body).Decode(&outInfo)
	if err != nil {
		runtime.LogDebugf(c.ctx, "Failed to decode json!:\n%v", err.Error())
		return nil, err
	}

	return &outInfo, nil
}
