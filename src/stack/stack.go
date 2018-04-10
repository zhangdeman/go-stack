package stack

import (
	"fmt"
	"errors"
)

type Stack struct {
	size int64 //栈容量
	used int64 //已使用容量
	data [] interface{}	//栈数据
}

/**
 * 创建一个栈
 */
func MakeStack(size int64) Stack {
	data := make([]interface{}, size)
	stack := Stack{size:size, data:data}
	return stack
}

/**
 * 栈是否已满,已用空间等于栈容量，即为栈满
 */
func (instance *Stack) IsFull() bool {
	return instance.size == instance.used
}

/**
 * 栈是否为空,已用空间为0即为空
 */
func (instance *Stack) IsEmpty() bool {
	return 0 == instance.used
}

/**
 * 清空栈,硬清除，数据清空
 */
func (instance *Stack) HardClear() {
	instance.data = make([]interface{}, instance.size)
}

/**
 * 清空栈,软清除，数据不清空
 */
func (instance *Stack) SoftClear() {
	instance.used = 0
}

/**
 * 获取栈长度
 */
func (instance *Stack) GetStackLength() int64 {
	return instance.used
}

/**
 * 获取栈容量
 */
func (instance *Stack) GetStackSize() int64 {
	return instance.size
}

/**
 * 获取栈全部数据
 */
func (instance *Stack) GetFullStackData() []interface{} {
	return instance.data
}

/**
 * 获取栈顶元素
 */
func (instance *Stack) GetTopData() interface{}  {
	return instance.data[instance.used - 1]
}

/**
 * 出栈操作
 */
func (instance *Stack) Pop() (interface{}, int64, error) {
	if instance.IsEmpty() {
		err := errors.New("栈已为空，无数据")
		fmt.Println("栈已为空，无数据")
		return nil, 0, err
	}
	data := instance.data[instance.used]
	afterPop := instance.data[0:instance.used]
	instance.used = instance.used - 1
	instance.data = afterPop
	return data, instance.used, nil
}

/**
 * 入栈操作
 */
func (instance *Stack) Push(data interface{}) (int64, error) {
	if instance.IsFull() {
		instance.size = instance.size * 2
		fmt.Println("栈已满，无法压入数据, 自动扩容")
	}
	afterPush := append(instance.data, data)
	instance.data = afterPush
	instance.used = instance.used + 1
	return instance.used, nil
}

/**
 * 修改指定索引值
 */
func (instance *Stack) ModifyIndexValue(index int64, data interface{}) (interface{}, error) {
	if index > instance.size {
		fmt.Println("索引越界")
		return  nil, errors.New("索引越界")
	}
	instance.data[index] = data
	return data, nil
}

/**
 * 获取指定索引值
 */
func (instance *Stack) GetIndexValue(index int64, data interface{}) (interface{}, error) {
	if index > instance.size || index < 0 {
		fmt.Println("索引越界")
		return  nil, errors.New("索引越界")
	}
	return instance.data[index], nil
}