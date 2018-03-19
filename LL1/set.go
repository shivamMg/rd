package main

// Set is a simple set implementation
type Set struct {
	m map[string]struct{}
}

func NewSet(elems []string) *Set {
	s := Set{m: make(map[string]struct{})}
	for _, ele := range elems {
		s.m[ele] = struct{}{}
	}
	return &s
}

func (s *Set) List() (l []string) {
	for e := range s.m {
		l = append(l, e)
	}

	return l
}

func (s *Set) Add(ele string) {
	(*s).m[ele] = struct{}{}
}

func (s *Set) Delete(ele string) {
	delete((*s).m, ele)
}

func (s *Set) Merge(s2 Set) {
	for ele := range s2.m {
		(*s).m[ele] = struct{}{}
	}
}

func (s Set) Has(ele string) bool {
	for e := range s.m {
		if e == ele {
			return true
		}
	}
	return false
}
