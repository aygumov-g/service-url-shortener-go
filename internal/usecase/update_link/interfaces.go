package update_link

import (
	"context"
	"time"
)

type LinkRepository interface {
	Update(ctx context.Context, id int64, now time.Time) error
}

type Clock interface {
	Now() time.Time
}
