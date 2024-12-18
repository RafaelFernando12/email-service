package domain

import (
	"context"
)

type emailService interface {
	Send(ctx context.Context, person *Person)
}
