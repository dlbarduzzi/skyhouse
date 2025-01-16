package skyhouse

import (
	"fmt"
	"log/slog"
	"sync"
)

type Skyhouse struct {
	config *Config
	logger *slog.Logger
	wg     *sync.WaitGroup
}

func NewSkyhouse(config *Config, logger *slog.Logger) (*Skyhouse, error) {
	cfg, err := config.parse()
	if err != nil {
		return nil, err
	}
	return &Skyhouse{
		config: cfg,
		logger: logger,
		wg:     &sync.WaitGroup{},
	}, nil
}

func (s *Skyhouse) Port() int {
	return s.config.Port
}

func (s *Skyhouse) Shutdown() {
	s.wg.Wait()
}

func (s *Skyhouse) Background(fn func()) {
	s.wg.Add(1)

	go func() {
		defer s.wg.Done()

		defer func() {
			if err := recover(); err != nil {
				s.logger.Error(fmt.Sprintf("%v", err))
			}
		}()

		fn()
	}()
}
