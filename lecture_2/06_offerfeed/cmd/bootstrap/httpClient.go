package bootstrap

import (
	"time"

	stdHttp "net/http"
)

func HttpClient() *stdHttp.Client {
	return &stdHttp.Client{Timeout: time.Second * 5}
}
