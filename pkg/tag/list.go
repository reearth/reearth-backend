package tag

import "github.com/reearth/reearth-backend/pkg/id"

type List struct {
	tags []id.TagID
}

func NewList() *List {
	return &List{tags: []id.TagID{}}
}

func NewListFromTags(tags []id.TagID) *List {
	return &List{tags: tags}
}

func (tl *List) Tags() []id.TagID {
	return tl.tags
}

func (tl *List) Has(tid id.TagID) bool {
	if tl == nil || tl.tags == nil {
		return false
	}
	for _, tag := range tl.tags {
		if tag == tid {
			return true
		}
	}
	return false
}

func (tl *List) Add(tags ...id.TagID) {
	if tl == nil || tl.tags == nil {
		return
	}
	tl.tags = append(tl.tags, tags...)
}

func (tl *List) Remove(tags ...id.TagID) {
	if tl == nil || tl.tags == nil {
		return
	}
	res := make([]id.TagID, 0)
	for _, t := range tl.tags {
		found := false
		for _, tid := range tags {
			if t == tid {
				found = true
				break
			}
		}
		if !found {
			res = append(res, t)
		}
	}
	tl.tags = res
}
