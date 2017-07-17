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

func colorOf(e *Entry) bool{
	if e == nil {
		return BLACK
	}
	return e.color
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
