package treemap

// Red-black mechanics
const (
	RED 	= true
	BLACK	= false
)

// Entry is an entry of a TreeMap.
type Entry struct {
	left	*Entry
	right	*Entry
	parent	*Entry
	color 	bool
	// The key of the entry.
	key 	interface{}
	// The value of the entry.
	value 	interface{}
}

// GetKey returns the key of the entry.
func (entry *Entry) GetKey() interface{} {
	return entry.key
}

// GetValue returns the value of the entry.
func (entry *Entry) GetValue() interface{} {
	return entry.value
}

func parentOf(e *Entry) *Entry {
	if e == nil {
		return nil
	}
	
	return e.parent
}

func setColor(e *Entry, color bool) {
	if e != nil {
		e.color = color
	}
}

func leftOf(e *Entry) *Entry {
	if e == nil {
		return nil
	}
	return e.left
}

func rightOf(e *Entry) *Entry {
	if e == nil {
		return nil
	}
	return e.right
}

func colorOf(e *Entry) bool {
	if e == nil {
		return BLACK
	}
	return e.color
}

// successor returns the successor of the specified Entry, or null if no such.
func successor(t *Entry) *Entry {
	if t == nil {
		return nil
	} else if t.right != nil {
		p := t.right
		for p.left != nil {
			p = p.left
		}
		return p
	} else {
		p := t.parent
		ch := t
		for p != nil && ch == p.right {
			ch = p
			p = p.parent
		}
		return p
	}
}

// TreeMap represents a map that is implemented by a red and black tree.
type TreeMap struct {
	root *Entry
	size int
}

// New returns an initialized TreeMap.
func New() *TreeMap {
	return new(TreeMap).Init()
}

// Init initializes or clears TreeMap tm.
func (tm *TreeMap) Init() *TreeMap {
	tm.root = nil
	tm.size = 0
	return tm
}

// Put insert a new entry with key and value in the TreeMap and returns previous value.
func (tm *TreeMap) Put(key, value interface{}) (pre interface{}, err error) {
	t := tm.root

	// If TreeMap is empty
	if t == nil {
		tm.root = &Entry{
			color: 	BLACK,
			key: 	key,
			value: 	value,
		}
		tm.size = 1
		return nil, nil
	}

	var parent *Entry
	var cmp int
	for {
		parent = t
		cmp, err = compare(key, t.key)
		if err != nil {
			return
		}
		switch {
			case cmp < 0:
				t = t.left
			case cmp > 0:
				t = t.right
			case cmp == 0:
				pre = t.value
				t.value = value
				return
		}
		if t == nil {
			break
		}
	}
	e := &Entry {
		parent: parent,
		key: key,
		value: value,
	}
	if cmp < 0 {
		parent.left = e
	} else {
		parent.right = e
	}
	tm.fixAfterInsertion(e)
	tm.size++

	return
}

// Get returns the value of the value to which the specified key is mapped or nil.
func (tm *TreeMap) Get(key interface{}) interface{} {
	p, err := tm.getEntry(key)
	if err != nil || p == nil{
		return nil
	}
	return p.GetValue()
}

func (tm *TreeMap) getEntry(key interface{}) (*Entry, error) {
	p := tm.root
	for p != nil {
		cmp, err := compare(key, p.GetKey())
		if err != nil {
			return nil, err
		}
		switch {
			case cmp < 0: 
				p = p.left
			case cmp > 0:
				p = p.right
			default:
				return p, nil
		}
	}
	return nil, nil
}

// Remove the mapping for this key from this TreeMap if present.
func (tm *TreeMap) Remove(key interface{}) (value interface{}, err error) {
	p, err := tm.getEntry(key)
	if err != nil {
		return
	}

	if p == nil {
		value, err = nil, nil
		return
	}

	value = p.value
	err = tm.deleteEntry(p)
	if err != nil {
		return
	}
	return
}

func (tm *TreeMap) deleteEntry(p *Entry) (err error) {
	tm.size--

	if p.left != nil && p.right != nil {
		s := successor(p)
		p.key = s.key
		p.value = s.value
		p = s
	}

	var replacement *Entry
	if p.left != nil {
		replacement = p.left
	} else {
		replacement = p.right
	}

	if replacement != nil {
		replacement.parent = p.parent
		if p.parent == nil {
			tm.root = replacement
		} else if p == p.parent.left {
			p.parent.left = replacement
		} else {
			p.parent.right = replacement
		}

		p.left, p.right, p.parent = nil, nil, nil

		if p.color == BLACK {
			tm.fixAfterDeletion(replacement)
		}
	} else if p.parent == nil {
		tm.root = nil
	} else {
		if p.color == BLACK {
			tm.fixAfterDeletion(p)
		}
		if p.parent != nil {
			if p == p.parent.left {
				p.parent.left = nil;
			} else if p == p.parent.right {
				p.parent.right = nil
			}
			p.parent = nil
		}
	}

	return
}

// Size returns the number of key-value mappings in this map.
func (tm *TreeMap) Size() int {
	return tm.size
}

func (tm *TreeMap) rotateLeft(p *Entry) {
	if p != nil {
		r := p.right
		p.right = r.left
		if r.left != nil {
			r.left.parent = p
		}
		r.parent = p.parent
		if p.parent == nil {
			tm.root = r
		} else if p.parent.left == p {
			p.parent.left = r
		} else {
			p.parent.right = r
		}
		r.left = p
		p.parent = r
	}
}

func (tm *TreeMap) rotateRight(p *Entry) {
	if p != nil {
		l := p.left
		p.left = l.left
		p.left = l.right
		if l.right != nil {
			l.right.parent = p
		}
		l.parent = p.parent
		if p.parent == nil {
			tm.root = l
		} else if p.parent.right == p {
			p.parent.right = l
		} else {
			p.parent.left = l
		}
		l.right = p
		p.parent = l
	}
}

func (tm *TreeMap) fixAfterInsertion(x *Entry) {
	x.color = RED
	for x != nil && x != tm.root && x.parent.color == RED {
		if parentOf(x) == leftOf(parentOf(parentOf(x))) {
			y := rightOf(parentOf(parentOf(x)))
			if colorOf(y) == RED {
				setColor(parentOf(x), BLACK)
				setColor(y, BLACK)
				setColor(parentOf(parentOf(x)), RED)
				x = parentOf(parentOf(x))
			} else {
				if x == rightOf(parentOf(x)) {
					x = parentOf(x)
					tm.rotateLeft(x)
				}
				setColor(parentOf(x), BLACK)
				setColor(parentOf(parentOf(x)), RED)
				tm.rotateRight(parentOf(parentOf(x)))
			}
		} else {
			y := leftOf(parentOf(parentOf(x)))
			if colorOf(y) == RED {
				setColor(parentOf(x), BLACK)
				setColor(y, BLACK)
				setColor(parentOf(parentOf(x)), RED)
				x = parentOf(parentOf(x))
			} else {
				if x == leftOf(parentOf(x)) {
					x = parentOf(x)
					tm.rotateRight(x)
				}
				setColor(parentOf(x), BLACK)
				setColor(parentOf(parentOf(x)), RED)
				tm.rotateLeft(parentOf(parentOf(x)))
			}
		}
	}
	tm.root.color = BLACK
}

func (tm *TreeMap) fixAfterDeletion(x *Entry) {
	for x != tm.root && colorOf(x) == BLACK {
		if x == leftOf(parentOf(x)) {
			sib := rightOf(parentOf(x))

			if colorOf(sib) == RED {
				setColor(sib, BLACK)
				setColor(parentOf(x), RED)
				tm.rotateLeft(parentOf(x))
			}

			if colorOf(leftOf(sib)) == BLACK && colorOf(rightOf(sib)) == BLACK {
					setColor(sib, RED)
					x = parentOf(x)
			} else {
				if colorOf(rightOf(sib)) == BLACK {
					setColor(leftOf(sib), BLACK)
					setColor(sib, RED)
					tm.rotateRight(sib)
					sib = rightOf(parentOf(x))
				}
				setColor(sib, colorOf(parentOf(x)))
				setColor(parentOf(x), BLACK)
				setColor(rightOf(sib), BLACK)
				tm.rotateLeft(parentOf(x))
				x = tm.root
			}
		} else {
			sib := leftOf(parentOf(x))
			if colorOf(sib) == RED {
				setColor(sib, BLACK)
				setColor(parentOf(x), RED)
				tm.rotateRight(parentOf(x))
				sib = leftOf(parentOf(x))
			}

			if colorOf(rightOf(sib)) == BLACK && colorOf(leftOf(sib)) == BLACK {
				setColor(sib, RED);
                x = parentOf(x);
			} else {
				if colorOf(leftOf(sib)) == BLACK {
                    setColor(rightOf(sib), BLACK);
                    setColor(sib, RED);
                    tm.rotateLeft(sib);
                    sib = leftOf(parentOf(x));
                }
                setColor(sib, colorOf(parentOf(x)));
                setColor(parentOf(x), BLACK);
                setColor(leftOf(sib), BLACK);
                tm.rotateRight(parentOf(x));
                x = tm.root;
			}
		}
	}
	setColor(x, BLACK)
}

func (tm *TreeMap) firstEntry() *Entry {
	p := tm.root
	if p != nil {
		for p.left != nil {
			p = p.left
		}
	}
	return p
}

// EntryIterator returns a EntryIterator of the TreeMap.
func (tm *TreeMap) EntryIterator() *EntryIterator {
	return &EntryIterator {
		next: tm.firstEntry(),
	}
}
