package logger

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"time"
)

func Remember(r *http.Request) {
	//Filename of each log consists of dummy URL and milliseconds (always unique)
	filename := fmt.Sprintf("../output/%s_%s.log", r.URL, fmt.Sprint(time.Now().UnixMilli()))
	file, err := os.Create(filename)
	if err != nil {
		log.Printf("Unable to create file: %v\n", err)
		return
	}
	defer file.Close()

	//Custom output for logger
	logger := log.New(file, "", 0)

	//using reflections to get all the exported fields for request
	v := reflect.ValueOf(r).Elem()
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.CanInterface() {
			logger.Printf("%s: %v\n", t.Field(i).Name, field.Interface())
		} else {

		}
	}
}
