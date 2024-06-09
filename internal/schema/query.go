package schema

import (
	"github.com/anduckhmt146/graphql-api/internal/dtos"
	"github.com/anduckhmt146/graphql-api/internal/services"
	"github.com/graphql-go/graphql"
)

func NewQueryType(userService services.UserService) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"users": &graphql.Field{
				Type: graphql.NewList(dtos.UserType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return userService.GetAllUsers()
				},
			},
			"user": &graphql.Field{
				Type: dtos.UserType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return userService.GetUserByID(p.Args["id"].(int))
				},
			},
		},
	})
}
