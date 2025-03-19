//go:build unix && !linux

package chrootarchive

import (
	"testing"

	"go.khulnasoft.com/sys/reexec"
)

func TestMain(m *testing.M) {
	if reexec.Init() {
		return
	}
	m.Run()
}
