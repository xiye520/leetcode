package list

import (
	"fmt"
	"testing"
)

type Intrusive interface {
	Next() Intrusive
	Prev() Intrusive
	AddNext(Intrusive)
	AddPrev(Intrusive)
}

// List provides the implementation of intrusive linked lists
type List struct {
	prev Intrusive
	next Intrusive
}

func (l *List) Next() Intrusive {
	return l.next
}

func (l *List) Prev() Intrusive {
	return l.prev
}

func (l *List) AddNext(i Intrusive) {
	l.next = i
}

func (l *List) AddPrev(i Intrusive) {
	l.prev = i
}

func (l *List) Reset() {
	l.prev = nil
	l.next = nil
}

func (l *List) Empty() bool {
	return l.prev == nil
}

// Front returns the first element of list l or nil.
func (l *List) Front() Intrusive {
	return l.prev
}

// Back returns the last element of list l or nil.
func (l *List) Back() Intrusive {
	return l.next
}

// PushFront inserts the element e at the front of list l.
func (l *List) PushFront(e Intrusive) {
	e.AddPrev(nil)
	e.AddNext(l.prev)

	if l.prev != nil {
		l.prev.AddPrev(e)
	} else {
		l.next = e
	}
	l.prev = e
}

// PushBack inserts the element e at the back of list l.
func (l *List) PushBack(e Intrusive) {
	e.AddNext(nil)
	e.AddPrev(l.next)

	if l.next != nil {
		l.next.AddNext(e)
	} else {
		l.prev = e
	}
	l.next = e
}

// InsertAfter inserts e after b.
func (l *List) InsertAfter(e, b Intrusive) {
	a := b.Next()
	e.AddNext(a)
	e.AddPrev(b)
	b.AddNext(e)

	if a != nil {
		a.AddPrev(e)
	} else {
		l.next = e
	}
}

// InsertBefore inserts e before a.
func (l *List) InsertBefore(e, a Intrusive) {
	b := a.Prev()
	e.AddNext(a)
	e.AddPrev(b)
	a.AddPrev(e)

	if b != nil {
		b.AddNext(e)
	} else {
		l.prev = e
	}
}

// Remove removes e from l.
func (l *List) Remove(e Intrusive) {
	prev := e.Prev()
	next := e.Next()

	if prev != nil {
		prev.AddNext(next)
	} else {
		l.prev = next
	}

	if next != nil {
		next.AddPrev(prev)
	} else {
		l.next = prev
	}
}

//----------------以下为使用demo-------------
//我们这里用List表示实现了Intrusive接口，也实现了链表的基本功能，所以任何实现了Intrusive接口的对象都是可以作为链表的节点，利用这个介入式链表就很简单了，只要在现有的struct嵌入List这个结构体即可，在举个例子：
func TestIntrusiveList(t *testing.T) {
	type E struct {
		List
		data int
	}
	// Create a new list and put some numbers in it.
	l := List{}
	e4 := &E{data: 4}
	e1 := &E{data: 1}
	l.PushBack(e4)
	l.PushFront(e1)
	l.InsertBefore(&E{data: 3}, e4)
	l.InsertAfter(&E{data: 2}, e1)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("e: %+v\n", e)
		fmt.Printf("data: %d\n", e.(*E).data)
	}
}

//=== RUN   TestIntrusiveList
//e: &{List:{prev:<nil> next:0xc04205c450} data:1}
//data: 1
//e: &{List:{prev:0xc04205c3f0 next:0xc04205c420} data:2}
//data: 2
//e: &{List:{prev:0xc04205c450 next:0xc04205c3c0} data:3}
//data: 3
//e: &{List:{prev:0xc04205c420 next:<nil>} data:4}
//data: 4
//--- PASS: TestIntrusiveList (0.00s)
//PASS
