package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

type Sale struct {
	Product int
	Volume  int
	Date    string
}

// String реализует метод интерфейса fmt.Stringer для Sale, возвращает строковое представление объекта Sale.
// Теперь, если передать объект Sale в fmt.Println(), то выведется строка, которую вернёт эта функция.
func (s Sale) String() string {
	return fmt.Sprintf("Product: %d Volume: %d Date:%s", s.Product, s.Volume, s.Date)
}

func selectSales(client int) ([]Sale, error) {
	var sales []Sale
	// Подключение базы данных.
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	// Соверешение sql-запроса.
	rows, err := db.Query("select product, volume, date from sales where id = :id", sql.Named("id", client))
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	// Чтение и запись значений из таблицы в массив sales.
	for rows.Next() {

		// Создаем объект newSales типа Sale.
		newSales := Sale{}

		err := rows.Scan(&newSales.Product, &newSales.Volume, &newSales.Date)
		if err != nil {
			log.Println(err)
		}

		sales = append(sales, newSales)
	}

	return sales, nil
}

func main() {
	client := 208

	sales, err := selectSales(client)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, sale := range sales {
		fmt.Println(sale)
	}
}
