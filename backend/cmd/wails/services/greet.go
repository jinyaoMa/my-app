package services

import (
	"context"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type GreetService struct{}

func (s *GreetService) ServiceStartup(ctx context.Context, options application.ServiceOptions) error {
	return nil
}

func (s *GreetService) ServiceShutdown() error {
	return nil
}

func (s *GreetService) Greet(name string) string {
	return "Hello " + name + "!"
}
