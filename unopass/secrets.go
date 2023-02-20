package unopass

import (
    "fmt"
    "os"
    "os/exec"
    "strings"
)

func opPath() string {
    op := "op"
    locate, err := exec.LookPath(op)
    if err != nil {
        fmt.Printf("error: %s\nhttps://github.com/amadotejada/unopass-go", err)
        os.Exit(1)
    }
    return locate
}

func Secret(vault, item, field string, deauthorize bool) string{
    op := opPath()
    cmd := []string{op, "read", "op://" + vault + "/" + item + "/" + field}
    results, err := exec.Command(cmd[0], cmd[1:]...).Output()
    if err != nil {
        Signout(deauthorize)
        fmt.Printf("%s", err)
        os.Exit(1)
    }
    return strings.TrimSpace(string(results))
}


func Signout(deauthorize bool) string {
    if deauthorize {
        op := opPath()
        results, err := exec.Command(op, "signout").Output()
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(1)
        }
        return strings.TrimSpace(string(results))
    }
    return ""
}