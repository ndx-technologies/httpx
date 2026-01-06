package httpx_test

import (
	"testing"
	"time"

	"github.com/ndx-technologies/httpx"
)

func TestCacheControlConfig(t *testing.T) {
	tests := map[string]httpx.CacheControlConfig{
		"private":               {},
		"no-cache":              {NoCache: true},
		"no-store":              {NoStore: true},
		"private, max-age=3600": {MaxAge: time.Second * 3600},
		"public, max-age=3600":  {Public: true, MaxAge: time.Second * 3600},
		"private, immutable":    {Immutable: true},
	}
	for s, v := range tests {
		if got := v.String(); got != s {
			t.Error(got, s)
		}
	}
}
