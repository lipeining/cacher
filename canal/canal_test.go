package canal

import (
	"fmt"
	"testing"

	"github.com/go-mysql-org/go-mysql/replication"
	"github.com/go-mysql-org/go-mysql/schema"
)

type Product struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type ProductListener struct{}

func (l *ProductListener) Scan(table *schema.Table, rows []any) *Product {
	product := &Product{}
	for index, column := range table.Columns {
		if column.Name == "id" {
			product.ID = rows[index].(int32)
		} else if column.Name == "name" {
			product.Name = rows[index].(string)
		}
	}
	return product
}

func (l *ProductListener) Insert(row *Product, header *replication.EventHeader) error {
	fmt.Println("Insert", row)
	return nil
}

func (l *ProductListener) Update(before *Product, after *Product, header *replication.EventHeader) error {
	fmt.Println("Update", before, after)
	return nil
}

func (l *ProductListener) Delete(row *Product, header *replication.EventHeader) error {
	fmt.Println("Delete", row)
	return nil
}

func TestCanal(t *testing.T) {
	canal := NewCanal("canal.toml",
		WithWhere("id > 0"),
	)

	listener := NewWrapListener[Product](&ProductListener{})

	canal.AddListener("product1", listener)

	canal.Start()
}
