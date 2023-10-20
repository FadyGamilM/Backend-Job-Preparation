package main

import "log"

type Stack struct {
	items []byte
}

func NewStack() *Stack {
	return &Stack{
		items: []byte{},
	}
}

func (s *Stack) Push(val byte) {
	s.items = append(s.items, val)
}

func (s *Stack) Pop() byte {
	removed := s.items[(len(s.items) - 1)]
	s.items = s.items[:(len(s.items) - 1)]
	return removed
}

func backspaceCompare(s string, t string) bool {
	s_store := NewStack()
	s_len := len(s)
	s_idx := 0
	t_store := NewStack()
	t_len := len(t)
	t_idx := 0

	for s_len > 0 {
		if s[s_idx] == byte('#') {
			// if stack isn't empty and we got a #
			if len(s_store.items) != 0 {
				s_store.Pop()
				s_len--
				s_idx++
			} else {
				s_idx++
				s_len--
			}
		} else {
			s_store.Push(s[s_idx])
			s_len--
			s_idx++
		}
	}

	for t_len > 0 {
		if t[t_idx] == byte('#') {
			// if stack isn't empty and we got a #
			if len(t_store.items) != 0 {
				t_store.Pop()
				t_len--
				t_idx++
			} else {
				t_idx++
				t_len--
			}
		} else {
			t_store.Push(t[t_idx])
			t_len--
			t_idx++
		}
	}

	if len(s_store.items) != len(t_store.items) {
		return false
	} else {
		for len(s_store.items) > 0 && len(t_store.items) > 0 {
			s_poped := s_store.Pop()
			t_poped := t_store.Pop()
			if s_poped != t_poped {
				return false
			}
		}
		return true
	}
}

func main() {
	log.Println(backspaceCompare("#", "#"))
}
