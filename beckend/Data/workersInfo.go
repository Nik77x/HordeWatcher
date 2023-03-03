package Data

import "log"

type WorkerInfo struct {
	RequestsFulfilled int          `json:"requests_fulfilled"`
	KudosRewards      float64      `json:"kudos_rewards"`
	KudosDetails      KudosDetails `json:"kudos_details"`
	Performance       string       `json:"performance"`
	Threads           int          `json:"threads"`
	Uptime            int          `json:"uptime"`
	MaintenanceMode   bool         `json:"maintenance_mode"`
	Nsfw              bool         `json:"nsfw"`
	Trusted           bool         `json:"trusted"`
	Flagged           bool         `json:"flagged"`
	UncompletedJobs   int          `json:"uncompleted_jobs"`
	Models            []string     `json:"models"`
	Team              struct {
		Name string `json:"name"`
		Id   string `json:"id"`
	} `json:"team"`
	BridgeAgent      string `json:"bridge_agent"`
	MaxLength        int    `json:"max_length"`
	MaxContextLength int    `json:"max_context_length"`
	Type             string `json:"type"`
	Name             string `json:"name"`
	Id               string `json:"id"`
	Online           bool   `json:"online"`
	Owner            string `json:"owner,omitempty"`
	Info             string `json:"info,omitempty"`
}

type KudosDetails struct {
	Generated float64 `json:"generated"`
	Uptime    int     `json:"uptime"`
}

func (wi WorkerInfo) Print() {
	log.Print("hwh")
}
