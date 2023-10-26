package request

// PageQry 分页参数
type PageQry struct {
	Page     int32
	PageSize int32
	OrderBy  string
}
