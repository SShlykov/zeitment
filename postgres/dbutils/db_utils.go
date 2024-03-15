package dbutils

type QueryOptions interface {
	Pagination
	Filter
}
type Pagination interface {
	GetPagination() (limit uint64, offset uint64)
}
type Filter interface {
	GetFilter() (key string, value any)
}

func NewQueryOptions(filter FilterImpl, pagination PaginationImpl) QueryOptions {
	return queryOptions{filter: filter, pagination: pagination}
}

type queryOptions struct {
	pagination PaginationImpl
	filter     FilterImpl
}

func (q queryOptions) GetFilter() (key string, value any) {
	return q.filter.Key, q.filter.Value
}

func (q queryOptions) GetPagination() (limit uint64, offset uint64) {
	return q.pagination.Limit, q.pagination.Offset
}

type PaginationImpl struct {
	Limit  uint64
	Offset uint64
}

func NewPagination(pageOptions PagedOptions) PaginationImpl {
	page, pageSize := pageOptions.GetPageAndPageSize()
	return PaginationImpl{Limit: pageSize, Offset: (page - 1) * pageSize}
}

type FilterImpl struct {
	Key   string
	Value any
}

func NewFilter(key string, value any) FilterImpl {
	return FilterImpl{Key: key, Value: value}
}

func (p PaginationImpl) GetPagination() (limit uint64, offset uint64) {
	return p.Limit, p.Offset
}

type PagedOptions interface {
	GetPageAndPageSize() (page uint64, pageSize uint64)
}
