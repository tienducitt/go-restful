package handler

import (
	"net/http"
	"os"
	"time"
)

var hostName string
var deployTime string

func init() {
	hostName, _ = os.Hostname()
	deployTime = time.Now().Format(time.RFC3339)
}

func Ping(req *http.Request) Response {
	resp := PingResponse{
		Status:      "OK",
		ServiceName: "my-service",
		Description: "",
		HostName:    hostName,
		DeployTime:  deployTime,
	}

	return Success(resp)
}

type PingResponse struct {
	Status      string
	ServiceName string
	HostName    string
	DeployTime  string
	Description string
}
