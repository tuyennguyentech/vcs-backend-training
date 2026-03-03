package main

import (
	"context"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	// Connect to PostgreSQL running as container in Podman
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Printf("%+v", db)
	ctx := context.Background()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	err = gorm.G[Product](db).Create(ctx, &Product{Code: "D42", Price: 100})

	// Read
	product, err := gorm.G[Product](db).Where("id = ?", 1).First(ctx)
	products, err := gorm.G[Product](db).Where("code = ?", "D42").Find(ctx)

	fmt.Println(product, products)

	// Update - update product's price to 200
	rowsAffected, err := gorm.G[Product](db).Where("id = ?", product.ID).Update(ctx, "Price", 200)
	fmt.Println(rowsAffected, err)
	// Update - update multiple fields
	rowsAffected, err = gorm.G[Product](db).Where("id = ?", product.ID).Updates(ctx, Product{Code: "D42", Price: 100})
	fmt.Println(rowsAffected, err)

	// Delete - delete product
	rowsAffected, err = gorm.G[Product](db).Where("id = ?", product.ID).Delete(ctx)
	fmt.Println(rowsAffected, err)
}
