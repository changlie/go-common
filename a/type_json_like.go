package a

type Obj Map

func ObjNew() Obj {
	var raw = M{}
	return raw
}
func ObjNew2(raw Map) Obj {
	return raw
}
func (o *Obj) S(k string, v any) *Obj {
	(*o)[k] = v
	return o
}
func (o *Obj) GetArr(k string) Arr {
	if v, ok := (*o)[k]; ok {
		if res, ok := v.(List); ok {
			return res
		}
		return L{}
	}
	return L{}
}
func (o *Obj) GetObj(k string) Obj {
	if v, ok := (*o)[k]; ok {
		if res, ok := v.(Map); ok {
			return res
		}
		return M{}
	}
	return M{}
}
func (o *Obj) Str(k string) string {
	if v, ok := (*o)[k]; ok {
		if res, ok := v.(string); ok {
			return res
		}
		return ""
	}
	return ""
}
func (o *Obj) Int(k string) int {
	if v, ok := (*o)[k]; ok {
		if res, ok := v.(int); ok {
			return res
		}
		return 0
	}
	return 0
}
func (o *Obj) Json() string {
	return Json(o)
}

type Arr List

func ArrNew() Arr {
	var arr = L{}
	return arr
}
func ArrNew2(arr List) Arr {
	return arr
}
func (a *Arr) A(items ...any) *Arr {
	for _, item := range items {
		*a = append(*a, item)
	}
	return a
}
func (a *Arr) Set(i int, item any) *Arr {
	(*a)[i] = item
	return a
}
func (a *Arr) GetArr(i int) Arr {
	if i >= 0 && len(*a) > i {
		if res, ok := (*a)[i].(List); ok {
			return res
		}
		return L{}
	}
	return L{}
}
func (a *Arr) GetObj(i int) Obj {
	if i >= 0 && len(*a) > i {
		if res, ok := (*a)[i].(Map); ok {
			return res
		}
		return M{}
	}
	return M{}
}
func (a *Arr) Str(i int) string {
	if i >= 0 && len(*a) > i {
		if res, ok := (*a)[i].(string); ok {
			return res
		}
		return ""
	}
	return ""
}
func (a *Arr) Int(i int) int {
	if i >= 0 && len(*a) > i {
		if res, ok := (*a)[i].(int); ok {
			return res
		}
		return 0
	}
	return 0
}
func (a *Arr) Json() string {
	return Json(a)
}
