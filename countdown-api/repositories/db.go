package repositories

import (
	"countdown/countdown-api/requests"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var err error

func init() {
	DB, err = DBInit()

	if err != nil {
		log.Fatal(err)
	}
}
func DBInit() (*gorm.DB, error) {

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_DB"),
	)
	for i := 0; i < 10; i++ {
		DB, err = gorm.Open("postgres", dbinfo) // gorm checks Ping on Open
		if err == nil {
			break
		}
		log.Printf("Error while connecting to DB : %s", err.Error())
		log.Println("Trying to connect ... waiting 5 seconds")
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		return DB, err
	}

	//create tables at start
	err = DB.AutoMigrate(&requests.Task{}).Error
	if err != nil {
		log.Fatal(err)
	}

	return DB, err
}

func CreateTask(t *requests.Task) error {
	if err := DB.Create(t).Error; err != nil {
		return err
	}

	return nil
}

func GetTasks() ([]requests.Task, error) {
	var tasks []requests.Task

	if err := DB.Order("timestamp asc").Find(&tasks).Error; err != nil {
		return tasks, err
	}

	return tasks, err
}
