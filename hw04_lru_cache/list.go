package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	count int
	front *ListItem
	back  *ListItem
}

func (l *list) Len() int {
	return l.count
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	current := &ListItem{v, nil, nil}

	if first := l.front; first == nil {
		l.front = current
		l.back = current
	} else {
		current.Next = first
		first.Prev = current
	}

	l.front = current
	l.count++

	return current
}

func (l *list) PushBack(v interface{}) *ListItem {
	current := &ListItem{v, nil, nil}

	if last := l.back; last == nil {
		l.front = current
		l.back = current
	} else {
		current.Prev = last
		last.Next = current
	}

	l.back = current
	l.count++

	return current
}

func (l *list) Remove(i *ListItem) {
	if i.Prev == nil {
		i.Next.Prev = nil
		l.front = i.Next
	} else {
		i.Prev.Next = i.Next
	}

	if i.Next == nil {
		i.Prev.Next = nil
		l.back = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}

	l.count--
}

func (l *list) MoveToFront(i *ListItem) {
	if l.front != i {
		i.Prev.Next = i.Next
	} else {
		l.front = i.Next
	}

	if l.back != i {
		i.Next.Prev = i.Prev
	} else {
		l.back = i.Prev
	}

	l.front.Prev = i
	i.Next = l.front
	i.Prev = nil
	l.front = i
}

func NewList() List {
	return new(list)
}
