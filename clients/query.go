package clients

import "github.com/flusaka/pandascore-go/clients/queries"

//
//func (r *Request) WithFilter(filter Filter) *Request {
//	if r.Filter != nil {
//		r.Filter.Merge(filter)
//	} else {
//		r.Filter = filter
//	}
//	return r
//}

func (r *Request) WithRange(rangeQuery queries.Range) *Request {
	r.rangeQuery = rangeQuery
	return r
}

func (r *Request) WithSort(sort queries.Sort) *Request {
	r.sortQuery = sort
	return r
}

//func (r *Request) WithSearch(search Search) *Request {
//	if r.Search != nil {
//		r.Search.Merge(search)
//	} else {
//		r.Search = search
//	}
//	return r
//}
