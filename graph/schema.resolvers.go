package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ledongthuc/goterators"
	"github.com/zeimbeekor/go-dapr-gql-service/graph/generated"
	"github.com/zeimbeekor/go-dapr-gql-service/graph/model"
	"github.com/zeimbeekor/go-dapr-gql-service/internal/utils"
)

// Create todo
func (r *mutationResolver) CreateTodo(ctx context.Context) ([]*model.Todo, error) {
	var (
		url       string
		err       error
		mTodos    = []*model.Todo{}
		mPosts    = []*model.Post{}
		mComments = []*model.Comment{}
		mUsers    = []*model.User{}
	)

	fmt.Println("Downloading todos metadata...")
	if body, _, err := utils.Request("todos"); err == nil {
		json.Unmarshal(body, &mTodos)
	}

	fmt.Println("Downloading users metadata...")
	if body, _, err := utils.Request("users"); err == nil {
		json.Unmarshal(body, &mUsers)
	}

	fmt.Println("Downloading posts metadata...")
	if body, u, err := utils.Request("posts"); err == nil {
		url = u
		json.Unmarshal(body, &mPosts)
	}

	fmt.Println("Downloading comments metadata...")
	if body, _, err := utils.Request("comments"); err == nil {
		json.Unmarshal(body, &mComments)
	}

	fmt.Println("Create comments...")
	for _, comment := range mComments {
		utils.SetComment(comment)
		r.comments = append(r.comments, comment)
	}

	fmt.Println("Create posts...")
	for _, post := range mPosts {
		utils.SetPost(post, url)
		filteredItems := goterators.Filter(r.comments, func(item *model.Comment) bool {
			return item.PostID == post.ID
		})
		utils.AddCommentsPost(post, filteredItems)
		r.posts = append(r.posts, post)
	}

	fmt.Println("Create users...")
	for _, user := range mUsers {
		utils.SetUser(user)
		filteredItems := goterators.Filter(r.posts, func(item *model.Post) bool {
			return item.UserID == user.ID
		})
		utils.AddPostsUser(user, filteredItems)
		r.users = append(r.users, user)
	}

	fmt.Println("Create todos...")
	for _, todo := range mTodos {
		filteredItems := goterators.Filter(r.users, func(item *model.User) bool {
			return item.ID == todo.UserID
		})
		todo := utils.SetTodo(todo, filteredItems[0])
		r.todos = append(r.todos, todo)
	}

	if err != nil {
		return nil, err
	}
	return r.todos, nil
}

// Get posts
func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	fmt.Println("Get posts...")
	if len(r.posts) > 0 {
		return r.posts, nil
	} else {
		return nil, errors.New("posts not found")
	}
}

// Get comments
func (r *queryResolver) Comments(ctx context.Context) ([]*model.Comment, error) {
	fmt.Println("Get comments...")
	if len(r.comments) > 0 {
		return r.comments, nil
	} else {
		return nil, errors.New("comments not found")
	}
}

// Get todos
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	fmt.Println("Get todos...")
	if len(r.todos) > 0 {
		return r.todos, nil
	} else {
		return nil, errors.New("todos not found")
	}
}

// Get users
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	fmt.Println("Get users...")
	if len(r.users) > 0 {
		return r.users, nil
	} else {
		return nil, errors.New("users not found")
	}
}

// Get user
func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
	fmt.Println("Get user...")
	filteredItems := goterators.Filter(r.users, func(item *model.User) bool {
		return item.ID == id
	})
	if len(filteredItems) > 0 {
		return filteredItems[0], nil
	} else {
		return nil, errors.New("user not found")
	}
}

// Get post
func (r *queryResolver) Post(ctx context.Context, id int) (*model.Post, error) {
	fmt.Println("Get post...")
	filteredItems := goterators.Filter(r.posts, func(item *model.Post) bool {
		return item.ID == id
	})
	if len(filteredItems) > 0 {
		return filteredItems[0], nil
	} else {
		return nil, errors.New("post not found")
	}
}

// Get comment
func (r *queryResolver) Comment(ctx context.Context, id int) (*model.Comment, error) {
	fmt.Println("Get comment...")
	filteredItems := goterators.Filter(r.comments, func(item *model.Comment) bool {
		return item.ID == id
	})
	if len(filteredItems) > 0 {
		return filteredItems[0], nil
	} else {
		return nil, errors.New("comment not found")
	}
}

// Get information the user by id
func (r *queryResolver) GetInfoByUserID(ctx context.Context, id int) (*model.User, error) {
	fmt.Println("Get information the user by id...")
	filteredItems := goterators.Filter(r.users, func(item *model.User) bool {
		return item.ID == id
	})
	if len(filteredItems) > 0 {
		return filteredItems[0], nil
	} else {
		return nil, errors.New("user not found")
	}
}

// Get all posts by user id
func (r *queryResolver) GetPostsByUserID(ctx context.Context, id int, postType string) ([]*model.Post, error) {
	fmt.Println("Get all posts by user id...")
	filteredItems := goterators.Filter(r.posts, func(item *model.Post) bool {
		return (item.Type == postType || postType == "") && item.UserID == id
	})
	if len(filteredItems) > 0 {
		return filteredItems, nil
	} else {
		return nil, errors.New("we did not find posts for the user")
	}
}

// Get all comments by post id
func (r *queryResolver) GetCommentsByPostID(ctx context.Context, id int) ([]*model.Comment, error) {
	fmt.Println("Get all comments by post id...")
	filteredItems := goterators.Filter(r.comments, func(item *model.Comment) bool {
		return item.PostID == id
	})
	if len(filteredItems) > 0 {
		return filteredItems, nil
	} else {
		return nil, errors.New("we did not find comments for the post")
	}
}

// Mutation returns generated.MutationResolver implementation
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
