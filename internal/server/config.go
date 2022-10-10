package server

import (
	"time"

	"github.com/NathanBak/go-server/pkg/widget"
)

const (
	defaultPort         = 8080
	defaultReadTimeout  = 3 * time.Second
	defaultWriteTimeout = 3 * time.Second
)

// Config contains information necessary to set up a Server.
type Config struct {
	Port         int           `json:"port" mapstructure:"PORT"`
	ReadTimeout  time.Duration `json:"readTimeout" mapstructure:"READ_TIMEOUT"`
	WriteTimeout time.Duration `json:"writeTimeout" mapstructure:"WRITE_TIMEOUT"`

	Logger Logger `json:"-" mapstructure:"-"`

	IncludeStatusCodeInMessages bool `json:"-" mapstructure:"-"`

	Storage Storage `json:"-" mapstructure:"-"`
}

// The Logger interface defines the methods required by the Server for logging.
type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warning(msg string)
	Error(msg string)
}

// The Storage interface defines the methods required to access backing Widget storage.
type Storage interface {
	Get(string) (widget.Widget, bool, error)
	Set(string, widget.Widget) error
	Delete(string) (widget.Widget, bool, error)
	Keys() ([]string, error)
}

// applyDefaultValues will set certain uninitialized properties to the default.
func (c *Config) applyDefaultValues() {
	if c.Port == 0 {
		c.Port = defaultPort
	}

	if c.ReadTimeout == 0 {
		c.ReadTimeout = defaultReadTimeout
	}

	if c.WriteTimeout == 0 {
		c.WriteTimeout = defaultWriteTimeout
	}

	if c.Logger == nil {
		c.Logger = defaultLogger{}
	}
}
