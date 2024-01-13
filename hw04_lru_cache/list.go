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

type list struct {
	Head *ListItem
	Tail *ListItem
	len  int
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.Head
}

func (l *list) Back() *ListItem {
	return l.Tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	if l.Head == nil {
		l.init(v)
	} else {
		tmp := l.Head
		l.Head = &ListItem{Value: v}
		l.Head.Next = tmp
		tmp.Prev = l.Head
	}
	l.len++
	return l.Head
}

func (l *list) PushBack(v interface{}) *ListItem {
	if l.Tail == nil {
		l.init(v)
	} else {
		tmp := l.Tail
		l.Tail = &ListItem{Value: v}
		l.Tail.Prev = tmp
		tmp.Next = l.Tail
	}
	l.len++
	return l.Tail
}

func (l *list) Remove(i *ListItem) {
	if i.Prev == nil {
		l.Head = i.Next
	} else {
		i.Prev.Next = i.Next
	}
	if i.Next == nil {
		l.Tail = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if l.Head == i {
		return
	}

	if i.Next == nil {
		l.Tail = i.Prev
		l.Tail.Next = nil
		i.Next = l.Head
		l.Head.Prev = i
		i.Prev = nil
		l.Head = i
		if l.Tail.Prev == nil {
			l.Tail.Prev = l.Head
		}
	} else {
		if i.Prev != nil {
			i.Prev.Next = i.Next
		}
		i.Next.Prev = i.Prev
		i.Next = l.Head
		i.Prev = nil
		l.Head = i
	}
}

func (l *list) init(v interface{}) {
	l.Head = &ListItem{Value: v}
	l.Tail = l.Head
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

func NewList() List {
	return new(list)
}
