package isms

type IProvider interface {
	Init(*Config)
	Send(string, []string) (string, error)
	SendMultiple([]string, []string) (string, error)
}

type Service struct {
	p IProvider
}

type Config struct {
	SecretId   string
	SecretKey  string
	SdkAppId   string
	Sign       string
	TemplateId string
}

type Option func(*Service)

type ProviderType int

var (
	QCloud ProviderType = 1
)

func New(provider ProviderType, config *Config) *Service {
	s := &Service{}
	switch provider {
	case QCloud:
		s.p = &QcloudProvider{}
		s.p.Init(config)
	}

	return s
}
