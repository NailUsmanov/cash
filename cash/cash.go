package Cache

import (
	"errors"
)

type Cache struct {
	Item map[string]interface{}
}

func New () *Cache {
	return &Cache{
		make(map[string]interface{}),
	}
}

func (c*Cache) Set (key string, value interface{}) error {
	if key == "" {
		return errors.New ("key is empty")
	}else{
		c.Item[key] = value
		return nil
	}
}

func (c*Cache) Get (key string)(interface{},error) {
	value, exists := c.Item[key]
	if exists {
		return value, nil
	}
return value,errors.New("error: file is not exists")
}

func (c*Cache) Delete (key string)  error{
	_, exists := c.Item[key]
	if exists {
		delete(c.Item,key)
		return nil	
	}else{
		return errors.New("already don't have this file")
	}
}

