
# unopass-go
##### Written by [Amado Tejada](https://www.linkedin.com/in/amadotejada/)

##
*unopass-go* is a Go package that allows you to retrieve secrets from the 1Password CLI using your biometrics.

<img src="https://cdn.jsdelivr.net/npm/programming-languages-logos/src/python/python.png" height="50"> Python version: [HERE](https://github.com/amadotejada/unopass)

#
#### Install

```bash
go get -u github.com/amadotejada/unopass-go
```
#### Usage

```go
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
```

#### Requirements:
 * [Go](https://go.dev/dl/) v1.20 or higher
 * [ 1Password CLI ](https://developer.1password.com/docs/cli/get-started#install) v2.4.1 or higher
 * [ 1Password App ](https://1password.com/downloads/) v8.7.1 or higher
 * [Biometrics](https://developer.1password.com/docs/cli/get-started#turn-on-biometric-unlock) enabled


### Security
Authorization expires after 10 minutes of inactivity in the session. There's a hard limit of 12 hours, after which you must reauthorize.
