package godict

import "fmt"

type Option func(d *dictionary) error

var (
	defaultOpts = []Option{}
)

func WithPath(paths ...string) Option {
	return func(d *dictionary) error {
		if len(paths) == 0 {
			return fmt.Errorf("invalid path length: %d", len(paths))
		}

		d.paths = append(d.paths, paths...)

		return nil
	}
}
