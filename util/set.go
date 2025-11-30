package util

type Settable interface {
	AsSet() UniqueSet
}

type Stringable interface {
	String() string
}

type UniqueSet map[string]struct{}

func NewSet(v ...Stringable) UniqueSet {
	set := make(UniqueSet)
	for _, val := range v {
		set.Insert(val)
	}
	return set
}

func (s UniqueSet) Insert(k Stringable) UniqueSet {
	s[k.String()] = struct{}{}
	return s
}

func (s UniqueSet) Include(k Stringable) bool {
	_, ok := s[k.String()]
	return ok
}

func Set(s Settable) UniqueSet {
	return s.AsSet()
}
