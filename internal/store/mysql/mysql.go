package mysql

import (
	"fmt"

	"github.com/gityokie/gocrud/internal/models"
	"github.com/gityokie/gocrud/internal/utils"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlDB struct {
	Db *gorm.DB
}

func New() (*MysqlDB, error) {
	user := viper.GetString(utils.DbUser)
	pass := viper.GetString(utils.DbPwd)
	port := viper.GetString(utils.DbPort)
	dbName := viper.GetString(utils.DbName)

	var dsn string

	dsn = fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, port, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Book{})
	if err != nil {
		return nil, err
	}

	return &MysqlDB{Db: db}, nil
}

func (m MysqlDB) GetAllBook(b *[]models.Book) (err error) {
	if err = m.Db.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func (m MysqlDB) AddNewBook(b *models.Book) (err error) {
	if err = m.Db.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func (m MysqlDB) GetOneBook(b *models.Book, id string) (err error) {
	if err := m.Db.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func (m MysqlDB) PutOneBook(b *models.Book, id string) (err error) {
	fmt.Println(b)
	m.Db.Save(b)
	return nil
}

func (m MysqlDB) DeleteBook(b *models.Book, id string) (err error) {
	m.Db.Where("id = ?", id).Delete(b)
	return nil
}
