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
	goose.AddMigrationContext(upProducts, downProducts)
}

func upProducts(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	err := db.DB.Set("gorm:table_options", TABLE_OPTIONS).AutoMigrate(&types.Product{})
	if err != nil {
		log.Fatalf("error migrate product: %v", err)
		return err
	}

	return nil
}

func downProducts(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	err := db.DB.Migrator().DropTable(&types.Product{})
	if err != nil {
		log.Fatalf("error drop table products: %v", err)
	}
	return nil
}
