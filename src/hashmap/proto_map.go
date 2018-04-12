package hashmap

import "sort"

/**
 * 原生map处理
 */
type ProtoMap struct {
	data map[string] interface{}	//map数据
}

/**
 * 创建一个map
 */
func MakeProtoMap(data map[string]interface{}) ProtoMap {
	protoMap := ProtoMap{
		data:data,
	}
	return protoMap
}

/**
 * 获取map的全部key
 */
func (protoMap *ProtoMap) GetMapKeyList() []string{
	len := len(protoMap.data)
	keyList := make([]string, len)
	return keyList
}

/**
 * key 降序排序
 */
func (protoMap *ProtoMap) SortDesc() []string {
	keyList := protoMap.SortAsc()
	return keyList
}

/**
 * key 升序排序
 */
func (protoMap *ProtoMap) SortAsc() []string {
	keyList := protoMap.GetMapKeyList()
	sort.Strings(keyList)
	return keyList
}