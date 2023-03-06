package main

import (
	"KAIWorkerViz/beckend"
	"KAIWorkerViz/beckend/Data"
	"KAIWorkerViz/beckend/httpClient"
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx    context.Context
	client httpClient.Client
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx

	a.client = httpClient.NewClient(ctx)

}

// domReady is called after the front-end dom has been loaded
func (a *App) domReady(ctx context.Context) {
	// Add your action here

	// download json from all sources
	// deserialize to struct
	// check for duplicates?
	// send to frontend

	backend.RunFuncAtInterval(3000, func() {
		workersInfo, err := a.client.GetJson("https://stablehorde.net/api/v2/workers?type=text")
		if err != nil {
			panic(err.Error())
			runtime.LogDebugf(a.ctx, err.Error())
		} else {
			runtime.EventsEmit(a.ctx, "update_list", *workersInfo)
		}

	}, true)
}

func (a *App) GetCurrentWorkers() []Data.WorkerInfo {
	wi, err := a.client.GetJson("https://stablehorde.net/api/v2/workers?type=text")
	if err != nil {
		return nil
	}
	return *wi
}
