package main

import (
	"KAIWorkerViz/beckend"
	"KAIWorkerViz/beckend/apiInterface"
	"KAIWorkerViz/beckend/data"
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx          context.Context
	apiInterface apiInterface.ApiInterface
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx

	a.apiInterface = apiInterface.New(ctx)

}

// domReady is called after the front-end dom has been loaded
func (a *App) domReady(_ context.Context) {
	// Add your action here

	backend.RunFuncAtInterval(3000, func() {

		textWorkers := a.apiInterface.GetTextWorkers()
		imageWorkers := a.apiInterface.GetImageWorkers()

		runtime.EventsEmit(a.ctx, "update_list", *textWorkers, *imageWorkers)

	}, true)
}

func (a *App) GetTextWorkers() []data.TextWorker {
	return *a.apiInterface.GetTextWorkers()
}
func (a *App) GetImageWorkers() []data.ImageWorker {
	return *a.apiInterface.GetImageWorkers()
}
