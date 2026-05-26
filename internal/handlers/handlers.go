package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "index.html")
}
func LoadHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Ожидался запрос", http.StatusInternalServerError)
		return
	}
	file, fileHeader, err := req.FormFile("myFile")
	if err != nil {
		http.Error(w, "Не удалось получить файл: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)

	if err != nil {
		http.Error(w, "Не удалось прочитать файл", http.StatusInternalServerError)
		return
	}
	content := string(bytes)
	result, err := service.AutoDetect(content)
	if err != nil {
		http.Error(w, "Ошибка конвертации: "+err.Error(), http.StatusInternalServerError)
		return
	}
	ext := filepath.Ext(fileHeader.Filename)
	newFileName := time.Now().Format("2006-01-02_15-04-05") + ext
	err = os.WriteFile(newFileName, []byte(result), 0755)
	if err != nil {
		http.Error(w, "Ошибка записи файла: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(result)); err != nil {
		log.Printf("Ошибка при отправке ответа: %v", err)
	}
}
