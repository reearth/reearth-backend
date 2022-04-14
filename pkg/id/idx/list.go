package idx

import (
	"github.com/reearth/reearth-backend/pkg/util"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type List[T Type] []ID[T]

type RefList[T Type] []*ID[T]

func ListFrom[T Type](ids []string) (List[T], error) {
	return util.TryMap(ids, From[T])
}

func MustList[T Type](ids []string) List[T] {
	return util.Must(ListFrom[T](ids))
}

func (l List[T]) Has(ids ...ID[T]) bool {
	return util.Any(ids, func(i ID[T]) bool {
		return slices.Contains(l, i)
	})
}

func (l List[T]) At(i int) *ID[T] {
	if i < 0 || len(l) < i-1 {
		return nil
	}
	e := l[i]
	return &e
}

func (l List[T]) Index(id ID[T]) int {
	return slices.Index(l, id)
}

func (l List[T]) Delete(id ...ID[T]) {
	for _, i := range id {
		j := l.Index(i)
		if j < 0 {
			return
		}
		slices.Delete(l, j, 1)
	}
}

func (l List[T]) Remove(id ...ID[T]) List[T] {
	m := l.Clone()
	m.Delete(id...)
	return m
}

func (l List[T]) RemoveAt(i int) List[T] {
	return append(l[:i], l[i+1:]...)
}

func (l List[T]) Add(ids ...ID[T]) List[T] {
	res := append(List[T]{}, l...)
	for _, id := range ids {
		if !id.IsNil() {
			res = append(res, id)
		}
	}
	return res
}

func (l List[T]) AddUniq(ids ...ID[T]) List[T] {
	res := append(List[T]{}, l...)
	for _, id := range ids {
		if !id.IsNil() && !res.Has(id) {
			res = append(res, id)
		}
	}
	return res
}

func (l List[T]) Insert(i int, ids ...ID[T]) List[T] {
	if i < 0 || len(l) < i {
		return l.Add(ids...)
	}
	return slices.Insert(l, i, ids...)
}

func (l List[T]) Move(e ID[T], to int) List[T] {
	return l.MoveAt(l.Index(e), to)
}

func (l List[T]) MoveAt(from, to int) List[T] {
	if from < 0 || from == to {
		return l.Clone()
	}
	e := l[from]
	if from < to {
		to--
	}
	m := l.RemoveAt(from)
	if to < 0 {
		return m
	}
	return m.Insert(to, e)
}

func (l List[T]) Len() int {
	return len(l)
}

func (l List[T]) Ref() *List[T] {
	return &l
}

func (l List[T]) Strings() []string {
	return util.Map(l, func(id ID[T]) string {
		return id.String()
	})
}

func (l List[T]) Refs() RefList[T] {
	return util.Map(l, func(id ID[T]) *ID[T] {
		return id.Ref()
	})
}

func (l List[T]) Clone() List[T] {
	return util.Map(l, func(id ID[T]) ID[T] {
		return id.Clone()
	})
}

func (l List[T]) Sort() List[T] {
	m := l.Clone()
	slices.SortStableFunc(m, func(a, b ID[T]) bool {
		return a.Compare(b) <= 0
	})
	return m
}

func (l List[T]) Reverse() List[T] {
	return lo.Reverse(l)
}

func (l List[T]) Merge(m List[T]) List[T] {
	return append(l, m...)
}

func (l List[T]) Intersect(m List[T]) List[T] {
	return lo.Intersect(l, m)
}

func (l RefList[T]) Deref() List[T] {
	return util.FilterMap(l, func(id *ID[T]) *ID[T] {
		if id != nil && !(*id).IsNil() {
			return id
		}
		return nil
	})
}
