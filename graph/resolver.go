package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/zeimbeekor/go-dapr-gql-service/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	posts []*model.Post
	comments []*model.Comment
	todos []*model.Todo
	users []*model.User
	user *model.User
}
