package graphqlTypes

import "github.com/graphql-go/graphql"

//TeacherType Objeto GraphQL para profesores.
var TeacherType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Teacher",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"userID": &graphql.Field{
				Type: graphql.Int,
			},
			"firstName": &graphql.Field{
				Type: graphql.String,
			},
			"lastName": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"phoneNumber": &graphql.Field{
				Type: graphql.String,
			},
			"ciType": &graphql.Field{
				Type: graphql.String,
			},
			"ci": &graphql.Field{
				Type: graphql.Int,
			},
			"message": &graphql.Field{
				Type: graphql.String,
			},
			"code": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)
