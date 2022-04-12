package stack

import (
	"testing"
)

func TestStack_Peek(t *testing.T) {
	testStack := NewStack[rune]()
	testStack.Push('a')
	testStack.Push('b')
	testStack.Push('c')

	tests := []struct {
		name  string
		stack *Stack[rune]
		want  func(any) bool
	}{
		{
			name:  "peek top of stack with three layers",
			stack: testStack,
			want: func(v any) bool {
				return v == 'c' && testStack.Length == 3
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.stack.Peek()
			if err != nil {
				t.Errorf("Stack.Peek() error = %v", err)
				return
			}

			if !tt.want(got) {
				t.Errorf("Stack.Peek() error = %v", err)
				return
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	testStack := NewStack[rune]()
	testStack.Push('a')
	testStack.Push('b')
	testStack.Push('c')

	tests := []struct {
		name  string
		Stack *Stack[rune]
		want  func(any) bool
	}{
		{
			name:  "pop from top of stack",
			Stack: testStack,
			want: func(s any) bool {
				return s == 'c' && testStack.Length == 2
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.Stack.Pop()
			if err != nil {
				t.Errorf("Stack.Pop() error = %v", err)
			}

			if !tt.want(got) {
				t.Errorf("Stack.Pop() error")
				return
			}
		})
	}
}

func TestStack_Push(t *testing.T) {
	testStack := NewStack[rune]()
	testStack.Push('a')
	testStack.Push('b')
	testStack.Push('c')

	type args struct {
		data rune
	}
	tests := []struct {
		name  string
		stack *Stack[rune]
		args  args
		want  func(*Stack[rune]) bool
	}{
		{
			name:  "push to stack",
			stack: testStack,
			args: args{
				data: 'd',
			},
			want: func(s *Stack[rune]) bool {
				got, _ := s.Peek()
				return got == 'd' && s.Length == 4
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.stack.Push(tt.args.data)
			if err != nil {
				t.Errorf("Stack.Push() error = %v", err)
				return
			}

			if !tt.want(tt.stack) {
				t.Errorf("Stack.Push() error")
			}
		})
	}
}
