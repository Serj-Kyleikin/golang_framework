package middlewares

import (
	"net/http"
	"os"
	"sync"

	"subscriptions/Infrastructure/LoadBalancer/libraries"

	"gopkg.in/yaml.v3"
)

const defaultConfigPath = "./settings/middlewares.yaml"

const configEnvVar = "MIDDLEWARES_CONFIG"

type Middleware func(http.Handler) http.Handler

func Apply(h http.Handler) http.Handler {
	return Chain(h, fromYAMLConfig()...)
}

func Chain(h http.Handler, m ...Middleware) http.Handler {
	for i := len(m) - 1; i >= 0; i-- {
		h = m[i](h)
	}
	return h
}

type MiddlewareFactory func(cfg map[string]any) (Middleware, error)

var (
	registryMu sync.RWMutex
	registry   = map[string]MiddlewareFactory{}
)

func Register(name string, f MiddlewareFactory) {
	registryMu.Lock()
	registry[name] = f
	registryMu.Unlock()
}

type yamlConfig struct {
	Middlewares []yamlMiddleware `yaml:"middlewares"`
}

type yamlMiddleware struct {
	Name    string         `yaml:"name"`
	Enabled *bool          `yaml:"enabled,omitempty"`
	Config  map[string]any `yaml:"config,omitempty"`
}

var (
	loadOnce sync.Once
	cfg      yamlConfig
	cfgOK    bool
)

func fromYAMLConfig() []Middleware {
	loadOnce.Do(func() {
		path := os.Getenv(configEnvVar)
		if path == "" {
			path = defaultConfigPath
		}

		b, err := os.ReadFile(path)
		if err != nil {
			libraries.Errorf("middlewares config read error (%s): %v", path, err)
			return
		}

		if err := yaml.Unmarshal(b, &cfg); err != nil {
			libraries.Errorf("middlewares config parse error (%s): %v", path, err)
			return
		}

		if len(cfg.Middlewares) == 0 {
			libraries.Errorf("middlewares config is empty (%s)", path)
			return
		}

		cfgOK = true
	})

	if !cfgOK {
		return nil
	}

	out := make([]Middleware, 0, len(cfg.Middlewares))

	for _, item := range cfg.Middlewares {
		if item.Name == "" {
			libraries.Errorf("middleware with empty name in config — skipped")
			continue
		}

		enabled := true
		if item.Enabled != nil {
			enabled = *item.Enabled
		}
		if !enabled {
			continue
		}

		registryMu.RLock()
		factory, ok := registry[item.Name]
		registryMu.RUnlock()

		if !ok {
			libraries.Errorf("middleware %q configured, but implementation not found (not registered) — skipped", item.Name)
			continue
		}

		mw, err := factory(item.Config)
		if err != nil {
			libraries.Errorf("middleware %q configured, but failed to build: %v — skipped", item.Name, err)
			continue
		}
		if mw == nil {
			libraries.Errorf("middleware %q factory returned nil — skipped", item.Name)
			continue
		}

		out = append(out, mw)
	}

	return out
}
