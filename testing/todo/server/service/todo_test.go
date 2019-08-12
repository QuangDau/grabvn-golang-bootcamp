package service

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/xuanit/testing/todo/pb"
	"github.com/xuanit/testing/todo/server/repository/mocks"
)

func TestGetToDo(t *testing.T) {
	mockToDoRep := &mocks.ToDo{}
	toDo := &pb.Todo{}
	req := &pb.GetTodoRequest{Id: "h"}
	mockToDoRep.On("Get", req.Id).Return(toDo, nil)
	service := ToDo{ToDoRepo: mockToDoRep}

	res, err := service.GetTodo(nil, req)

	expectedRes := &pb.GetTodoResponse{Item: toDo}

	assert.Nil(t, err)
	assert.Equal(t, expectedRes, res)
	mockToDoRep.AssertExpectations(t)
}

func TestListToDo(t *testing.T) {
	mockToDoRep := &mocks.ToDo{}
	toDos := []*pb.Todo{}
	req := &pb.ListTodoRequest{Limit: 10, NotCompleted: false}
	mockToDoRep.On("List", req.Limit, req.NotCompleted).Return(toDos, nil)

	service := ToDo{ToDoRepo: mockToDoRep}

	res, err := service.ListTodo(nil, req)

	expectedRes := &pb.ListTodoResponse{Items: toDos}

	assert.Nil(t, err)
	assert.Equal(t, expectedRes, res)
	mockToDoRep.AssertExpectations(t)
}

func TestCreateToDo(t *testing.T) {
	mockToDoRep := &mocks.ToDo{}
	toDo := &pb.Todo{}

	id, _ := uuid.NewV4()
	toDo.Id = id.String()
	req := &pb.CreateTodoRequest{Item: toDo}
	mockToDoRep.On("Insert", req.Item).Return(nil)

	service := ToDo{ToDoRepo: mockToDoRep}

	res, err := service.CreateTodo(nil, req)

	expectedRes := &pb.CreateTodoResponse{Id: toDo.Id}

	assert.Nil(t, err)
	assert.Equal(t, expectedRes, res)
	mockToDoRep.AssertExpectations(t)
}

func TestDeleteToDo(t *testing.T) {
	mockToDoRep := &mocks.ToDo{}
	toDo := &pb.Todo{}

	id, _ := uuid.NewV4()
	toDo.Id = id.String()
	req := &pb.DeleteTodoRequest{Id: toDo.Id}
	mockToDoRep.On("Delete", id.String()).Return(nil)

	service := ToDo{ToDoRepo: mockToDoRep}

	res, err := service.DeleteTodo(nil, req)

	expectedRes := &pb.DeleteTodoResponse{}

	assert.Nil(t, err)
	assert.Equal(t, expectedRes, res)
	mockToDoRep.AssertExpectations(t)
}
