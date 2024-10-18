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
