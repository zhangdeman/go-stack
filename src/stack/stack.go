package stack

import (
	"fmt"
	"errors"
)

type Stack struct {
	size int64 //栈容量
	used int64 //已使用容量
	isExtend bool //栈满是否允许自动扩展
	data [] interface{}	//栈数据
}

/**
 * 创建一个栈
 * @param size 栈大小
 * @param isExtend 是否允许扩展
 */
func MakeStack(size int64, isExtend bool) Stack {
	data := make([]interface{}, size)
	stack := Stack{
		size:size,
		data:data[0:1],
		used:0,
		isExtend:isExtend,
	}
	return stack
}

/**
 * 栈是否已满,已用空间等于栈容量，即为栈满
 */
func (stack *Stack) IsFull() bool {
	return stack.size == stack.used
}

/**
 * 栈是否为空,已用空间为0即为空
 */
func (stack *Stack) IsEmpty() bool {
	return 0 == stack.used
}

/**
 * 清空栈,硬清除，数据清空
 */
func (stack *Stack) HardClear() {
	stack.data = make([]interface{}, stack.size)
}

/**
 * 清空栈,软清除，数据不清空
 */
func (stack *Stack) SoftClear() {
	stack.used = 0
}

/**
 * 获取栈长度
 */
func (stack *Stack) GetStackLength() int64 {
	return stack.used
}

/**
 * 获取栈容量
 */
func (stack *Stack) GetStackSize() int64 {
	return stack.size
}

/**
 * 获取栈全部数据
 */
func (stack *Stack) GetFullStackData() []interface{} {
	return stack.data
}

/**
 * 获取栈顶元素
 */
func (stack *Stack) GetTopData() interface{}  {
	return stack.data[stack.used - 1]
}

/**
 * 出栈操作
 */
func (stack *Stack) Pop() (interface{}, int64, error) {
	if stack.IsEmpty() {
		err := errors.New("栈已为空，无数据")
		fmt.Println("栈已为空，无数据")
		return nil, 0, err
	}
	data := stack.data[stack.used]
	afterPop := stack.data[0:stack.used]
	stack.used = stack.used - 1
	stack.data = afterPop
	return data, stack.used, nil
}

/**
 * 入栈操作
 */
func (stack *Stack) Push(data interface{}) (int64, error) {
	if stack.IsFull() {
		if stack.isExtend {
			//允许扩展
			stack.size = stack.size * 2
			fmt.Println("栈已满，无法压入数据, 自动扩容")
		} else {
			//不允许扩展
			return -1, errors.New("栈已满，定长队列不允许自动扩展")
		}
	}
	afterPush := append(stack.data, data)
	stack.data = afterPush
	stack.used = stack.used + 1
	return stack.used, nil
}

/**
 * 修改指定索引值
 */
func (stack *Stack) ModifyIndexValue(index int64, data interface{}) (interface{}, error) {
	if index > stack.size {
		fmt.Println("索引越界")
		return  nil, errors.New("索引越界")
	}
	stack.data[index] = data
	return data, nil
}

/**
 * 获取指定索引值
 */
func (stack *Stack) GetIndexValue(index int64, data interface{}) (interface{}, error) {
	if index > stack.size || index < 0 {
		fmt.Println("索引越界")
		return  nil, errors.New("索引越界")
	}
	return stack.data[index], nil
}

/**
 * 逆转一个栈
 * @return slice
 */
func (stack *Stack) ReversalStack() []interface{} {
	stackLen := len(stack.data)
	reversalStack := make([]interface{}, 1)
	for index := stackLen - 1; index > 0; index-- {
		reversalStack = append(reversalStack, stack.data[index])
	}
	return reversalStack
}