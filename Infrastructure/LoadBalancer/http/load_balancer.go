package http

import (
	"context"
	"errors"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
	"time"

	"subscriptions/Infrastructure/LoadBalancer/libraries"
)

type lb struct {
	proxies []*httputil.ReverseProxy
	counter uint64
}

func LoadBalancer(list []string) http.Handler {

	if len(list) == 0 {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "no backends configured", http.StatusServiceUnavailable)
		})
	}

	proxies := make([]*httputil.ReverseProxy, 0, len(list))

	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          200,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   5 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	for _, s := range list {
		u, err := url.Parse(s)
		if err != nil || u.Scheme == "" || u.Host == "" {
			libraries.Errorf("bad backend url: %q", s)
			panic("bad backend url")
		}

		p := httputil.NewSingleHostReverseProxy(u)
		p.Transport = transport
		p.ErrorHandler = func(w http.ResponseWriter, r *http.Request, e error) {
			code := http.StatusBadGateway
			if errors.Is(e, context.DeadlineExceeded) {
				code = http.StatusGatewayTimeout
			}
			http.Error(w, "upstream error", code)
		}

		proxies = append(proxies, p)
	}

	return &lb{proxies: proxies}
}

func (l *lb) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	i := atomic.AddUint64(&l.counter, 1) - 1
	p := l.proxies[i%uint64(len(l.proxies))]
	p.ServeHTTP(w, r)
}
