package db

import (
	"gopkg.in/mgo.v2"
	"reflect"
)

// Collection - Struct
type Collection struct {
	col   *mgo.Collection
	model interface{}
	count int
}

// Find - Return queried results list
func (c *Collection) Find(options QueryOptions) (interface{}, error) {
	modelType := reflect.New(reflect.TypeOf(c.model))
	err := c.col.Find(options.GetFilters()).Limit(options.Amount).All(modelType.Interface())
	if err != nil {
		return nil, err
	}
	return modelType.Elem().Interface(), nil
}

// All - Return all results
func (c *Collection) All() (interface{}, error) {
	modelType := reflect.New(reflect.TypeOf(c.model))
	err := c.col.Find(nil).All(modelType.Interface())
	if err != nil {
		return nil, err
	}
	return modelType.Elem().Interface(), nil
}
