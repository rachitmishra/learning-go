package todo

import (
	"context"

	pb "rachitmishra.com/yc/generated/data/proto/todo"
)

type Handler struct {
	pb.UnimplementedTodoServer
}

func (h *Handler) ReadTodos(ctx context.Context, v *pb.Void) (*pb.TodoItems, error) {
	return &pb.TodoItems{
		Items: []*pb.TodoItem{
			{
				Id:   1,
				Text: "One",
			},
			{
				Id:   2,
				Text: "Two",
			},
		},
	}, nil
}

func (h *Handler) CreateTodo(ctx context.Context, item *pb.TodoItem) (*pb.TodoItem, error) {
	return &pb.TodoItem{
		Id:   1,
		Text: "Created",
	}, nil
}
