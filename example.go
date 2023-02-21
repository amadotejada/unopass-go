package main

import (
	"fmt"
	"github.com/amadotejada/unopass-go/unopass"
)

/*
Parameters:

    VAULT (string): the name of the 1Password vault
    ITEM (string): the name of the item within the specified vault
    FIELD (string): the name of the field to retrieve within the specified item

Returns:
(string) the value of the specified fields within the specified item

OPTIONAL: To sign out of the 1Password CLI session, call 'unopass.Signout(true)'
*/

func main() {
	username := unopass.Secret("personal", "server", "username", true)
	password := unopass.Secret("personal", "server", "password", true)
	fmt.Println(username, password, unopass.Signout(true))
}
