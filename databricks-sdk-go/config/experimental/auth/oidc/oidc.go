// Package oidc provides utilities for working with OIDC ID tokens.
//
// This package is experimental and subject to change.
package oidc

import (
	"context"
	"fmt"
	"os"
)

// IDToken represents an OIDC ID token that can be exchanged for a Databricks
// access token.
type IDToken struct {
	Value string
}

// IDTokenSource is anything that returns an IDToken given an audience.
type IDTokenSource interface {
	IDToken(ctx context.Context, audience string) (*IDToken, error)
}

// IDTokenSourceFn is an adapter to allow the use of ordinary functions as
// IDTokenSource.
//
// Example:
//
//	   ts := IDTokenSourceFn(func(ctx context.Context, aud string) (*IDToken, error) {
//			return &IDToken{}, nil
//	   })
type IDTokenSourceFn func(ctx context.Context, audience string) (*IDToken, error)

func (fn IDTokenSourceFn) IDToken(ctx context.Context, audience string) (*IDToken, error) {
	return fn(ctx, audience)
}

// NewEnvIDTokenSource returns an IDTokenSource that reads the IDtoken from
// environment variable env.
//
// Note that the IDTokenSource does not cache the token and will read the token
// from environment variable env each time.
func NewEnvIDTokenSource(env string) IDTokenSource {
	return IDTokenSourceFn(func(ctx context.Context, _ string) (*IDToken, error) {
		t := os.Getenv(env)
		if t == "" {
			return nil, fmt.Errorf("missing env var %q", env)
		}
		return &IDToken{Value: t}, nil
	})
}

// NewFileTokenSource returns an IDTokenSource that reads the ID token from a
// file. The file should contain a single line with the token.
func NewFileTokenSource(path string) IDTokenSource {
	return IDTokenSourceFn(func(ctx context.Context, _ string) (*IDToken, error) {
		if path == "" {
			return nil, fmt.Errorf("missing path")
		}
		t, err := os.ReadFile(path)
		if err != nil {
			if os.IsNotExist(err) {
				return nil, fmt.Errorf("file %q does not exist", path)
			}
			return nil, err
		}
		if len(t) == 0 {
			return nil, fmt.Errorf("file %q is empty", path)
		}
		return &IDToken{Value: string(t)}, nil
	})
}
