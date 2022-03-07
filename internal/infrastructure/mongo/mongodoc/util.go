package mongodoc

import "go.mongodb.org/mongo-driver/bson"

func convertDToM(i interface{}) interface{} {
	if i == nil {
		return nil
	}
	switch i2 := i.(type) {
	case bson.D:
		return i2.Map()
	case bson.A:
		a := make([]interface{}, 0, len(i2))
		for _, e := range i2 {
			a = append(a, convertDToM(e))
		}
		return a
	}
	return i
}
