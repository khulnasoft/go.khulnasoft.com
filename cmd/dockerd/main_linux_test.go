package main

import (
	"testing"

	"go.khulnasoft.com/sys/reexec"
)

func TestMain(m *testing.M) {
	reexec.Register(testListenerNoAddrCmdPhase1, initListenerTestPhase1)
	reexec.Register(testListenerNoAddrCmdPhase2, initListenerTestPhase2)
	if reexec.Init() {
		return
	}
	m.Run()
}
