package a

type Set map[any]Void

func NewSet(items ...any) Set {
	var res Set = make(map[any]Void)
	for _, item := range items {
		res.Add(item)
	}
	return res
}

func (s *Set) Add(item any) {
	(*s)[item] = empty
}
func (s *Set) Contains(item any) bool {
	_, ok := (*s)[item]
	return ok
}

func (s *Set) Del(item any) {
	delete(*s, item)
}

func (s *Set) Each(acceptor Consumer) {
	for k := range *s {
		acceptor(k)
	}
}

func (s *Set) ToArr() []any {
	var res []any
	for k := range *s {
		res = append(res, k)
	}
	return res
}
