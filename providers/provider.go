package providers

import (
	"github.com/willbarkoff/autoscout/config"
	"github.com/willbarkoff/autoscout/data"
)

// A Provider provides data to be used by autoscout.
type Provider interface {
	GetName() string
	GetData(config.Config) (map[int]data.Team, [][]string)
}
