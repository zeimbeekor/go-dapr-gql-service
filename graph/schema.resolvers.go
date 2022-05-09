package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/zeimbeekor/go-dapr-gql-service/graph/generated"
	"github.com/zeimbeekor/go-dapr-gql-service/graph/model"
	"github.com/zeimbeekor/go-dapr-gql-service/internal/utils"
)

func (r *mutationResolver) CreatePosts(ctx context.Context) ([]*model.Post, error) {
	var err error

	fmt.Println("... Create Posts ...")

	if body, url, err := utils.Request("posts"); err == nil {
		mPosts := []*model.Post{}
		json.Unmarshal(body, &mPosts)
		for _, v := range mPosts {
			n := rand.Int() % len(utils.PostType)
			createdAt := time.Now().UTC()
			post := &model.Post{
				ID:        v.ID,
				Title:     v.Title,
				Body:      v.Body,
				URL:       fmt.Sprintf("%s/%d", url, v.ID),
				Type:      utils.PostType[n],
				UserID:    v.UserID,
				CreatedAt: createdAt.Format(time.RFC3339Nano),
			}
			r.posts = append(r.posts, post)
		}
		return r.posts, nil
	}
	return nil, err
}

func (r *mutationResolver) CreateComments(ctx context.Context) ([]*model.Comment, error) {
	var err error

	fmt.Println("... Create Comments ...")

	if body, _, err := utils.Request("comments"); err == nil {
		mComments := []*model.Comment{}
		json.Unmarshal(body, &mComments)
		for _, v := range mComments {
			createdAt := time.Now().UTC()
			comment := &model.Comment{
				ID:        v.ID,
				Name:      v.Name,
				Email:     v.Email,
				Body:      v.Body,
				PostID:    v.PostID,
				CreatedAt: createdAt.Format(time.RFC3339Nano),
			}
			r.comments = append(r.comments, comment)
		}
		return r.comments, nil
	}
	return nil, err
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUsers(ctx context.Context) ([]*model.User, error) {
	var err error

	fmt.Println("... Create Users ...")

	if body, _, err := utils.Request("users"); err == nil {
		mUsers := []*model.User{}
		json.Unmarshal(body, &mUsers)
		for _, v := range mUsers {
			createdAt := time.Now().UTC()
			user := &model.User{
				ID:        v.ID,
				Name:      v.Name,
				Username:  v.Username,
				Email:     v.Email,
				Photo:     v.Photo,
				Website:   v.Website,
				Address:   v.Address,
				Company:   v.Company,
				CreatedAt: createdAt.Format(time.RFC3339Nano),
			}
			r.users = append(r.users, user)
		}
		return r.users, nil
	}
	return nil, err
}

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	fmt.Println("... Get Posts ...")
	return r.posts, nil
}

func (r *queryResolver) Comments(ctx context.Context) ([]*model.Comment, error) {
	fmt.Println("... Get Comments ...")
	return r.comments, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	fmt.Println("... Get Todos ...")
	return r.todos, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	fmt.Println("... Get Users ...")
	return r.users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
