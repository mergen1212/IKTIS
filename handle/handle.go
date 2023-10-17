package handle

import (
	"Go+HTMX/dbInter"
	"Go+HTMX/dto"
	"encoding/json"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

func YourHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		param1 := r.URL.Query().Get("gru")
		param2 := r.URL.Query().Get("week")
		week, err := strconv.Atoi(param2)
		if err != nil {
			log.Println(err)
			http.Error(w, "не валидные параметры", http.StatusUnprocessableEntity)
			return
		}
		var kok dto.YourModel
		if err := dbInter.FirstWeek(db, &kok, param1, week); err != nil {
			log.Println(err)
			http.Error(w, "NotFound", http.StatusNotFound)
			return
		}
		var jsonData dto.T
		if err := json.Unmarshal([]byte(kok.DATA), &jsonData); err != nil {
			log.Println(err)
			http.Error(w, "ошибка декодирования", http.StatusInternalServerError)
			return
		}
		Q, err := json.Marshal(jsonData)
		if err != nil {
			log.Println(err)
			http.Error(w, "ошибка декодирования", http.StatusInternalServerError)
			return
		}
		w.Write(Q)
	}
}

func YourHandle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("kkkkk"))
	}
}
