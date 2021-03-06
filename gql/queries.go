package gql

import (
	"github.com/austen-wade/go-graphql/postgres"
	"github.com/graphql-go/graphql"
)

type Root struct {
	Query *graphql.Object
}

func NewRoot(db *postgres.Db) *Root {
	resolver := Resolver{db: db}

	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"users": &graphql.Field{
						Type: graphql.NewList(User),
						Args: graphql.FieldConfigArgument{
							"name": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
							"id": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: resolver.UserResolver,
					},
				},
			},
		),
	}
	return &root
}
