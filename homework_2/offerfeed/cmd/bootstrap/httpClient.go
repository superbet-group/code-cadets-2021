package bootstrap

import (
	stdhttp "net/http"
	"time"
)

func HttpClient(timeout time.Duration) *stdhttp.Client {
	return &stdhttp.Client{Timeout: timeout}
}
