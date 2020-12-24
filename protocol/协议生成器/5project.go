package main

type MSG_PROJECT_tree_getLinePairs struct {
	QueryID uint32
}
type MSG_PROJECT_tree_getLinePairs_result struct {
	QueryResultID uint32
	List          []HtmlKeyValueStr
}
