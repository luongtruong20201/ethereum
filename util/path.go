package util

import (
	"os/user"
	"strings"
)

func ExpandHomePath(p string) (path string) {
	path = p
	if path[:2] == "~/" {
		usr, _ := user.Current()
		dir := usr.HomeDir
		path = strings.Replace(p, "~", dir, 1)
	}
	return
}
