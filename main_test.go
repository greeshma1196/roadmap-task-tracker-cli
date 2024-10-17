package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddTask(t *testing.T) {
	t.Run("Successfully adds task", func(t *testing.T) {
		data := []task{}
		var id int
		var err error
		data, id, err = addTask(data, "Task1")
		require.Nil(t, err)
		require.NotEqual(t, 0, id)
		require.NotEqual(t, 0, len(data))
	})
}
