//go:build windows

package metrics

import (
	"go.khulnasoft.com/pkg/plugingetter"
	"go.khulnasoft.com/plugin"
)

func RegisterPlugin(*plugin.Store, string) error { return nil }
func CleanupPlugin(plugingetter.PluginGetter)    {}
