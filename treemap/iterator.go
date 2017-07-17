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

// HasNext returns true if has more entry.
func (iter *EntryIterator) HasNext() bool {
	return iter.next != nil
}
