package model

import (
	"github.com/aaalik/ke-jepang/bootstrap"
	"github.com/aaalik/ke-jepang/helper"
	"github.com/aaalik/ke-jepang/structs"
)

func GetSingleItem(id int) (structs.Item, error) {
	sql := `
		SELECT
			id,
			name,
			price
		FROM
			items
		WHERE
			id = ?
	`

	rows := bootstrap.DB.QueryRow(sql, id)

	result := structs.Item{}

	err := rows.Scan(
		&result.Id,
		&result.Name,
		&result.Price,
	)

	return result, err
}

func GetItems() []structs.Item {
	sql := `
		SELECT
			id,
			name,
			price
		FROM
			items
	`

	rows, err := bootstrap.DB.Query(sql)
	if err != nil {
		helper.Log.Error(err)
	}

	results := []structs.Item{}

	for rows.Next() {
		row := structs.Item{}

		err := rows.Scan(
			&row.Id,
			&row.Name,
			&row.Price,
		)

		if err != nil {
			helper.Log.Error(err)
		}

		results = append(results, row)
	}

	return results
}

func SaveItem(name string, price int) (bool, error) {
	sql := `
		INSERT INTO items(name, price) VALUES(?,?)
	`

	rows, err := bootstrap.DB.Prepare(sql)

	if err != nil {
		helper.Log.Error(err)
		return false, err
	}

	rows.Exec(name, price)

	return true, err
}

func UpdateItem(id int, name string, price int) (bool, error) {
	sql := `
		UPDATE items SET name = ? , price = ? WHERE id = ?
	`

	rows, err := bootstrap.DB.Prepare(sql)

	if err != nil {
		helper.Log.Error(err)
		return false, err
	}

	rows.Exec(name, price, id)

	return true, err
}

func DeleteItem(id int) (bool, error) {
	sql := `
		DELETE FROM items WHERE id = ?
	`

	rows, err := bootstrap.DB.Prepare(sql)

	if err != nil {
		helper.Log.Error(err)
		return false, err
	}

	rows.Exec(id)

	return true, err
}
