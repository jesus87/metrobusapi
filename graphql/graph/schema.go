package graph

import (
	"github.com/graphql-go/graphql"
)

type GraphqlSchema struct {
	Fields graphql.Fields
}

func NewGraphqlSchema() *GraphqlSchema {
	schema := &GraphqlSchema{}

	schema.Fields = graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}

	return schema
}
