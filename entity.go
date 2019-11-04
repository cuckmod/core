package core

import (
	"database/sql"
	"database/sql/driver"
	"reflect"
	"sync"
)

type (
	Entity interface {
		Get(tx *sql.Tx) error
		Post(tx *sql.Tx) error
		Put(tx *sql.Tx) error
		Delete(tx *sql.Tx) error

		// SQL
		Scan(interface{}) error
		Value() (driver.Value, error)
	}
)

func GetAll(tx *sql.Tx, e Entity) (err error) {
	errChan := make(chan error)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		recursiveGet(&wg, errChan, tx, e)
	}()
	go func() {
		wg.Wait()
		close(errChan)
	}()
	for err := range errChan {
		if err != nil {
			return err
		}
	}
	return
}

func recursiveGet(wg *sync.WaitGroup, errChan chan error, tx *sql.Tx, e Entity) {
	if err := e.Get(tx); err != nil {
		errChan <- err
		return
	}
	modelValue := reflect.ValueOf(e).Elem()
	for i := 0; i < modelValue.NumField(); i++ {
		field := modelValue.Field(i)
		if field.Type().Implements(reflect.ValueOf(new(Entity)).Type().Elem()) {
			wg.Add(1)
			go func() {
				defer wg.Done()
				recursiveGet(wg, errChan, tx, field.Interface().(Entity))
			}()
		}
	}
}
