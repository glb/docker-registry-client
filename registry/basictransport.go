package registry

import (
	"net/http"
	"strings"
)

type BasicTransport struct {
	Transport http.RoundTripper
	URL       string
	Username  string
	Password  string
}

func (t *BasicTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := t.Transport.RoundTrip(req)
	if err != nil {
		return resp, err
	}

	if resp.StatusCode == http.StatusUnauthorized && strings.HasPrefix(req.URL.String(), t.URL) {
		if strings.HasPrefix(resp.Header.Get("WWW-Authenticate"), "Basic") {
			if t.Username != "" || t.Password != "" {
				req.SetBasicAuth(t.Username, t.Password)
				return t.Transport.RoundTrip(req)
			}
		}
	}
	return resp, err
}
