package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddTask(t *testing.T) {
	t.Run("Successfully adds task 'Task1'", func(t *testing.T) {
		data := []task{}
		var id int
		var err error
		data, id, err = addTask(data, "Task1")
		require.Nil(t, err)
		require.Equal(t, 1, id)
		require.Equal(t, "Task1", data[0].Description)
		require.Equal(t, StatusToDo, data[0].Status)
	})
}

func TestUpdateTask(t *testing.T) {
	t.Run("Successfully updates task 'Task1'", func(t *testing.T) {
		data := []task{
			{ID: 1, Description: "Task1", Status: StatusToDo, CreatedAt: 0, UpdatedAt: 0},
			{ID: 2, Description: "Task2", Status: StatusToDo, CreatedAt: 0, UpdatedAt: 0},
		}
		var err error
		data, err = updateTask(data, 1, "TestTask1")
		require.Nil(t, err)
		require.Equal(t, "TestTask1", data[0].Description)
		require.NotEqual(t, 0, data[0].UpdatedAt)
	})
}
