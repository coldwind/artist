package ihttp

import "golang.org/x/time/rate"

var (
	MethodGet     = "GET"
	MethodPost    = "POST"
	MethodOptions = "OPTIONS"
)

func WithAddress(host string, port int) Option {
	return func(opt *Service) {
		opt.host = host
		opt.port = port
	}
}

func WithRate(sec, capacity int) Option {
	return func(opt *Service) {
		opt.rateLimiter = rate.NewLimiter(rate.Limit(sec), capacity)
	}
}

func WithCertificate(cert, key string) Option {
	return func(opt *Service) {
		if cert != "" && key != "" {
			opt.https = true
			opt.httpsCertFile = cert
			opt.httpsKeyFile = key
		}
	}
}
