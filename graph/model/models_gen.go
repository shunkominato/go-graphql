// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Node interface {
	IsNode()
	GetID() string
}

type Pagination interface {
	IsPagination()
	GetPageInfo() *PaginationInfo
	GetNodes() []Node
}

type CreateTodoInput struct {
	Todo         string `json:"todo"`
	TodoStatusID int    `json:"todoStatusId"`
	UserID       int    `json:"userId"`
}

type PaginationInfo struct {
	Page             int  `json:"page"`
	PaginationLength int  `json:"paginationLength"`
	HasNextPage      bool `json:"hasNextPage"`
	HasPreviousPage  bool `json:"hasPreviousPage"`
	Count            int  `json:"count"`
	TotalCount       int  `json:"totalCount"`
}

type Todo struct {
	ID         string      `json:"id"`
	Todo       string      `json:"todo"`
	TodoStatus *TodoStatus `json:"todoStatus"`
	User       *User       `json:"user"`
}

func (Todo) IsNode()            {}
func (this Todo) GetID() string { return this.ID }

type TodoPagination struct {
	PageInfo *PaginationInfo `json:"pageInfo"`
	Nodes    []*Todo         `json:"nodes"`
}

func (TodoPagination) IsPagination()                     {}
func (this TodoPagination) GetPageInfo() *PaginationInfo { return this.PageInfo }
func (this TodoPagination) GetNodes() []Node {
	if this.Nodes == nil {
		return nil
	}
	interfaceSlice := make([]Node, 0, len(this.Nodes))
	for _, concrete := range this.Nodes {
		interfaceSlice = append(interfaceSlice, concrete)
	}
	return interfaceSlice
}

type TodoStatus struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type UpdateTodoInput struct {
	ID           int    `json:"id"`
	Todo         string `json:"todo"`
	TodoStatusID int    `json:"todoStatusId"`
	UserID       int    `json:"userId"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
