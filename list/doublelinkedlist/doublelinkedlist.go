package doublelinkedlist

import (
	"errors"
)

var (
	// ErrIndex is returned when the index is out of the list
	ErrIndex = errors.New("index is out of the list")
	// ErrIndexOf is returned when the index of value can not found
	ErrIndexOf = errors.New("index of value can not found")
)

// List
type List struct {
	first *element
	last  *element
	size  int
}

type element struct {
	value interface{}
	prev  *element
	next  *element
}

// New
func New(value ...interface{}) *List {
	list := &List{}

	if len(value) > 0 {
		list.Append(value...)
	}

	return list
}

// List Interface

// Append values (one or more than one) to list.
func (list *List) Append(value ...interface{}) {
	if len(value) == 0 {
		return
	}

	// if size is equal to 0, it is a new double linked list
	if list.size == 0 {
		for i, v := range value {
			newElement := &element{value: v}
			// first element
			if i == 0 {
				list.first = newElement
				list.last = newElement
			} else {
				list.last.next = newElement
				newElement.prev = list.last
				list.last = newElement
			}
			list.size++
		}
	} else {
		for _, v := range value {
			newElement := &element{value: v}
			list.last.next = newElement
			newElement.prev = list.last
			list.last = newElement
			list.size++
		}
	}
}

// PreAppend
func (list *List) PreAppend(values ...interface{}) {
	if len(values) == 0 {
		return
	}

	if list.size == 0 {
		list.Append(values...)
	} else {
		for i := len(values) - 1; i >= 0; i-- {
			newElement := &element{value: values[i]}
			newElement.next = list.first
			list.first.prev = newElement
			list.first = newElement
			list.size++
		}
	}
}

// Check if the index is within the length of the list
func (list *List) indexInRange(index int) bool {
	if index >= 0 && index < list.size {
		return true
	}
	return false
}

// Get
func (list *List) Get(index int) (interface{}, error) {
	if !list.indexInRange(index) {
		return nil, ErrIndex
	}

	foundElement := list.first
	for i := 0; i != index; i++ {
		foundElement = foundElement.next
	}

	return foundElement.value, nil
}

// Remove
func (list *List) Remove(index int) error {
	if !list.indexInRange(index) {
		return ErrIndex
	}

	foundElement := list.first

	if index == 0 {
		list.first = list.first.next
		list.first.next.prev = nil
	} else {
		preFoundElement := new(element)
		for i := 0; i != index; i++ {
			preFoundElement = foundElement
			foundElement = foundElement.next
		}
		preFoundElement.next = foundElement.next
		foundElement.next.prev = preFoundElement
	}

	foundElement.value = nil
	foundElement.next = nil
	foundElement.prev = nil
	list.size--

	return nil
}

// Contains
func (list *List) Contains(values ...interface{}) bool {
	if len(values) == 0 {
		return true
	}
	if list.size == 0 {
		return false
	}

	for _, value := range values {
		found := false
		for element := list.first; element != nil; element = element.next {
			if element.value == value {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Swap
func (list *List) Swap(i, j int) error {
	if !list.indexInRange(i) || !list.indexInRange(j) {
		return ErrIndex
	}
	if i == j {
		return nil
	}

	foundElementI := list.first
	for index := 0; index != i; index++ {
		foundElementI = foundElementI.next
	}

	foundElementJ := list.first
	for index := 0; index != j; index++ {
		foundElementJ = foundElementJ.next
	}

	foundElementI.value, foundElementJ.value = foundElementJ.value, foundElementI.value

	return nil
}

// Insert
func (list *List) Insert(index int, values ...interface{}) error {
	if len(values) == 0 {
		return nil
	}
	if !list.indexInRange(index) {
		return ErrIndex
	}
	if index == list.size-1 {
		list.Append(values...)
		return nil
	}

	foundElement := list.first
	for i := 0; i != index; i++ {
		foundElement = foundElement.next
	}
	foundElementNext := foundElement.next

	// insert
	for _, v := range values {
		element := &element{value: v}
		foundElement.next = element
		element.prev = foundElement
		foundElement = element
	}
	foundElement.next = foundElementNext
	foundElementNext.prev = foundElement
	list.size += len(values)

	return nil
}

// Set
func (list *List) Set(index int, value interface{}) error {
	if !list.indexInRange(index) {
		return ErrIndex
	}

	foundElement := list.first
	for i := 0; i != index; i++ {
		foundElement = foundElement.next
	}
	foundElement.value = value

	return nil
}

// IndexOf
func (list *List) IndexOf(value interface{}) (int, error) {
	for i, v := range list.Values() {
		if v == value {
			return i, nil
		}
	}
	return -1, ErrIndexOf
}

// Reverse
func (list *List) Reverse() {
	if list.size == 0 || list.size == 1 {
		return
	}

	preElement := new(element)
	preElement.value = nil
	preElement.next = nil
	preElement.prev = nil
	currElement := list.first
	nextElement := list.first.next

	// reset the last pointer
	list.last = currElement

	// reverse
	for currElement != nil {
		currElement.next = preElement
		currElement.prev = nextElement
		preElement = currElement
		currElement = nextElement
		if nextElement != nil {
			nextElement = nextElement.next
		}
	}

	// reset the first pointer
	list.first = preElement
}

// Container Interface

// Empty
func (list *List) Empty() bool {
	return list.size == 0
}

// Size
func (list *List) Size() int {
	size := list.size
	return size
}

// Clear
func (list *List) Clear() {
	list.size = 0
	list.first = nil
	list.last = nil
}

// Values
func (list *List) Values() []interface{} {
	values := make([]interface{}, 0)

	iterator := list.Iterator()
	iterator.Begin()
	for iterator.Next() {
		values = append(values, iterator.Value())
	}

	return values
}
