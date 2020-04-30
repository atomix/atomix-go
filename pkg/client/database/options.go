// Copyright 2019-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package database

import (
	"os"
	"time"
)

func applyOptions(opts ...Option) databaseOptions {
	options := &databaseOptions{
		scope:          os.Getenv("ATOMIX_SCOPE"),
		sessionTimeout: 1 * time.Minute,
	}
	for _, opt := range opts {
		opt.apply(options)
	}
	return *options
}

type databaseOptions struct {
	scope          string
	sessionTimeout time.Duration
}

// Option provides a database option
type Option interface {
	apply(options *databaseOptions)
}

// WithScope configures the application scope for the client
func WithScope(scope string) Option {
	return &scopeOption{scope: scope}
}

type scopeOption struct {
	scope string
}

func (o *scopeOption) apply(options *databaseOptions) {
	options.scope = o.scope
}

// WithSessionTimeout sets the session timeout for the client
func WithSessionTimeout(timeout time.Duration) Option {
	return &sessionTimeoutOption{
		timeout: timeout,
	}
}

type sessionTimeoutOption struct {
	timeout time.Duration
}

func (s *sessionTimeoutOption) apply(options *databaseOptions) {
	options.sessionTimeout = s.timeout
}
