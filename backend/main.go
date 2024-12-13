package main
import "strings"

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Apartment представляет информацию о квартире
type Apartment struct {
	ID          int
	Title       string
	Address     string
	ImageLink   string
	Description string
	SquareMeters int
	Bedrooms     int
	Price        int
	Favourite    bool
}

// Пример списка квартир
var apartments = []Apartment{
	{ID: 1, Title: "Однокомнатная квартира", Address: "ул. Ленина, 12", ImageLink: "https://static1.abitant.com/uploads/project_image/115/9682/a5dgdgialh5hzvlo164u.jpg", Description: "Уютная однокомнатная квартира в центре города.", SquareMeters: 45, Bedrooms: 1, Price: 25000, Favourite: false},
	{ID: 2, Title: "Двухкомнатная квартира", Address: "ул. Гагарина, 8", ImageLink: "https://static2.abitant.com/uploads/project_image/193/6020/b1eop4yu8e2gyh2tspm3.jpg", Description: "Просторная двухкомнатная квартира с видом на парк.", SquareMeters: 65, Bedrooms: 2, Price: 40000, Favourite: false},
	{ID: 3, Title: "Однокомнатная квартира", Address: "ул. Ленина, 12", ImageLink: "https://cdn0.youla.io/files/images/780_780/63/29/6329d9f543eedb62b7695786-1.jpg", Description: "Уютная однокомнатная квартира в центре города.", SquareMeters: 45, Bedrooms: 1, Price: 25000, Favourite: false},
    {ID: 4, Title: "Двухкомнатная квартира", Address: "ул. Гагарина, 8", ImageLink: "https://www.amocrm.ru/uploads/2019/06/huszzitbio4.jpg", Description: "Просторная двухкомнатная квартира с видом на парк.", SquareMeters: 65, Bedrooms: 2, Price: 40000, Favourite: false},
    {ID: 5, Title: "Однокомнатная квартира", Address: "ул. Ленина, 12", ImageLink: "https://static1.abitant.com/uploads/project_image/115/9682/a5dgdgialh5hzvlo164u.jpg", Description: "Уютная однокомнатная квартира в центре города.", SquareMeters: 45, Bedrooms: 1, Price: 25000, Favourite: false},
    {ID: 6, Title: "Двухкомнатная квартира", Address: "ул. Гагарина, 8", ImageLink: "https://static.tildacdn.com/tild6231-6263-4235-b834-353836373433/IMG_0374__--1_k_.jpg", Description: "Просторная двухкомнатная квартира с видом на парк.", SquareMeters: 65, Bedrooms: 2, Price: 40000, Favourite: false},
    {ID: 7, Title: "Однокомнатная квартира", Address: "ул. Ленина, 12", ImageLink: "https://bigfoto.name/uploads/posts/2022-02/1645541861_34-bigfoto-name-p-kvartira-s-krasivim-vidom-72.jpg", Description: "Уютная однокомнатная квартира в центре города.", SquareMeters: 45, Bedrooms: 1, Price: 25000, Favourite: false},
    {ID: 8, Title: "Двухкомнатная квартира", Address: "ул. Гагарина, 8", ImageLink: "https://attaches.1001tur.ru/hotels/gallery/651238/28891598889415.jpg", Description: "Просторная двухкомнатная квартира с видом на парк.", SquareMeters: 65, Bedrooms: 2, Price: 40000, Favourite: false},
}

// обработчик для GET-запроса, возвращает список квартир
func getApartmentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(apartments)
}

// обработчик для POST-запроса, добавляет квартиру
func createApartmentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var newApartment Apartment
	err := json.NewDecoder(r.Body).Decode(&newApartment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newApartment.ID = len(apartments) + 1
	apartments = append(apartments, newApartment)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newApartment)
}

// обработчик для получения квартиры по ID
func getApartmentByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/apartments/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Apartment ID", http.StatusBadRequest)
		return
	}

	for _, apartment := range apartments {
		if apartment.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(apartment)
			return
		}
	}

	http.Error(w, "Apartment not found", http.StatusNotFound)
}

// удаление квартиры по ID
func deleteApartmentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/apartments/delete/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Apartment ID", http.StatusBadRequest)
		return
	}

	for i, apartment := range apartments {
		if apartment.ID == id {
			apartments = append(apartments[:i], apartments[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Apartment not found", http.StatusNotFound)
}

// обновление информации о квартире по ID
func updateApartmentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	parts := strings.Split(r.URL.Path, "/")
    if len(parts) < 3 {
        http.Error(w, "Invalid Apartment ID", http.StatusBadRequest)
        return
    }
    idStr := parts[2]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Apartment ID", http.StatusBadRequest)
		return
	}

	var updatedApartment Apartment
	err = json.NewDecoder(r.Body).Decode(&updatedApartment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, apartment := range apartments {
		if apartment.ID == id {
			apartments[i] = updatedApartment
			apartments[i].ID = id

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(apartments[i])
			return
		}
	}

	http.Error(w, "Apartment not found", http.StatusNotFound)
}
// Обработчик для изменения статуса Favourite
func toggleFavouriteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Извлекаем ID квартиры из URL
	idStr := strings.TrimPrefix(r.URL.Path, "/apartments/favourite/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Apartment ID", http.StatusBadRequest)
		return
	}

	// Поиск квартиры по ID
	for i, apartment := range apartments {
		if apartment.ID == id {
			// Переключаем состояние Favourite
			apartments[i].Favourite = !apartments[i].Favourite

			// Возвращаем обновленный объект
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(apartments[i])
			return
		}
	}

	http.Error(w, "Apartment not found", http.StatusNotFound)
}


func main() {
	http.HandleFunc("/apartments", getApartmentsHandler)             // Получить все квартиры
	http.HandleFunc("/apartments/create", createApartmentHandler)    // Создать квартиру
	http.HandleFunc("/apartments/", getApartmentByIDHandler)         // Получить квартиру по ID
	http.HandleFunc("/apartments/update/", updateApartmentHandler)   // Обновить квартиру
	http.HandleFunc("/apartments/delete/", deleteApartmentHandler)   // Удалить квартиру
    http.HandleFunc("/apartments/favourite/", toggleFavouriteHandler) // Изменить Favourite
	fmt.Println("Server is running on port 8080!")
	http.ListenAndServe(":8080", nil)
}