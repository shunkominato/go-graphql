package graph

import "story.com/story/app/internal/infrastructure/server/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.


type Resolver struct{
	todos []*model.Todo
	todoStatus []*model.TodoStatus
	user *model.User
}
