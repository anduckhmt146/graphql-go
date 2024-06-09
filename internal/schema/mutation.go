package schema

import (
	"github.com/anduckhmt146/graphql-api/internal/dtos"
	"github.com/anduckhmt146/graphql-api/internal/services"
	"github.com/graphql-go/graphql"
)

func NewMutationType(userService services.UserService) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createUser": &graphql.Field{
				Type: dtos.UserType,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"age": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return userService.CreateUser(p.Args["name"].(string), p.Args["age"].(int))
				},
			},
			"updateUser": &graphql.Field{
				Type: dtos.UserType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"age": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return userService.UpdateUser(p.Args["id"].(int), p.Args["name"].(string), p.Args["age"].(int))
				},
			},
			"deleteUser": &graphql.Field{
				Type: graphql.Int,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return userService.DeleteUser(p.Args["id"].(int))
				},
			},
		},
	})
}
