package hashmap

import (
	"errors"
)

/**
 * map数据结构
 */
type MapType struct {
	index int64
	key string
	value interface{}
}

/**
 * 存储的hashmap
 */
type HashMap struct {
	data []MapType
	keyList map[string]int64	//记录每一个key对应的索引，删除使用
}

/**
 * 创建一个hashmap
 */
func MakeHashMap() HashMap {
	keyList := make(map[string]int64)
	hashMap := HashMap{
		keyList:keyList,
	}
	return hashMap
}

/**
 * 获取map的长度
 */
func (hashMap *HashMap) GetMapLength() int64 {
	return int64(len(hashMap.GetAllData()))
}

/**
 * 获取map的value
 */
func (hashMap *HashMap) GetMapValue(key string) (interface{},int64, error)  {
	index, ok := hashMap.keyList[key]
	if !ok {
		//不存在
		return nil, -1, errors.New(key + " 索引不存在")
	}
	//存在
	data := hashMap.data[index].value
	return data, index, nil
}

/**
 * 添加map数据,不存在，追加新数据，存在，值覆盖
 * @return 添加后map的长度
 */
func (hashMap *HashMap) Add(key string, value interface{}) int64 {
	//先检测key是否已存在，若存在，值覆盖
	_,index, err := hashMap.GetMapValue(key)
	len := hashMap.GetMapLength()
	if err != nil {
		//说明不存在，追加即可
		data := MapType{
			index:len,
			key:key,
			value:value,
		}
		hashMap.data = append(hashMap.data, data)
		//维护keyList
		hashMap.keyList[key] = len
		return len + 1
	}

	//存在，修改原有值
	hashMap.data[index].value = value
	return len
}

/**
 * 获取全部的数据
 */
func (hashMap *HashMap) GetAllData() []MapType {
	return hashMap.data
}

/**
 * 删除一个key
 */
func (hashMap *HashMap) Delete(key string) []MapType {
	_,index,err := hashMap.GetMapValue(key)
	if err != nil {
		//不存在,不需要删除，直接返回
		return hashMap.GetAllData()
	}
	//存在，需要删除
	//1)删除keyMap
	delete(hashMap.keyList, key)

	//2)后续map索引修改，后续的前进补位，保证索引是连续的，不会断裂，同时后面的数据迁移
	len := hashMap.GetMapLength()
	for i := index + 1; i < len; i++ {
		hashMap.keyList[hashMap.data[i].key] = i - 1
		hashMap.data[i].index = i - 1
		//后面数据前移
		hashMap.data[i - 1] = hashMap.data[i]
	}

	//3) 删除最后一条数据
	hashMap.data = hashMap.data[0:len - 1]
	return hashMap.GetAllData()
}

/**
 * 批量删除指定key
 */
func (hashMap *HashMap) BatchDelete(key ...string) []MapType {
	data := hashMap.GetAllData()
	for _, mapKey := range key{
		data = hashMap.Delete(mapKey)
	}
	return data
}