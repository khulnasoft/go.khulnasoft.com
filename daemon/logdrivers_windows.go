package daemon // import "go.khulnasoft.com/daemon"

import (
	// Importing packages here only to make sure their init gets called and
	// therefore they register themselves to the logdriver factory.
	_ "go.khulnasoft.com/daemon/logger/awslogs"
	_ "go.khulnasoft.com/daemon/logger/etwlogs"
	_ "go.khulnasoft.com/daemon/logger/fluentd"
	_ "go.khulnasoft.com/daemon/logger/gcplogs"
	_ "go.khulnasoft.com/daemon/logger/gelf"
	_ "go.khulnasoft.com/daemon/logger/jsonfilelog"
	_ "go.khulnasoft.com/daemon/logger/loggerutils/cache"
	_ "go.khulnasoft.com/daemon/logger/splunk"
	_ "go.khulnasoft.com/daemon/logger/syslog"
)
