package skyhouse

import (
	"fmt"
	"log/slog"
	"sync"
)

type Skyhouse struct {
	logger *slog.Logger
	wg     *sync.WaitGroup
}

func NewSkyhouse(logger *slog.Logger) (*Skyhouse, error) {
	return &Skyhouse{
		logger: logger,
		wg:     &sync.WaitGroup{},
	}, nil
}

// TODO: Get port from config.
func (s *Skyhouse) Port() int {
	return 8000
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
