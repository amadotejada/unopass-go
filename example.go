package main

import (
	"fmt"
	"github.com/amadotejada/unopass-go/unopass"
)

func main() {
	username := unopass.Secret("personal", "server", "username", true)
	password := unopass.Secret("personal", "server", "password", true)
	fmt.Println(username, password, unopass.Signout(true))
}

