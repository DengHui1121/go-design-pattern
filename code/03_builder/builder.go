package _3_builder

import "fmt"

type ResourcePoolConfig struct {
	name     string // 资源名称, 必填, 默认值: 无
	maxTotal int    // 最大总资源数量, 非必填, 默认值: 8
	maxIdle  int    // 最大空闲资源数量, 非必填, 默认值: 8
	minIdle  int    // 最小空闲资源数量, 非必填, 默认值: 0
}

var (
	defaultMaxTotal = 8
	defaultMaxIdle  = 8
	defaultMinIdle  = 0
)

type ResourcePoolConfigBuilder struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

func (b *ResourcePoolConfigBuilder) SetName(name string) error {
	if name == "" {
		return fmt.Errorf("name can not be empty")
	}
	b.name = name
	return nil
}

func (b *ResourcePoolConfigBuilder) SetMaxTotal(maxTotal int) error {
	if maxTotal < 0 {
		return fmt.Errorf("maxTotal can not be 0")
	}
	b.maxTotal = maxTotal
	return nil
}

func (b *ResourcePoolConfigBuilder) SetMaxIdle(maxIdle int) error {
	if maxIdle < 0 {
		return fmt.Errorf("maxIdle can not be 0")
	}
	b.maxIdle = maxIdle
	return nil
}

func (b *ResourcePoolConfigBuilder) SetMinIdle(minIdle int) error {
	if minIdle < 0 {
		return fmt.Errorf("minIdle can not be 0")
	}
	b.minIdle = minIdle
	return nil
}

func (b *ResourcePoolConfigBuilder) Build() (*ResourcePoolConfig, error) {
	if b.name == "" {
		return nil, fmt.Errorf("name can not be empty")
	}
	if b.maxTotal == 0 {
		b.maxTotal = defaultMaxTotal
	}
	if b.maxIdle == 0 {
		b.maxIdle = defaultMaxIdle
	}
	if b.minIdle == 0 {
		b.minIdle = defaultMinIdle
	}
	if b.maxIdle < b.minIdle {
		return nil, fmt.Errorf("maxIdle can not be less than minIdle")
	}
	if b.maxIdle > b.maxTotal {
		return nil, fmt.Errorf("maxIdle can not be greater than maxTotal")
	}
	return &ResourcePoolConfig{
		name:     b.name,
		maxIdle:  b.maxIdle,
		maxTotal: b.maxTotal,
		minIdle:  b.minIdle,
	}, nil
}
