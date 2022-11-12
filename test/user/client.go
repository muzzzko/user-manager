package user

import (
	"net/http"
	"time"

	httptransport "github.com/go-openapi/runtime/client"

	"github/user-manager/pkg/client/client"
)

var umClient *client.UserManager

func init() {
	httpClient := http.Client{
		Timeout: time.Second * 5,
	}
	r := httptransport.NewWithClient("user-manager:8000", "/", nil, &httpClient)
	umClient = client.New(r, nil)
}
