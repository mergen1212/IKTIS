package dbInter

import (
	"Go+HTMX/dto"
	"encoding/json"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"net/http"
)

func Conn() (*gorm.DB, error) {
	dsn := "host=localhost user=habrpguser password=pgpwd4habr dbname=habrdb port=5432 sslmode=disable TimeZone=Europe/Moscow"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&dto.YourModel{})
	return db, nil
}

func FirstWeek(db *gorm.DB, kok *dto.YourModel, param1 string, week int) error {
	db.Where("id = ? AND week = ?", param1, week).First(&kok)
	return db.Error
}

func Parse(db *gorm.DB) (*gorm.DB, error) {
	for w := 1; w < 12; w++ {
		for i := 1; i < 113; i++ {
			url := fmt.Sprintf("https://ictis.ru/api?request=schedule&group=%d.html&week=%d", i, w)
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("Ошибка при выполнении GET-запроса:", err)

			}
			defer resp.Body.Close()

			responseBody, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Ошибка при чтении ответа:", err)

			}
			var T dto.T
			json.Unmarshal(responseBody, &T)
			newRecord := dto.YourModel{ID: T.Table.Name, Week: w, DATA: string(responseBody)}
			db.Create(newRecord)
		}
	}
	return db, nil
}
