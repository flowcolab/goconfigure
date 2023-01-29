package goconfigure

type ConfigSource interface {
	GetProperties() (*ConfigProperties, error)
}
