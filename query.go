package druid

import (
	"fmt"
	"net/http"

	"github.com/grafadruid/go-druid/builder"
	"github.com/grafadruid/go-druid/builder/query"
)

const (
	NativeQueryEndpoint = "druid/v2"
	SQLQueryEndpoint    = "druid/v2/sql"
)

type QueryService struct {
	client *Client
}

func (q *QueryService) Execute(qry builder.Query, result interface{}, headers http.Header) (*Response, error) {
	var path string
	switch qry.Type() {
	case "sql":
		path = SQLQueryEndpoint
	default:
		path = NativeQueryEndpoint
	}
	fmt.Printf("BEFORE NEW REQUEST\n")
	r, err := q.client.NewRequest("POST", path, qry)
	fmt.Printf("AFTER NEW REQUEST %v\n", err)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		for _, vv := range v {
			r.Header.Set(k, vv)
		}
	}
	fmt.Printf("BEFORE CLIENT DO\n")
	resp, err := q.client.Do(r, result)
	fmt.Printf("AFTER CLIENT DO %v\n", err)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//func (q *QueryService) Cancel(query builder.Query) () {}

//func (q *QueryService) Candidates(query builder.Query, result interface{}) (*Response, error) {}

func (q *QueryService) Load(data []byte) (builder.Query, error) {
	return query.Load(data)
}
