package graph

import (
	"github.com/graphql-go/graphql"
	"github.com/mszsgo/hgraph"
)

var (
	GraphqlHttpHandler = hgraph.GraphqlHttpHandler
	QueryFields        = hgraph.MergeQueryFields
	MutationFields     = hgraph.MergeMutationFields

	String           = graphql.String
	NewNonNullString = graphql.NewNonNull(graphql.String)
)

// 字段响应模型，针对有错误码的字段
type FieldResponseModel struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
