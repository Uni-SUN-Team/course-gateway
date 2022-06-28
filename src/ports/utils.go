package ports

import "net/http"

type HTTPRequest interface {
	HTTPRequest(url string, method string, payload []byte) (*http.Response, error)
}
