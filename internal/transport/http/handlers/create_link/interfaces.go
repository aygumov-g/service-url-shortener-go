package create_link

import "context"

type CreateLinkUsecase interface {
	Execute(ctx context.Context, original string) (string, error)
}
