package models

import (
	"fmt"
	"database/sql"
)

type Category struct {
	Id int64 `json:"id"`
	Title string `json:"title"`
}

// Получение всех категории ответа нейронной сети
func (self * SessionDb) GetCategories() ([]Category,error) {
	var (
		cat []Category
		id sql.NullInt64
		title sql.NullString
	)
	rows,err := self.GetDb().Query("SELECT id,title FROM neural_category")
	if err != nil {
		fmt.Println("GetCategories Query Error",err)
		return cat,err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id,&title)
		if err != nil {
			fmt.Println("GetCategories Scan Error",err)
			return cat,err
		}
		cat = append(cat,Category{ Id : id.Int64, Title : title.String })
	}
	return cat,nil
}
// Получение категории ответа нейронной сети
func (self * SessionDb) GetCategory(idcat int64) (Category,error) {
	var (
		cat Category
		id sql.NullInt64
		title sql.NullString
	)
	err := self.GetDb().QueryRow("SELECT id,title FROM neural_category WHERE id = ? ",idcat).Scan(&id,&title)
	if err != nil {
		fmt.Println("GetCategory Error",err)
		return cat,err
	}
	cat.Id = id.Int64
	cat.Title = title.String
	return cat,nil
}
// Получение категории ответа нейронной сети
func (self * SessionDb) SaveCategory(title string) (Category,error) {
	cat := Category{}
	cat.Title =title
	smtp,err := self.GetDb().Prepare("INSERT INTO neural_category( title ) VALUES( ? )")

	defer smtp.Close()

	if err != nil {
		fmt.Println("SaveCategory Prepare Error",err)
		return cat,err
	}
	res,err := smtp.Exec(title)
	if err != nil {
		fmt.Println("SaveCategory Exec Error",err)
		return cat,err
	}
	cat.Id,err = res.LastInsertId()
	if err != nil {
		fmt.Println("SaveCategory LastInsertId Error",err)
		return cat,err
	}
	return cat,nil
}

func (self * SessionDb) UpdateCategory(cat Category) (Category,error){

	smtp,err := self.GetDb().Prepare("UPDATE neural_category SET title = ? WHERE id = ?")

	defer smtp.Close()

	if err != nil {
		fmt.Println("SaveCategory Prepare Error",err)
		return cat,err
	}
	_,err = smtp.Exec(cat.Title,cat.Id)
	if err != nil {
		fmt.Println("SaveCategory Exec Error",err)
		return cat,err
	}
	return cat,nil
}

func (self * SessionDb) DeleteCategory(cat Category) bool{
	smtp,err := self.GetDb().Prepare("DELETE FROM neural_category WHERE id = ?")
	defer smtp.Close()
	if err != nil {
		fmt.Println("SaveCategory Prepare Error",err)
		return false
	}
	_,err = smtp.Exec(cat.Id)
	if err != nil {
		fmt.Println("SaveCategory Exec Error",err)
		return false
	}
	return true
}