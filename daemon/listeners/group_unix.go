//go:build !windows

package listeners // import "go.khulnasoft.com/daemon/listeners"

import (
	"fmt"
	"strconv"

	"go.khulnasoft.com/internal/usergroup"
)

const defaultSocketGroup = "docker"

func lookupGID(name string) (int, error) {
	group, err := usergroup.LookupGroup(name)
	if err == nil {
		return group.Gid, nil
	}
	gid, err := strconv.Atoi(name)
	if err == nil {
		return gid, nil
	}
	return -1, fmt.Errorf("group %s not found", name)
}
