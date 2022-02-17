package config

//datatable用的设置定义
type ConfigDatatable struct {
	Order int
	Id    string `json:"id"`
	Show  bool   `json:"show"`
	Width string `json:"Width"`
	Fixed string `json:"Fixed"`
	Title string `json:"Title"`
	Sort  string `json:"Sort"`
	Name  string `json:"Name"`
}
