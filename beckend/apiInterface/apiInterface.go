package apiInterface

import (
	"KAIWorkerViz/beckend/data"
	"context"
	"encoding/json"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"net/http"
	"time"
	_ "time"
)

const textWorkersURL = "https://horde.koboldai.net/api/v2/workers?type=text"
const textModelsStatusURL = "https://horde.koboldai.net/api/v2/status/models?type=text"
const imageWorkersURL = "https://horde.koboldai.net/api/v2/workers?type=image"
const imageModelsStatusURL = "https://horde.koboldai.net/api/v2/status/models?type=image"

type ApiInterface struct {
	client http.Client
	ctx    context.Context
}

func New(ctx context.Context) ApiInterface {
	return ApiInterface{
		client: http.Client{Timeout: 10 * time.Second},
		ctx:    ctx,
	}

}

func (ai ApiInterface) GetTextWorkers() *[]data.TextWorker {

	var outWorkers []data.TextWorker

	var modelsStatus []data.ModelStatus

	ai.parseJson(textWorkersURL, &outWorkers)

	ai.parseJson(textModelsStatusURL, &modelsStatus)

	// assign queued jobs to workers. If model is not found set queue to -1.

	for i := range outWorkers {
		wk := &outWorkers[i]
		wk.Queue = -1
		for j := range modelsStatus {
			ms := &modelsStatus[j]
			if wk.Models[0] == ms.Name {
				wk.Queue = ms.Queued
				break
			}
		}
	}

	return &outWorkers
}

func (ai ApiInterface) GetImageWorkers() *[]data.ImageWorker {

	var outWorkers []data.ImageWorker

	var modelsStatus []data.ModelStatus

	ai.parseJson(imageWorkersURL, &outWorkers)

	ai.parseJson(imageModelsStatusURL, &modelsStatus)

	// assign queued jobs to workers. If model is not found set queue to -1.
	for i := range outWorkers {
		wk := &outWorkers[i]
		wk.Queue = -1
		for j := range modelsStatus {
			ms := &modelsStatus[j]
			if wk.Models[0] == ms.Name {

				// queue is in pixelSteps I think, so divide by 1,000,000 to get MegaPixelSteps.
				wk.Queue = ms.Queued / 1_000_000
				break
			}
		}
	}

	return &outWorkers
}

func (ai ApiInterface) parseJson(url string, target any) {
	r, err := ai.client.Get(url)
	defer r.Body.Close()
	if err != nil {
		runtime.LogFatalf(ai.ctx, "Failed to get url: %q", err.Error())
	}

	err = json.NewDecoder(r.Body).Decode(target)
	if err != nil {
		runtime.LogFatalf(ai.ctx, "Failed to parse json!:\n", err.Error())
	}
}
