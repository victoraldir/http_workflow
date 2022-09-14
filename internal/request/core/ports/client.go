package ports

import (
	"github.com/victoraldir/http-follower/internal/request/core/domain"
)

type Client interface {
	Do(req *domain.Request) (*domain.Response, error)
}
