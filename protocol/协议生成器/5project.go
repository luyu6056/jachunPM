package main

type MSG_PROJECT_tree_getLinePairs struct {
	QueryID uint32
}
type MSG_PROJECT_tree_getLinePairs_result struct {
	QueryResultID uint32
	List          []HtmlKeyValueStr
}
type MSG_PROJECT_product_cache struct {
	Id          int32
	Name        string
	Code        string
	Line        int32
	Type        string
	Status      string
	Desc        string
	PO          int32
	QD          int32
	RD          int32
	Acl         string
	Whitelist   []int32
	CreatedBy   int32
	CreatedDate int64
	Order       int32
	Deleted     bool
	TimeStamp   int64
}
type MSG_PROJECT_product_insert struct {
	QueryID uint32
	DocName string
	Data    *MSG_PROJECT_product_cache
}
type MSG_PROJECT_product_insert_result struct {
	QueryResultID uint32
	ID            int32
}
