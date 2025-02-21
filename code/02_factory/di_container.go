package _2_factory

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

// Container 是一个简单的 IoC 容器
type Container struct {
	beans sync.Map // 存储所有注册的 Bean
}

// NewContainer 创建一个新的容器
func NewContainer() *Container {
	return &Container{}
}

// RegisterBean 注册一个 Bean 到容器中
func (c *Container) RegisterBean(name string, bean interface{}) {
	c.beans.Store(name, bean)
}

// GetBean 从容器中获取一个 Bean
func (c *Container) GetBean(name string) (interface{}, error) {
	bean, ok := c.beans.Load(name)
	if !ok {
		return nil, errors.New("bean not found")
	}
	return bean, nil
}

// InjectBean 实现依赖注入
func (c *Container) InjectBean(target interface{}) error {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() != reflect.Ptr || targetValue.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("target must be a pointer to a struct")
	}

	targetValue = targetValue.Elem()
	targetType := targetValue.Type()

	for i := 0; i < targetValue.NumField(); i++ {
		field := targetValue.Field(i)
		fieldType := targetType.Field(i)

		// 检查字段是否有 `inject` 标签
		tag := fieldType.Tag.Get("inject")
		if tag == "" {
			continue
		}

		// 从容器中获取依赖的 Bean
		bean, err := c.GetBean(tag)
		if err != nil {
			return fmt.Errorf("failed to inject field %s: %v", fieldType.Name, err)
		}

		// 将 Bean 注入到字段中
		beanValue := reflect.ValueOf(bean)
		if beanValue.Type().AssignableTo(field.Type()) {
			field.Set(beanValue)
		} else {
			return fmt.Errorf("type mismatch for field %s: expected %s, got %s", fieldType.Name, field.Type(), beanValue.Type())
		}
	}

	return nil
}
