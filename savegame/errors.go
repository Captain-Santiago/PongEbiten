package savegame

import "fmt"

type ErrUserNotFound struct {
	name string
}

func (u *ErrUserNotFound) Error() string {
	return fmt.Sprintf("user: \"%s\" not found in the database", u.name)
}
