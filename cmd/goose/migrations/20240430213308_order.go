package migrations

import (
	"context"
	"database/sql"
	"log"

	"github.com/pressly/goose/v3"

	"github.com/dodo/ecom/db"
	"github.com/dodo/ecom/types"
)

func init() {
	goose.AddMigrationContext(upOrders, downOrders)
}

func upOrders(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	err := db.DB.Set("gorm:table_options", TABLE_OPTIONS).AutoMigrate(&types.Order{})
	if err != nil {
		log.Fatalf("error migarate order: %v", err)
		return err
	}
	return nil
}

func downOrders(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	err := db.DB.Migrator().DropTable(&types.Order{})
	if err != nil {
		log.Fatalf("erro drop table order: %v", err)
		return err
	}
	return nil
}
