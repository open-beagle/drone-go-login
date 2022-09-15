// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package other

import (
	"net/http"
	"strings"

	"github.com/drone/go-login/login"
	"github.com/drone/go-login/login/internal/oauth2"
	"github.com/drone/go-login/login/logger"
)

var _ login.Middleware = (*Config)(nil)

// Config configures the GitLab auth provider.
type Config struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Server       string
	Scope        []string
	Client       *http.Client
	Logger       logger.Logger
	Dumper       logger.Dumper
}

// Handler returns a http.Handler that runs h at the
// completion of the GitLab authorization flow. The GitLab
// authorization details are available to h in the
// http.Request context.
func (c *Config) Handler(h http.Handler) http.Handler {
	if strings.Contains(c.Server, "github.com") {
		return oauth2.Handler(h, &oauth2.Config{
			BasicAuthOff:     true,
			Client:           c.Client,
			ClientID:         c.ClientID,
			ClientSecret:     c.ClientSecret,
			AccessTokenURL:   c.Server + "/login/oauth/access_token",
			AuthorizationURL: c.Server + "/login/oauth/authorize",
			// Scope: []string{"repo", "user", "read:org"},
			Scope:  c.Scope,
			Logger: c.Logger,
			Dumper: c.Dumper,
		})
	} else if strings.Contains(c.Server, "gitee.com") {
		return oauth2.Handler(h, &oauth2.Config{
			BasicAuthOff:     true,
			Client:           c.Client,
			ClientID:         c.ClientID,
			ClientSecret:     c.ClientSecret,
			RedirectURL:      c.RedirectURL,
			AccessTokenURL:   c.Server + "/oauth/token",
			AuthorizationURL: c.Server + "/oauth/authorize",
			Scope:            []string{"user_info", "projects", "pull_requests", "hook"},
		})
	} else {
		return oauth2.Handler(h, &oauth2.Config{
			BasicAuthOff:     true,
			Client:           c.Client,
			ClientID:         c.ClientID,
			ClientSecret:     c.ClientSecret,
			RedirectURL:      c.RedirectURL,
			AccessTokenURL:   c.Server + "/oauth/token",
			AuthorizationURL: c.Server + "/oauth/authorize",
			Scope:            c.Scope,
		})
	}

}
