package models

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB
var dbURL = "root:Test@1234@tcp(localhost:3306)/testdb?parseTime=true"

var err error

func GetDB() *gorm.DB {
	return DB
}

func InitDB() *gorm.DB {
	var db *gorm.DB
	db, err = gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	DB = db
	return DB
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&Category{}, &Subcategory{}, &Product{}, &WeeklyProductPrice{}, &User{}, &Subscription{})

}

func SaveRecord(dto interface{}) error {
	database := GetDB()
	err := database.Create(dto).Error
	return err
}

func UpdateRecord(model interface{}, paramMap map[string]string) error {
	database := GetDB()
	err := database.Model(model).Where(paramMap).Updates(model).Debug().Error
	return err
}

func FetchTableRowId(model interface{}, param interface{}) interface{} {
	database := GetDB()
	entity := database.Debug().Where(param).First(model)
	return entity

}

func Upsert(dto interface{}, columnNames []string, paramMap map[string]interface{}) error {
	database := GetDB()

	cols := make([]clause.Column, len(columnNames))
	for element := range columnNames {
		cols[element] = clause.Column{Name: columnNames[element]}
	}
	err := database.Clauses(clause.OnConflict{
		Columns:   cols,
		DoUpdates: clause.Assignments(paramMap),
	}).Create(dto).Error
	return err
}

func GetRecordsInStruct(query string, dto interface{}) {
	database := GetDB()
	database.Raw(query).Scan(dto)
}
