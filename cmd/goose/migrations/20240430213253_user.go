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
	goose.AddMigrationContext(upUser, downUser)
}

func upUser(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	err := db.DB.Set("gorm:table_options", TABLE_OPTIONS).AutoMigrate(&types.User{})
	if err != nil {
		log.Fatalf("error migrate order_item: %v", err)
		return err
	}
	return nil
}

func downUser(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	err := db.DB.Migrator().DropTable(&types.User{})
	if err != nil {
		log.Fatalf("error drop order_item: %v", err)
		return err
	}
	return nil
}
