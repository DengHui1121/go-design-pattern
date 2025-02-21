package _3_builder

import (
	"errors"
	"fmt"
)

type ResourcePoolConfigOption struct {
	maxTotal int
	maxIdle  int
	minIdle  int
}

type ResourcePoolConfigOptFunc func(option *ResourcePoolConfigOption)

func NewResourcePollConfig(name string, opts ...ResourcePoolConfigOptFunc) (*ResourcePoolConfig, error) {
	if name == "" {
		return nil, errors.New("name is empty")
	}

	option := &ResourcePoolConfigOption{
		maxTotal: 10,
		maxIdle:  9,
		minIdle:  1,
	}

	for _, opt := range opts {
		opt(option)
	}
	if option.maxTotal < 0 || option.maxIdle < 0 || option.minIdle < 0 {
		return nil, fmt.Errorf("args err, option: %v", option)
	}

	if option.maxTotal < option.maxIdle || option.minIdle > option.maxIdle {
		return nil, fmt.Errorf("args err, option: %v", option)
	}

	return &ResourcePoolConfig{
		name:     name,
		maxTotal: option.maxTotal,
		maxIdle:  option.maxIdle,
		minIdle:  option.minIdle,
	}, nil
}

func WithMaxTotal(total int) ResourcePoolConfigOptFunc {
	return func(option *ResourcePoolConfigOption) {
		option.maxTotal = total
	}
}
func WithMaxIdle(idle int) ResourcePoolConfigOptFunc {
	return func(option *ResourcePoolConfigOption) {
		option.maxIdle = idle
	}
}
func WithMinIdle(minIdle int) ResourcePoolConfigOptFunc {
	return func(option *ResourcePoolConfigOption) {
		option.minIdle = minIdle
	}
}
