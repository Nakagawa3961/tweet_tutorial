package repositories

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/ryuta06012/tweet_backend/src/configs"
)

var dsn = ""

func init() {
	mysqlConfig := configs.NewConfig()
	dsn = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		mysqlConfig.Mysql.User,
		mysqlConfig.Mysql.Password,
		mysqlConfig.Mysql.Host,
		mysqlConfig.Mysql.Port,
		mysqlConfig.Mysql.DatabaseName,
	)
}

type MysqlRepository struct {
	Client *gorm.DB
}

func NewMysqlHandler() *MysqlRepository {
	//DB接続
	fmt.Printf("dsn: %v\n", dsn)
	client, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}

	return &MysqlRepository{
		Client: client,
	}
}

func (db *MysqlRepository) Begin() *gorm.DB {
	return db.Client.Begin()
}

type Repository interface {
	SetClient(*gorm.DB)
}

func (db *MysqlRepository) WithTransaction(repositories []Repository, dofunc func() error) error {
	// NOTE: commitしてからClientを元に戻さないと"transaction has already been committed or rolled back"エラーになる
	defer func() {
		for _, repository := range repositories {
			repository.SetClient(db.Client)
		}
	}()

	return db.Client.Transaction(func(tx *gorm.DB) error {
		for _, repository := range repositories {
			repository.SetClient(tx)
		}
		return dofunc()
	})
}
