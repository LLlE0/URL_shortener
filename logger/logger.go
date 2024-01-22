package logger

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func Remember(r *http.Request) {
	filename := fmt.Sprintf("../output/%s.log", time.Now().Format("20060102_150405"))

	// Создаем файл
	file, err := os.Create(filename)
	if err != nil {
		log.Printf("Unable to create file: %v\n", err)
		return
	}
	defer file.Close()

	logger := log.New(file, "", 0)

	logger.Printf("Method: %s\nURL: %s\nHeaders: %v\n", r.Method, r.URL, r.Header)
}
