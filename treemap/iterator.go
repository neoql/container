package treemap

// EntryIterator is the iterator of the TreeMap.
type EntryIterator struct {
	next *Entry
}

// Next returns next Entry.
func (iter *EntryIterator) Next() *Entry {
	e := iter.next
	iter.next = successor(e)
	return e
}

func (iter *EntryIterator) HasNext() bool {
	return iter.next != nil
}
