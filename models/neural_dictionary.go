package models

import (
	"fmt"
	"database/sql"
)

type NeuralDictionary struct {
	Id int64 `json:"id"`
	Title string `json:"title"`
	Binary string `json:"binary"`
}


// Получение всех категории ответа нейронной сети
func (self * SessionDb) GetDictionaryes() ([]NeuralDictionary,error) {
	var (
		cat []NeuralDictionary
		id sql.NullInt64
		title sql.NullString
		binary sql.NullString
	)
	rows,err := self.GetDb().Query("SELECT id,title,binary FROM neural_dictionary")
	if err != nil {
		fmt.Println("GetReplys Query Error",err)
		return cat,err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id,&title,&binary)
		if err != nil {
			fmt.Println("GetReplys Scan Error",err)
			return cat,err
		}
		cat = append(cat,NeuralDictionary{ Id : id.Int64, Title : title.String, Binary : binary.String })
	}
	return cat,nil
}
// Получение категории ответа нейронной сети
func (self * SessionDb) GetDictionary(idcat int64) (NeuralDictionary,error) {
	var (
		cat NeuralDictionary
		id sql.NullInt64
		title sql.NullString
		binary sql.NullString
	)
	err := self.GetDb().QueryRow("SELECT id,title,binary FROM neural_dictionary WHERE id = ? ",idcat).Scan(&id,&title,&binary)
	if err != nil {
		fmt.Println("GetReply Error",err)
		return cat,err
	}
	cat.Id = id.Int64
	cat.Title = title.String
	cat.Binary = binary.String
	return cat,nil
}
// Поиск ответа нейронной сети
func (self * SessionDb) FindDictionary(search string) (NeuralDictionary,error) {
	var (
		cat NeuralDictionary
		id sql.NullInt64
		title sql.NullString
		binary sql.NullString
	)
	cat.Id = 2
	err := self.GetDb().QueryRow("SELECT id,title,binary FROM neural_dictionary WHERE title LIKE '%"+search+"%'").Scan(&id,&title,&binary)
	if err != nil {
		fmt.Println("FindBinary Error",err)
		return cat,err
	}
	cat.Id = id.Int64
	cat.Title = title.String
	cat.Binary = binary.String
	return cat,nil
}
// Получение категории ответа нейронной сети
func (self * SessionDb) SaveDictionary(title string,binary string) (NeuralDictionary,error) {
	cat := NeuralDictionary{}
	cat.Binary = binary
	smtp,err := self.GetDb().Prepare("INSERT INTO neural_dictionary( title, binary ) VALUES( ?,? )")

	defer smtp.Close()

	if err != nil {
		fmt.Println("SaveReply Prepare Error",err)
		return cat,err
	}
	res,err := smtp.Exec(title,binary)
	if err != nil {
		fmt.Println("SaveReply Exec Error",err)
		return cat,err
	}
	cat.Id,err = res.LastInsertId()
	if err != nil {
		fmt.Println("SaveReply LastInsertId Error",err)
		return cat,err
	}
	return cat,nil
}

func (self * SessionDb) UpdateDictionary(cat NeuralDictionary) (NeuralDictionary,error){

	smtp,err := self.GetDb().Prepare("UPDATE neural_dictionary SET title = ?,binary = ? WHERE id = ?")

	defer smtp.Close()

	if err != nil {
		fmt.Println("UpdateReply Prepare Error",err)
		return cat,err
	}
	_,err = smtp.Exec(cat.Title,cat.Binary,cat.Id)
	if err != nil {
		fmt.Println("UpdateReply Exec Error",err)
		return cat,err
	}
	return cat,nil
}

func (self * SessionDb) DeleteDictionary(cat NeuralDictionary) bool{
	smtp,err := self.GetDb().Prepare("DELETE FROM neural_dictionary WHERE id = ?")
	defer smtp.Close()
	if err != nil {
		fmt.Println("DeleteReply Prepare Error",err)
		return false
	}
	_,err = smtp.Exec(cat.Id)
	if err != nil {
		fmt.Println("DeleteReply Exec Error",err)
		return false
	}
	return true
}
func (self * SessionDb) DeleteDictionaryAll() bool{
	smtp,err := self.GetDb().Prepare("DELETE FROM neural_dictionary")
	defer smtp.Close()
	if err != nil {
		fmt.Println("DropBinary Prepare Error",err)
		return false
	}
	_,err = smtp.Exec()
	if err != nil {
		fmt.Println("DropBinary Exec Error",err)
		return false
	}
	return true
}
