package stack

/**
 * 双栈模拟队列
 */
type Queue struct {
	size int64 			//队列最大长度
	used int64 			//已使用长度
	isExtend bool		//队列长度是否允许自动扩展
	bufferStack Stack	//缓冲数据栈(只push,不pop)
	queueData Stack		//队列数据栈(只pop,不push)
}

/**
 * 创建一个队列
 */
func MakeQueue(size int64, isExtend bool) Queue {
	queue := Queue{
		size:size,
		used:0,
		isExtend:isExtend,
		bufferStack:MakeStack(size, isExtend),
		queueData:MakeStack(size, isExtend),
	}
	return queue
}

/**
 * 判断队列是否为空
 * 队列数据不为空或者缓冲栈不为空，则队列不为空
 */
func (queue *Queue) IsEmpty() bool {
	if !queue.queueData.IsEmpty() || !queue.bufferStack.IsEmpty() {
		return true
	}
	return false
}

/**
 * 获取队列长度
 * 队列长度 + 缓冲区长度 = 真实队列长度
 */
func (queue *Queue) GetQueueLen() int64 {
	return queue.used
}

/**
 * 获取队列长度
 * 队列长度 + 缓冲区长度 = 真实队列长度
 */
func (queue *Queue) GetQueueSize() int64 {
	return queue.size
}

/**
 * 向队列中追加数据
 * @param data interface 压入队列的数据
 * @return int64 数据压入后，队列的长度
 */
func (queue *Queue) Push(data interface{}) int64 {
	if queue.used == queue.size {
		//队列已用尽自动扩容队列容量至当前2倍
		queue.size = queue.size * 2
	}
	queue.bufferStack.Push(data)
	//队列长度加一
	queue.used = queue.used + 1
	return queue.GetQueueLen()
}

/**
 * 弹出队列中的数据
 * @return interface 弹出的数据
 */
func (queue *Queue) Pop() interface{}  {
	if queue.IsEmpty() {
		//队列为空
		return nil
	}

	isQueueDataEmpty := queue.queueData.IsEmpty()
	if !isQueueDataEmpty {
		//队列数据不为空
		data, _, _ := queue.queueData.Pop()
		return data
	}

	//队列数据为空时
	isBufferEmpty := queue.bufferStack.IsEmpty()
	if isBufferEmpty {
		//缓冲数据为空
		return nil
	}

	//缓冲数据不为空,将缓冲数据全部刷新至队列
	queue.bufferStack.data = queue.bufferStack.ReversalStack()
	queue.queueData = queue.bufferStack
	queue.used = queue.bufferStack.used
	//清空缓冲栈
	queue.bufferStack = MakeStack(queue.size, queue.isExtend)

	//队列长度减一
	queue.used = queue.used - 1

	//取出数据
	data, _, _ := queue.queueData.Pop()
	return data
}
