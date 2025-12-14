package rate

import "crud/domain/user"

type Rate struct {
	ID        string
	Name      string
	Rate      int
	TmdbId    string
	ImagePath string
	Comment   string
	User      user.User
}
