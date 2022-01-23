package mysql

//用于where参数map[string]interface{}的interface{}构建
//如map[string]interface{}{
//  "Id":WhereOperatorNOTIN([]int{1,2,3}),
//}

//参数为大部分slice类型与string
func WhereOperatorNOTIN(i interface{}) []interface{} { //"notin"
	return []interface{}{"notin", i}
}
func WhereOperatorIN(i interface{}) []interface{} { //"in"
	return []interface{}{"in", i}
}
func UpdateValueRaw(rawstr string) []string {
	return []string{"exp", rawstr}
}

//不进行任何转义
func WhereOperatorRaw(where string) []interface{} {
	return []interface{}{"raw", where}
}
