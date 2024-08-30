package queue_test

import (
	"testing"

	queues "github.com/kauebonfimm/go-data-structures/queues"
)

func TestQueues(t *testing.T) {
	tests := []struct {
		name       string
		lenght     int
		parameters []string
		expected   []string
	}{
		{
			name:       "Success instance",
			lenght:     2,
			parameters: []string{"1", "2"},
			expected:   []string{"1", "2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := queues.NewQueue(tt.lenght)
			for _, data := range tt.parameters {
				if err := queue.Enqueue(data); err != nil {
					t.Error(err)
					return
				}
			}

			if queue.Lenght() != len(tt.expected) {
				t.Error("The lenght expected is different by stack")
				return
			}

			for _, data := range tt.expected {
				if res, err := queue.Peek(); err != nil || res != data {
					t.Errorf("Error to implement peek, expected: %s, but return %s", data, res)
					return
				}
				res, err := queue.Dequeue()
				if err != nil {
					t.Error(err)
				}
				if res != data {
					t.Errorf("Error to implement pop, expected: %s, but return %s", data, res)
					return
				}
			}
		})
	}
}
