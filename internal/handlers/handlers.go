package handlers

import (
	"bufio"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// Функция MainHandler возвращает HTML из файла index.html
func MainHandler(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("C:\\Users\\User\\Dev\\go-sprint6\\index.html")
	if err != nil {
		http.Error(w, "ошибка при загрузке из файла", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// Функция HttpParcerHandler делает загрузку файла, конвертирует его с помощью вызова Convert и сохраняет результат
func HttpParcerHandler(w http.ResponseWriter, r *http.Request) {

	file, fileHeader, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "ошибка при получении файла", http.StatusInternalServerError)
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
		http.Error(w, "ошибка при чтении файла", http.StatusInternalServerError)
		return
	}

	// Вызывается функция Convert из пакета service для конвертации
	convertedData, err := service.Convert(fileData)
	if err != nil {
		http.Error(w, "ошибка при конвертации", http.StatusInternalServerError)
		return
	}
	// Создание нового файла, определение имени и пути нового файла
	fileName := "indexres_" + time.Now().UTC().Format("20060102150405") + ext
	filePath := filepath.Join("C:\\Users\\User\\Dev\\go-sprint6", fileName)

	// Создаеться новый файл и внего записываются данные
	outputFile, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "ошибка при создании файла", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	// Запись данных во вновь созданный файл.
	_, err = outputFile.WriteString(convertedData)
	if err != nil {
		http.Error(w, "ошибка при записи файла", http.StatusInternalServerError)
		return
	}

	// Возвращение результата пользователю
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Файл успешно обработан. Результат сохранён в " + filePath))
}
