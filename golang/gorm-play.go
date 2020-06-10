package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "L1212", Price: 1000})

	// Read
	var product Product
	db.First(&product, 1)                   // find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	db.Delete(&product)
}

/*
Model Example:

type User struct {
  gorm.Model
  Name         string
  Age          sql.NullInt64
  Birthday     *time.Time
  Email        string  `gorm:"type:varchar(100);unique_index"`
  Role         string  `gorm:"size:255"` // set field size to 255
  MemberNumber *string `gorm:"unique;not null"` // set member number to unique and not null
  Num          int     `gorm:"AUTO_INCREMENT"` // set num to auto incrementable
  Address      string  `gorm:"index:addr"` // create index with name `addr` for address
  IgnoreMe     int     `gorm:"-"` // ignore this field
}
*/
/*
Struct tags

Tags are optional to use when declaring models. GORM supports the following tags:

Supported Struct tags

Tag				    Description
Column 			    Specifies column name
Type 			    Specifies column data type
Size 			    Specifies column size, default 255
PRIMARY_KEY 	    Specifies column as primary key
UNIQUE 	            Specifies column as unique
DEFAULT 	        Specifies column default value
PRECISION 	        Specifies column precision
NOT NULL 	        Specifies column as NOT NULL
AUTO_INCREMENT 	    Specifies column auto incrementable or not
INDEX 	            Create index with or without name, same name creates composite indexes
UNIQUE_INDEX 	    Like INDEX, create unique index
EMBEDDED 	        Set struct as embedded
EMBEDDED_PREFIX 	Set embedded struct’s prefix name
- 	                Ignore this fields
*/
/*
Struct tags for Associations

Check out the Associations section for details
Tag									Description
MANY2MANY							Specifies join table name
FOREIGNKEY							Specifies foreign key
ASSOCIATION_FOREIGNKEY				Specifies association foreign key
POLYMORPHIC							Specifies polymorphic type
POLYMORPHIC_VALUE					Specifies polymorphic value
JOINTABLE_FOREIGNKEY				Specifies foreign key of jointable
ASSOCIATION_JOINTABLE_FOREIGNKEY 	Specifies association foreign key of jointable
SAVE_ASSOCIATIONS					AutoSave associations or not
ASSOCIATION_AUTOUPDATE				AutoUpdate associations or not
ASSOCIATION_AUTOCREATE				AutoCreate associations or not
ASSOCIATION_SAVE_REFERENCE			AutoSave associations reference or not
PRELOAD								Auto Preload associations or not
*/
