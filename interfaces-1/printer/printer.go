package printer

import (
	"fmt"
	"github.com/bfoc/gocasts/interfaces-1/user"
	"github.com/pkg/errors"
)

type Printer struct {
	User *user.Client
}

func (p *Printer) Print() error {
	users, err := p.User.Get()
	if err != nil {
		return errors.Wrap(err, "unable to get users for printing")
	}

	// NOTE: Use an actual CSV Formatter for production use cases
	fmt.Println("ID,Email,Username,Name")
	for _, entry := range users {
		fmt.Printf("%v,%v,%v,%v\n", entry.ID, entry.Email, entry.Username, entry.Name)
	}

	return nil
}
