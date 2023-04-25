package data

import (
	"regexp"
	"strconv"
)

type Worker struct {
	// common data
	RequestsFulfilled int      `json:"requests_fulfilled"`
	KudosRewards      float32  `json:"kudos_rewards"`
	Performance       string   `json:"performance"`
	Uptime            int      `json:"uptime"`
	MaintenanceMode   bool     `json:"maintenance_mode"`
	UncompletedJobs   int      `json:"uncompleted_jobs"`
	Models            []string `json:"models"`
	Queue             float64  `json:"queue"`
	Name              string   `json:"name"`
	Type              string   `json:"type"`
	Id                string   `json:"id"`
	Online            bool     `json:"online"`
	Info              string   `json:"info,omitempty"`
}

func (w *Worker) GetPerformance() float32 {
	xp := regexp.MustCompile("[0-9.]")
	out, err := strconv.ParseFloat(xp.FindString(w.Performance), 32)
	if err != nil {
		return -1
	}
	return float32(out)
}

func (w *Worker) SetWorker(worker *Worker) {
	w = worker
}

//type KudosDetails struct {
//	Generated float64 `json:"generated"`
//	Uptime    int     `json:"uptime"`
//}

//KudosDetails      KudosDetails `json:"kudos_details"`
//Threads           int          `json:"threads"`
//Nsfw              bool         `json:"nsfw"`
//Trusted           bool         `json:"trusted"`
//Flagged           bool         `json:"flagged"`
//BridgeAgent string `json:"bridge_agent"`
