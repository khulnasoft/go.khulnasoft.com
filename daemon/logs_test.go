package daemon // import "go.khulnasoft.com/daemon"

import (
	"testing"

	containertypes "go.khulnasoft.com/api/types/container"
)

func TestMergeAndVerifyLogConfigNilConfig(t *testing.T) {
	d := &Daemon{defaultLogConfig: containertypes.LogConfig{Type: "json-file", Config: map[string]string{"max-file": "1"}}}
	cfg := containertypes.LogConfig{Type: d.defaultLogConfig.Type}
	if err := d.mergeAndVerifyLogConfig(&cfg); err != nil {
		t.Fatal(err)
	}
}
