package httpx

import (
	"net/http"
	"strconv"
	"strings"
	"time"
)

type CacheControlConfig struct {
	Public          bool          `json:"public"`
	MaxAge          time.Duration `json:"max_age,format:units"`
	NoCache         bool          `json:"no_cache"`
	NoStore         bool          `json:"no_store"`
	Immutable       bool          `json:"immutable"`
	AllowAllMethods bool          `json:"allow_all_methods"` // cache control is usually defined for GET and OPTIONS. set this to true to allow all methods.
}

func (s CacheControlConfig) String() string {
	if s.NoCache {
		return "no-cache"
	}
	if s.NoStore {
		return "no-store"
	}

	var b strings.Builder

	if s.Public {
		b.WriteString("public")
	} else {
		b.WriteString("private")
	}

	if s.MaxAge > 0 {
		b.WriteString(", max-age=")
		b.WriteString(strconv.Itoa(int(s.MaxAge.Seconds())))
	}

	if s.Immutable {
		b.WriteString(", immutable")
	}

	return b.String()
}

func CacheControl(c CacheControlConfig) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodGet || r.Method == http.MethodOptions || c.AllowAllMethods {
				w.Header().Set("Cache-Control", c.String())
			}
			h.ServeHTTP(w, r)
		})
	}
}
