package main

import "fmt"

type node struct {
	name string
	next *node
}

type Singlelinkedlist struct {
	len  int
	head *node
}

func initList() *Singlelinkedlist {
	return &Singlelinkedlist{}
}

func (s *Singlelinkedlist) AddFront(name string) {
	node := &node{
		name: name,
	}
	if s.head == nil {
		s.head = node
	} else {
		node.next = s.head
		s.head = node
	}
	s.len++
	return
}

func (s *Singlelinkedlist) AddBack(name string) {
	node := &node{
		name: name,
	}
	if s.head == nil {
		s.head = node
	} else {
		current := s.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}
	s.len++
	return
}

func (s *Singlelinkedlist) RemoveFront() error {
	if s.head == nil {
		return fmt.Errorf("List is empty")
	}
	s.head = s.head.next
	s.len--
	return nil
}

func (s *Singlelinkedlist) RemoveBack() error {
	if s.head == nil {
		return fmt.Errorf("removeBack: List is empty")
	}
	var prev *node
	current := s.head
	for current.next != nil {
		prev = current
		current = current.next
	}
	if prev != nil {
		prev.next = nil
	} else {
		s.head = nil
	}
	s.len--
	return nil
}

func (s *Singlelinkedlist) Front() (string, error) {
	if s.head == nil {
		return "", fmt.Errorf("Single List is empty")
	}
	return s.head.name, nil
}

func (s *Singlelinkedlist) Size() int {
	return s.len
}

func (s *Singlelinkedlist) Traverse() error {
	if s.head == nil {
		return fmt.Errorf("TranverseError: List is empty")
	}
	current := s.head
	for current != nil {
		fmt.Println(current.name)
		current = current.next
	}
	return nil
}

func main() {
	Singlelinkedlist := initList()
	fmt.Printf("AddFront: A\n")
	Singlelinkedlist.AddFront("A")
	fmt.Printf("AddFront: B\n")
	Singlelinkedlist.AddFront("B")
	fmt.Printf("AddBack: C\n")
	Singlelinkedlist.AddBack("C")

	fmt.Printf("Size: %d\n", Singlelinkedlist.Size())

	err := Singlelinkedlist.Traverse()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("RemoveFront\n")
	err = Singlelinkedlist.RemoveFront()
	if err != nil {
		fmt.Printf("RemoveFront Error: %s\n", err.Error())
	}

	fmt.Printf("RemoveBack\n")
	err = Singlelinkedlist.RemoveBack()
	if err != nil {
		fmt.Printf("RemoveBack Error: %s\n", err.Error())
	}

	fmt.Printf("RemoveBack\n")
	err = Singlelinkedlist.RemoveBack()
	if err != nil {
		fmt.Printf("RemoveBack Error: %s\n", err.Error())
	}

	fmt.Printf("RemoveBack\n")
	err = Singlelinkedlist.RemoveBack()
	if err != nil {
		fmt.Printf("RemoveBack Error: %s\n", err.Error())
	}

	err = Singlelinkedlist.Traverse()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Size: %d\n", Singlelinkedlist.Size())

}
