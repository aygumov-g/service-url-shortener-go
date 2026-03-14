package update_link

import "time"

type LinkRepository interface {
	Update(id int64, now time.Time) error
}

type Clock interface {
	Now() time.Time
}
