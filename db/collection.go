package db

import (
	"gopkg.in/mgo.v2"
	"math/rand"
	"reflect"
)

// QueryOptions - Search options
type QueryOptions struct {
	Amount  int
	Filters map[string]string
}

// Collection - Struct
type Collection struct {
	col   *mgo.Collection
	model interface{}
	count int
}

// FindRandom - Return random results list
func (c *Collection) FindRandom(options QueryOptions) (interface{}, error) {
	modelType := reflect.New(reflect.TypeOf(c.model))
	skip := c.makeRandomSkip()
	err := c.col.Find(options.Filters).Limit(options.Amount).Skip(skip).All(modelType.Interface())
	if err != nil {
		return nil, err
	}
	return modelType.Elem().Interface(), nil
}

func (c *Collection) makeRandomSkip() int {
	if c.count == 0 {
		return 1
	}
	return rand.Intn(c.count)
}
