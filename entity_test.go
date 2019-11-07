package core

import (
	"database/sql"
	"database/sql/driver"
	"testing"
)

type FooEntity struct {
	Id  int
	Bar *BarEntity
}

func (f *FooEntity) Get(*sql.Tx) error {
	f.Bar = &BarEntity{
		Id: f.Id + 1,
	}
	return nil
}

func (f *FooEntity) Post(tx *sql.Tx) error {
	return nil
}

func (f *FooEntity) Put(tx *sql.Tx) error {
	return nil
}

func (f *FooEntity) Delete(tx *sql.Tx) error {
	return nil
}

func (f *FooEntity) Scan(interface{}) error {
	return nil
}

func (f *FooEntity) Value() (driver.Value, error) {
	return nil, nil
}

type BarEntity struct {
	Id  int
	Baz BazEntity
}

func (b *BarEntity) Get(*sql.Tx) error {
	b.Baz = BazEntity{
		Id: b.Id + 1,
	}
	return nil
}

func (b *BarEntity) Post(tx *sql.Tx) error {
	return nil
}

func (b *BarEntity) Put(tx *sql.Tx) error {
	return nil
}

func (b *BarEntity) Delete(tx *sql.Tx) error {
	return nil
}

func (b *BarEntity) Scan(interface{}) error {
	return nil
}

func (b *BarEntity) Value() (driver.Value, error) {
	return nil, nil
}

type BazEntity struct {
	Id int
}

func TestGetAll(t *testing.T) {
	a := FooEntity{
		Id: 1,
	}
	err := GetAll(nil, &a)
	if err != nil {
		t.Error(err)
	}
	if a.Bar.Id != a.Id+1 {
		t.Log(a.Bar.Id)
		t.Fail()
	}
	if a.Bar.Baz.Id != a.Bar.Id+1 {
		t.Log(a.Bar.Baz.Id)
		t.Fail()
	}
}
