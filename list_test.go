package ringlist_test

import (
	"testing"

	"github.com/mgnsk/ringlist"
	. "github.com/onsi/gomega"
)

func TestPushFront(t *testing.T) {
	var l ringlist.List[int]

	g := NewWithT(t)

	l.PushFront(ringlist.NewElement(0))
	g.Expect(l.Len()).To(Equal(1))

	l.PushFront(ringlist.NewElement(1))
	g.Expect(l.Len()).To(Equal(2))

	expectValidRing(g, &l)
}

func TestPushBack(t *testing.T) {
	var l ringlist.List[int]

	g := NewWithT(t)

	l.PushFront(ringlist.NewElement(0))
	g.Expect(l.Len()).To(Equal(1))

	l.PushFront(ringlist.NewElement(1))
	g.Expect(l.Len()).To(Equal(2))

	expectValidRing(g, &l)
}

func TestMoveToFront(t *testing.T) {
	t.Run("moving the back element", func(t *testing.T) {
		var l ringlist.List[string]

		g := NewWithT(t)

		l.PushBack(ringlist.NewElement("one"))
		l.PushBack(ringlist.NewElement("two"))
		l.MoveToFront(l.Back())

		expectValidRing(g, &l)
		g.Expect(l.Front().Value).To(Equal("two"))
		g.Expect(l.Back().Value).To(Equal("one"))
	})

	t.Run("moving the middle element", func(t *testing.T) {
		var l ringlist.List[string]

		g := NewWithT(t)

		l.PushBack(ringlist.NewElement("one"))
		l.PushBack(ringlist.NewElement("two"))
		l.PushBack(ringlist.NewElement("three"))
		l.MoveToFront(l.Front().Next())

		expectValidRing(g, &l)
		g.Expect(l.Front().Value).To(Equal("two"))
		g.Expect(l.Back().Value).To(Equal("three"))
	})
}

func TestMoveToBack(t *testing.T) {
	t.Run("moving the front element", func(t *testing.T) {
		var l ringlist.List[string]

		g := NewWithT(t)

		l.PushBack(ringlist.NewElement("one"))
		l.PushBack(ringlist.NewElement("two"))
		l.MoveToBack(l.Front())

		expectValidRing(g, &l)
		g.Expect(l.Front().Value).To(Equal("two"))
		g.Expect(l.Back().Value).To(Equal("one"))
	})

	t.Run("moving the middle element", func(t *testing.T) {
		var l ringlist.List[string]

		g := NewWithT(t)

		l.PushBack(ringlist.NewElement("one"))
		l.PushBack(ringlist.NewElement("two"))
		l.PushBack(ringlist.NewElement("three"))
		l.MoveToBack(l.Front().Next())

		expectValidRing(g, &l)
		g.Expect(l.Front().Value).To(Equal("one"))
		g.Expect(l.Back().Value).To(Equal("two"))
	})
}

func TestMoveBefore(t *testing.T) {
	t.Run("before itself", func(t *testing.T) {
		var l ringlist.List[string]

		g := NewWithT(t)

		l.PushBack(ringlist.NewElement("one"))
		l.PushBack(ringlist.NewElement("two"))
		l.PushBack(ringlist.NewElement("three"))
		expectValidRing(g, &l)

		one := l.Front()
		two := l.Front().Next()
		three := l.Front().Next().Next()

		g.Expect(one.Value).To(Equal("one"))
		g.Expect(two.Value).To(Equal("two"))
		g.Expect(three.Value).To(Equal("three"))

		l.MoveToFront(one)
		l.MoveToFront(two)
		l.MoveToFront(three)
		g.Expect(l.Len()).To(Equal(3))

		expectHasExactElements(g, &l, "three", "two", "one")
	})
}

func TestMoveAfter(t *testing.T) {
	t.Run("after itself", func(t *testing.T) {
		var l ringlist.List[string]

		g := NewWithT(t)

		l.PushBack(ringlist.NewElement("one"))
		l.PushBack(ringlist.NewElement("two"))
		l.PushBack(ringlist.NewElement("three"))
		expectValidRing(g, &l)

		one := l.Front()
		two := l.Front().Next()
		three := l.Front().Next().Next()

		g.Expect(one.Value).To(Equal("one"))
		g.Expect(two.Value).To(Equal("two"))
		g.Expect(three.Value).To(Equal("three"))

		l.MoveToBack(three)
		l.MoveToBack(two)
		l.MoveToBack(one)
		g.Expect(l.Len()).To(Equal(3))

		expectHasExactElements(g, &l, "three", "two", "one")
	})
}

func TestMoveForward(t *testing.T) {
	t.Run("overflow", func(t *testing.T) {
		var l ringlist.List[string]

		g := NewWithT(t)

		l.PushBack(ringlist.NewElement("one"))
		l.PushBack(ringlist.NewElement("two"))
		l.Move(l.Front(), 3)

		expectValidRing(g, &l)
		g.Expect(l.Front().Value).To(Equal("two"))
		g.Expect(l.Back().Value).To(Equal("one"))
	})

	t.Run("not overflow", func(t *testing.T) {
		var l ringlist.List[string]

		g := NewWithT(t)

		l.PushBack(ringlist.NewElement("one"))
		l.PushBack(ringlist.NewElement("two"))
		l.PushBack(ringlist.NewElement("three"))
		l.Move(l.Front(), 1)

		expectValidRing(g, &l)
		g.Expect(l.Front().Value).To(Equal("two"))
		g.Expect(l.Front().Next().Value).To(Equal("one"))
		g.Expect(l.Back().Value).To(Equal("three"))
	})
}

func TestMoveBackwards(t *testing.T) {
	t.Run("overflow", func(t *testing.T) {
		var l ringlist.List[string]

		g := NewWithT(t)

		l.PushBack(ringlist.NewElement("one"))
		l.PushBack(ringlist.NewElement("two"))
		l.Move(l.Back(), -3)

		expectValidRing(g, &l)
		g.Expect(l.Front().Value).To(Equal("two"))
		g.Expect(l.Back().Value).To(Equal("one"))
	})

	t.Run("not overflow", func(t *testing.T) {
		var l ringlist.List[string]

		g := NewWithT(t)

		l.PushBack(ringlist.NewElement("one"))
		l.PushBack(ringlist.NewElement("two"))
		l.PushBack(ringlist.NewElement("three"))
		l.Move(l.Back(), -1)

		expectValidRing(g, &l)
		g.Expect(l.Front().Value).To(Equal("one"))
		g.Expect(l.Front().Next().Value).To(Equal("three"))
		g.Expect(l.Back().Value).To(Equal("two"))
	})
}

func TestDo(t *testing.T) {
	var l ringlist.List[string]

	g := NewWithT(t)

	l.PushBack(ringlist.NewElement("one"))
	l.PushBack(ringlist.NewElement("two"))
	l.PushBack(ringlist.NewElement("three"))

	g.Expect(l.Len()).To(Equal(3))
	expectValidRing(g, &l)

	var elems []string
	l.Do(func(e *ringlist.Element[string]) bool {
		elems = append(elems, e.Value)
		return true
	})

	g.Expect(elems).To(Equal([]string{"one", "two", "three"}))
}

func expectHasExactElements[T any](g *WithT, l *ringlist.List[T], elements ...any) {
	var elems []T

	l.Do(func(e *ringlist.Element[T]) bool {
		elems = append(elems, e.Value)

		return true
	})

	g.Expect(elems).To(HaveExactElements(elements...))
}

func expectValidRing[T any](g *WithT, l *ringlist.List[T]) {
	g.Expect(l.Len()).To(BeNumerically(">", 0))
	g.Expect(l.Front()).To(Equal(l.Back().Next()))
	g.Expect(l.Back()).To(Equal(l.Front().Prev()))

	{
		expectedFront := l.Front()

		front := l.Front()

		for i := 0; i < l.Len(); i++ {
			front = front.Next()
		}

		g.Expect(front).To(Equal(expectedFront))
	}

	{
		expectedBack := l.Back()

		back := l.Back()

		for i := 0; i < l.Len(); i++ {
			back = back.Prev()
		}

		g.Expect(back).To(Equal(expectedBack))
	}
}
