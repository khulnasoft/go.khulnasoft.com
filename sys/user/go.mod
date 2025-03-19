module go.khulnasoft.com/sys/user

go 1.17

require golang.org/x/sys v0.1.0

retract v0.2.0 // Package go.khulnasoft.com/sys/user/userns was included in this module, but should've been a separate module; see https://go.khulnasoft.com/sys/pull/140#issuecomment-2250644304.
