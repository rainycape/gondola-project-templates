package main

import (
	"gnd.la/apps/users"
	"gnd.la/orm"
)

type User struct {
	users.User
	// Uncomment this field to enable Facebook integration
	Facebook *users.Facebook
	// Uncomment this field to enable Twitter integration
	Twitter *users.Twitter
}

func init() {
	orm.Register(User{}, nil)
}
