package stack_test

import (
	"testing"

	stacks "github.com/kauebonfimm/go-data-structures/stacks"
)

func TestStack(t *testing.T) {
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
			expected:   []string{"2", "1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := stacks.NewStack(tt.lenght)
			for _, data := range tt.parameters {
				if err := stack.Push(data); err != nil {
					t.Error(err)
					return
				}
			}

			if stack.Lenght() != len(tt.expected) {
				t.Error("The lenght expected is different by stack")
				return
			}

			for _, data := range tt.expected {
				if res, err := stack.Peek(); err != nil || res != data {
					t.Errorf("Error to implement peek, expected: %s, but return %s", data, res)
					return
				}
				res, err := stack.Pop()
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
