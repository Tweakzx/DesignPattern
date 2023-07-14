package di

import (
	"fmt"
	"reflect"
)

// Container DI 容器
type Container struct {
	// 假设一种类型只能有一个 provider 提供， 存储各种类型的工厂类
	providers map[reflect.Type]provider

	// 缓存已生成的对象， 可以看做是原材料的仓库
	results map[reflect.Type]reflect.Value
}

type provider struct {
	value reflect.Value

	params []reflect.Type
}

// New 创建一个容器
func New() *Container {
	return &Container{
		providers: map[reflect.Type]provider{},
		results:   map[reflect.Type]reflect.Value{},
	}
}

// isError 判断是否是 error 类型
func isError(t reflect.Type) bool {
	if t.Kind() != reflect.Interface {
		return false
	}
	return t.Implements(reflect.TypeOf(reflect.TypeOf((*error)(nil)).Elem()))
}

// Provide 对象提供者，需要传入一个对象的工厂方法，后续会用于对象的创建
func (c *Container) Provide(constructor interface{}) error {
	v := reflect.ValueOf(constructor)

	// 仅支持函数 provider
	if v.Kind() != reflect.Func {
		return fmt.Errorf("constructor must be a func")
	}

	vt := v.Type()

	// 获取工厂生产所需的原材料
	params := make([]reflect.Type, vt.NumIn())
	for i := 0; i < vt.NumIn(); i++ {
		params[i] = vt.In(i)
	}

	// 获取可以生产的产品列表
	results := make([]reflect.Type, vt.NumOut())
	for i := 0; i < vt.NumOut(); i++ {
		results[i] = vt.Out(i)
	}

	provider := provider{
		value:  v,      // 工厂本体
		params: params, // 工厂所需的原材料参数
	}

	// 根据产品类型保存不同的工厂
	for _, result := range results {
		// 判断返回值是不是 error
		if isError(result) {
			continue
		}

		if _, ok := c.providers[result]; ok {
			return fmt.Errorf("%s had a provider", result)
		}

		c.providers[result] = provider
	}

	return nil
}

// Invoke 函数执行入口
func (c *Container) Invoke(function interface{}) error {
	v := reflect.ValueOf(function)

	// 判断是否是函数
	if v.Kind() != reflect.Func {
		return fmt.Errorf("constructor must be a func")
	}

	vt := v.Type()

	// 获取生产所需的原材料
	var err error
	params := make([]reflect.Value, vt.NumIn())
	for i := 0; i < vt.NumIn(); i++ {
		// buildParam 会递归生产以获取原材料
		params[i], err = c.buildParam(vt.In(i))
		if err != nil {
			return err
		}
	}

	//原材料获取完成， 执行生产
	v.Call(params)

	return nil
}

func (c *Container) buildParam(param reflect.Type) (val reflect.Value, err error) {
	// 如果已经生成了原材料，直接返回
	if result, ok := c.results[param]; ok {
		return result, nil
	}

	// 获取原材料的生产工厂
	provider, ok := c.providers[param]
	if !ok {
		return reflect.Value{}, fmt.Errorf("can not found provider: %s", param)
	}

	// 递归获取生成原材料的原材料， 以生成原材料
	params := make([]reflect.Value, len(provider.params))
	for i, p := range provider.params {
		params[i], _ = c.buildParam(p)
	}

	// 执行工厂生产，生产原材料， 存储到c.results
	results := provider.value.Call(params)
	for _, result := range results {
		// 判断是否报错
		if isError(result.Type()) && !result.IsNil() {
			return reflect.Value{}, result.Interface().(error)
		}

		if !isError(result.Type()) && !result.IsNil() {
			c.results[result.Type()] = result
		}
	}

	// 返回原材料
	return c.results[param], nil
}

//
// +---------------------------------------------------+
// |  *Container          +-----------------+          |
// |                      | providers       |          |
// |              +------>|                 |-----+    |
// |              |       |  A: NewA(B)     |     |    |
// |              |       |  B: NewB(C)     |     |    |
// |      调用工厂 |生产   |  C: NewC()      | 递归|查找 |
// |              |       |                 |     |    |
// |              |       +-----------------+     |    |
// |              |       +-----------------+     |    |
// |              +------ | result          |<----+    |
// |       +----------->  |                 |          |
// |       |              | C: c            |          |
// |       | 从仓库查找A   | B: b            |          |
// |       |              | A: A            |          |
// |    func(A)           |                 |          |
// |       ^              +-----------------+          |
// |       |                                           |
// +-------|-------------------------------------------+
//         |
//         |  Invoke(func(A){})
//         |
//         |
// +--------------------+
// | 外界请求  传入函数  |
// +--------------------+
//
