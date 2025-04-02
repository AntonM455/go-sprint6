package handlers

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// Функция MainHandler возвращает HTML из файла index.html
func MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Header().Add("Content-Type", "text/html")
	http.ServeFile(w, r, "./index.html")
}

// Функция HttpParcerHandler делает загрузку файла, конвертирует его с помощью вызова Convert и сохраняет результат
func HttpParcerHandler(w http.ResponseWriter, r *http.Request) {

	file, fileHeader, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "error when receiving file", http.StatusInternalServerError)
		return
	}
	// Закрыть файл после использования.
	defer file.Close()

	// Определить  расширение файла
	ext := filepath.Ext(fileHeader.Filename)

	// Прочитать весь полученный файл
	scanner := bufio.NewScanner(file)
	var fileData string
	for scanner.Scan() {
		fileData += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		http.Error(w, "error reading file", http.StatusInternalServerError)
		return
	}

	// Вызывается функция Convert из пакета service для конвертации
	convertedData, err := service.Convert(fileData)
	if err != nil {
		http.Error(w, "error during conversion", http.StatusInternalServerError)
		return
	}
	// Создание нового файла, определение имени и пути нового файла
	fileName := "indexres_" + time.Now().UTC().Format("20060102150405") + ext
	filePath := filepath.Join(".", fileName)

	// Создаеться новый файл и внего записываются данные
	outputFile, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "error creating file", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	// Запись данных во вновь созданный файл.
	if _, err = outputFile.Write([]byte(convertedData)); err != nil {
		http.Error(w, "error while writing file", http.StatusInternalServerError)
		return
	}

	// Возвращение результата пользователю в HTTP-ответе.
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	if _, err := w.Write([]byte(convertedData)); err != nil {
		log.Printf("Error sending response: %v", err)
	}
}
