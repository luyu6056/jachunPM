package protocol

//定义一些http渲染模板的结构体
type HtmlKeyValueStr struct {
	Key   string
	Value string
}
type HtmlKeyValueInterface struct {
	Key   string
	Value interface{}
}
type HtmlMenu struct {
	Key   string
	Value map[string]string
}
type HtmlBlock struct {
	Source string
	Grid   int
	Title  string
	Block  string
	Prams  map[string]string
}

func HtmlKeyValueStr2MapStringInterface(in []HtmlKeyValueStr) (out map[string]interface{}) {
	out = make(map[string]interface{}, len(in))
	for _, kv := range in {
		out[kv.Key] = kv.Value
	}
	return
}
