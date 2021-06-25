package plugin

import "time"

type Metadata struct {
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	ThumbnailUrl string    `json:"thumbnailUrl"`
	Author       string    `json:"author"`
	CreatedAt    time.Time `json:"createdAt"`
}

type MetadataList []*Metadata

func (ml *MetadataList) FindByAuthorAndName(name, author string) *Metadata {
	for _, m := range *ml {
		if m.Author == author && m.Name == name {
			return m
		}
	}
	return nil
}
