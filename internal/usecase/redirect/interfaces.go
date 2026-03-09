package redirect

import (
	"time"

	link_d "github.com/aygumov-g/service-url-shortener-go/internal/domain/link"
)

type LinkRepository interface {
	Update(id int64, now time.Time) error
	GetByID(id int64) (*link_d.Link, error)
	GetByCustomCode(code string) (*link_d.Link, error)
}

type Generator interface {
	Decode(code string) (int64, error)
}

type Clock interface {
	Now() time.Time
}
