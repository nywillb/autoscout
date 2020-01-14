package main

import (
	"github.com/willbarkoff/autoscout/config"
	"github.com/willbarkoff/autoscout/data"
)

// A Provider provides data to be used by autoscout.
type Provider interface {
	CheckIsProvider(string) bool
	GetData(config.Config) (map[int]data.Team, [][]string)
}
