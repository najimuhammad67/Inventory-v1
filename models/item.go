package models

import "inventory-app/config"

type Item struct {
	ID       int
	Name     string
	Category string
	Quantity int
}

func GetAllItems() []Item {
	rows, _ := config.DB.Query("SELECT id, name, category, quantity FROM inventory")
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var i Item
		rows.Scan(&i.ID, &i.Name, &i.Category, &i.Quantity)
		items = append(items, i)
	}
	return items
}

func InsertItem(name, category string, quantity int) error {
	_, err := config.DB.Exec("INSERT INTO inventory (name, category, quantity) VALUES (?, ?, ?)", name, category, quantity)
	return err
}

func UpdateItem(id int, name, category string, quantity int) error {
	_, err := config.DB.Exec("UPDATE inventory SET name = ?, category = ?, quantity = ? WHERE id = ?", name, category, quantity, id)
	return err
}

func DeleteItem(id int) error {
	_, err := config.DB.Exec("DELETE FROM inventory WHERE id = ?", id)
	return err
}
