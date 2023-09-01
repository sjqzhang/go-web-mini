package global

import (
	"context"
	"gorm.io/gorm"
	"sync"
)

var tx *transaction
var once sync.Once

type txFunc func(context.Context, func(context.Context) error) error

type transaction struct {
	handle txFunc
}

func (t *transaction) Execute(ctx context.Context, f func(c context.Context) error) error {
	return t.handle(ctx, f)
}

func NewTransaction(f txFunc) *transaction {
	return &transaction{handle: f}
}

func Context(ctx context.Context) *gorm.DB {
	return GetDB(ctx)
}


func GetTransaction() *transaction {
	once.Do(func() { // <-- atomic, does not allow repeating
		f := func(ctx context.Context, f func(context.Context) error) error {
			db := Context(ctx)
			return db.Transaction(func(tx *gorm.DB) error {
				c := context.WithValue(ctx, dbKey, tx)
				return f(c)
			})
		}
		tx = NewTransaction(f) // <-- thread safe
	})
	return tx
}


