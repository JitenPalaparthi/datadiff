package compares

import (
	"encoding/json"
	"errors"
	"reflect"

	"gopkg.in/yaml.v2"
)

type Comparer interface {
	AreEqual(...[]byte) (bool, int, error)
	IsEqual([]byte, []byte) (bool, error)
	Compare([]byte, []byte) (bool, []string, []string, []any, error)
}

type JsonCompares struct{}

var (
	ErrOnlyOneItem = errors.New("there is only one item to compare")
	ErrNoItem      = errors.New("there is no item to compare")
)

// AreEqual checks all the items as arguments are equal or not.
// It only checks with the previous item. Does not check against each item as a bubble
// As soon as two items are different, it immediately return false,index,nil
// If all items are same then it returns true, 0, nil
// If there is no item or one item then it returns respective errors as false,0, ErrNoItem or false,0,ErrOnlyOneItem
// It is ideal to pass only two items to compare
func (c *JsonCompares) AreEqual(items ...[]byte) (bool, int, error) {
	if len(items) == 0 {
		return false, 0, ErrNoItem
	}
	if len(items) == 1 {
		return false, 0, ErrOnlyOneItem
	}
	equal := false
	data1 := items[0]
	for i := 1; i < len(items); i++ {
		data2 := items[i]
		if reflect.DeepEqual(data1, data2) {
			equal = true
			data1 = items[i]
		} else {
			return equal, i, nil
		}
	}
	return equal, 0, nil
}

// IsEqual is to check whether x and y are equal or not
// IsEqual returns bool, error
// if argument x or argument y is nil it returns false, "x is nil" or false, "y is nil"
func (c *JsonCompares) IsEqual(x []byte, y []byte) (bool, error) {
	if x == nil {
		return false, errors.New("x is nil item")
	}
	if y == nil {
		return false, errors.New("y is nil item")
	}
	if reflect.DeepEqual(x, y) {
		return true, nil
	}
	return false, nil
}

// Compare compares y against x provided by x and y are byte arrays
// the first return type is true of both of them are equal
// the second return type string slice returns new keys in y against x
// the third return type string slice returns deleted keys in y against x
// the fourth return type string slice returns common keys with different values in x and y
// the fifth return type error returns error if there is any error in data or in unmarshal
func (c *JsonCompares) Compare(x []byte, y []byte) (equal bool, newKeys []string, delKeys []string, diffValueKeys []any, err error) {
	if x == nil {
		return false, nil, nil, nil, errors.New("x is nil item")
	}
	if y == nil {
		return false, nil, nil, nil, errors.New("y is nil item")
	}

	if !json.Valid(x) {
		return false, nil, nil, nil, errors.New("x is invalid json")
	}

	if !json.Valid(y) {
		return false, nil, nil, nil, errors.New("x is invalid json")
	}

	xmap, ymap := make(map[string]any), make(map[string]any)

	newKeys = make([]string, 0)
	delKeys = make([]string, 0)
	diffValueKeys = make([]any, 0)
	err = json.Unmarshal(x, &xmap)
	if err != nil {
		return false, nil, nil, nil, err
	}
	err = json.Unmarshal(y, &ymap)
	if err != nil {
		return false, nil, nil, nil, err
	}
	// search for del keys and diff values
	for key, value := range xmap {
		val, ok := ymap[key]
		if !ok {
			delKeys = append(delKeys, key)
		} else if val != value {
			diffValueKeys = append(diffValueKeys, key)
		}
	}

	for key := range ymap {
		_, ok := xmap[key]
		if !ok {
			newKeys = append(newKeys, key)
		}
	}

	if len(newKeys) == 0 && len(delKeys) == 0 && len(diffValueKeys) == 0 {
		return true, nil, nil, nil, nil
	}
	return false, newKeys, delKeys, diffValueKeys, nil
}

type YamlCompares struct{}

// AreEqual checks all the items as arguments are equal or not.
// It only checks with the previous item. Does not check against each item as a bubble
// As soon as two items are different, it immediately return false,index,nil
// If all items are same then it returns true, 0, nil
// If there is no item or one item then it returns respective errors as false,0, ErrNoItem or false,0,ErrOnlyOneItem
// It is ideal to pass only two items to compare
func (c *YamlCompares) AreEqual(items ...[]byte) (bool, int, error) {
	if len(items) == 0 {
		return false, 0, ErrNoItem
	}
	if len(items) == 1 {
		return false, 0, ErrOnlyOneItem
	}
	equal := false
	data1 := items[0]
	for i := 1; i < len(items); i++ {
		data2 := items[i]
		if reflect.DeepEqual(data1, data2) {
			equal = true
			data1 = items[i]
		} else {
			return equal, i, nil
		}
	}
	return equal, 0, nil
}

// IsEqual is to check whether x and y are equal or not
// IsEqual returns bool, error
// if argument x or argument y is nil it returns false, "x is nil" or false, "y is nil"
func (c *YamlCompares) IsEqual(x []byte, y []byte) (bool, error) {
	if x == nil {
		return false, errors.New("x is nil item")
	}
	if y == nil {
		return false, errors.New("y is nil item")
	}
	if reflect.DeepEqual(x, y) {
		return true, nil
	}
	return false, nil
}

// Compare compares y against x provided by x and y are byte arrays
// the first return type is true of both of them are equal
// the second return type string slice returns new keys in y against x
// the third return type string slice returns deleted keys in y against x
// the fourth return type string slice returns common keys with different values in x and y
// the fifth return type error returns error if there is any error in data or in unmarshal
func (c *YamlCompares) Compare(x []byte, y []byte) (equal bool, newKeys []string, delKeys []string, diffValueKeys []any, err error) {
	if x == nil {
		return false, nil, nil, nil, errors.New("x is nil item")
	}
	if y == nil {
		return false, nil, nil, nil, errors.New("y is nil item")
	}

	xmap, ymap := make(map[string]any), make(map[string]any)
	newKeys = make([]string, 0)
	delKeys = make([]string, 0)
	diffValueKeys = make([]any, 0)
	err = yaml.Unmarshal(x, &xmap)
	if err != nil {
		return false, nil, nil, nil, err
	}
	err = yaml.Unmarshal(y, &ymap)
	if err != nil {
		return false, nil, nil, nil, err
	}
	// search for del keys and diff values
	for key, value := range xmap {
		val, ok := ymap[key]
		if !ok {
			delKeys = append(delKeys, key)
		} else if val != value {
			diffValueKeys = append(diffValueKeys, key)
		}
	}
	for key := range ymap {
		_, ok := xmap[key]
		if !ok {
			newKeys = append(newKeys, key)
		}
	}
	if len(newKeys) == 0 && len(delKeys) == 0 && len(diffValueKeys) == 0 {
		return true, nil, nil, nil, nil
	}
	return false, newKeys, delKeys, diffValueKeys, nil
}
