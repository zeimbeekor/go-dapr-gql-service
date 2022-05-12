package utils

import (
	"io/ioutil"
	"fmt"
	"os"
	"bytes"
	"math/rand"
	"time"
	"strconv"
	"net/http"

	"github.com/zeimbeekor/go-dapr-gql-service/graph/model"
)

const (
	USE_BINDING      = false
	HTTP_URL         = "https://jsonplaceholder.typicode.com"
	HTTP_BINDING_URL = "http://127.0.0.1:3500/v1.0/bindings/jsonbinding"
	QUERY_TEMPLATE   = `{
		"operation": "get",
		"metadata": {
			"path": "/%s"
		}
	}`
)

var client = &http.Client{}
var PostType = []string{"story", "job", "poll"}

// Generic request
func Request(path string) ([]byte, string, error) {
	var (
		url       string
		err       error
		reqMethod string
		reqData   bytes.Buffer
	)
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	useBinding, err := strconv.ParseBool(os.Getenv("USE_BINDING"))
	if err != nil {
		return nil, "", err
	}
	switch useBinding {
	case false:
		url = os.Getenv("HTTP_URL")
		if url == "" {
			url = HTTP_URL
		}
		url = url + "/" + path
		reqMethod = "GET"
	default:
		url = os.Getenv("HTTP_BINDING_URL")
		if url == "" {
			url = HTTP_BINDING_URL
		}
		reqMethod = "POST"
		q := fmt.Sprintf(QUERY_TEMPLATE, path)
		reqData = *bytes.NewBuffer([]byte(q))
	}
	if req, err := http.NewRequest(reqMethod, url, &reqData); err == nil {
		req.Header.Add("Content-Type", "application/json")
		// req.Header.Add("dapr-app-id", "gql")
		if resp, err := client.Do(req); err == nil {
			defer resp.Body.Close()
			if body, err := ioutil.ReadAll(resp.Body); err == nil {
				return body, url, nil
			}
		}
	}
	return nil, "", err
}

// Set todo
func SetTodo(todo *model.Todo, user *model.User) *model.Todo {
	m := &model.Todo{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
		User:      user,
	}
	return m
}

// Set user
func SetUser(user *model.User) {
	createdAt := time.Now().UTC()
	url := fmt.Sprintf("%s%s", "https://i.pravatar.cc/100?u=", user.Email)
	user.Photo = url
	user.CreatedAt = createdAt.Format(time.RFC3339Nano)
}

// Add posts user
func AddPostsUser(user *model.User, posts []*model.Post) {
	user.Posts = posts
}

// Set post
func SetPost(post *model.Post, url string) {
	createdAt := time.Now().UTC()
	n := rand.Int() % len(PostType)
	post.URL = fmt.Sprintf("%s/%d", url, post.ID)
	post.Type = PostType[n]
	post.CreatedAt = createdAt.Format(time.RFC3339Nano)
}

// Add comments post
func AddCommentsPost(post *model.Post, comments []*model.Comment) {
	post.Comments = comments
}

// Set comment
func SetComment(comment *model.Comment) {
	createdAt := time.Now().UTC()
	comment.CreatedAt = createdAt.Format(time.RFC3339Nano)
}

// Generic slice filter
func Filter(data []map[string]interface{}, f filterFunc) []map[string]interface{} {
	filtered := []map[string]interface{}{}
	for _, d := range data {
		if f(d) {
			filtered = append(filtered, d)
		}
	}
	return filtered
}

type filterFunc func(d map[string]interface{}) bool