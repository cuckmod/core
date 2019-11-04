package core

import (
	"database/sql"
	"database/sql/driver"
	"errors"
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

func Get(tx *sql.Tx, query string, ID int, dest ...interface{}) (err error) {
	stmt, err := tx.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(ID)
	return row.Scan(dest...)
}

func Post(tx *sql.Tx, query string, args ...interface{}) (ID int, err error) {
	stmt, err := tx.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(args...)
	if err != nil {
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		return
	}
	ID = int(id)
	return
}

func Put(tx *sql.Tx, query string, args ...interface{}) error {
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(args...)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected != 1 {
		return errors.New("can't update entity")
	}
	return nil
}

func Delete(tx *sql.Tx, query string, ID int) error {
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(ID)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected != 1 {
		return errors.New("can't delete entity")
	}
	return nil
}

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
