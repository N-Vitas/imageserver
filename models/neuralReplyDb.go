package models

import (
	"fmt"
	"database/sql"
)

type NeuralRandomReply struct {
	Id int64 `json:"id"`
	CatId int64 `json:"catId"`
	Title string `json:"title"`
}


// Получение всех категории ответа нейронной сети
func (self * SessionDb) GetReplys() ([]NeuralRandomReply,error) {
	var (
		cat []NeuralRandomReply
		id sql.NullInt64
		id_neural_cat sql.NullInt64
		title sql.NullString
	)
	rows,err := self.GetDb().Query("SELECT id,id_neural_cat,title FROM neural_random_reply")
	if err != nil {
		fmt.Println("GetReplys Query Error",err)
		return cat,err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id,&id_neural_cat,&title)
		if err != nil {
			fmt.Println("GetReplys Scan Error",err)
			return cat,err
		}
		cat = append(cat,NeuralRandomReply{ Id : id.Int64, CatId : id_neural_cat.Int64, Title : title.String })
	}
	return cat,nil
}
// Получение категории ответа нейронной сети
func (self * SessionDb) GetReply(idcat int64) (NeuralRandomReply,error) {
	var (
		cat NeuralRandomReply
		id sql.NullInt64
		id_neural_cat sql.NullInt64
		title sql.NullString
	)
	err := self.GetDb().QueryRow("SELECT id,id_neural_cat,title FROM neural_random_reply WHERE id = ? ",idcat).Scan(&id,&id_neural_cat,&title)
	if err != nil {
		fmt.Println("GetReply Error",err)
		return cat,err
	}
	cat.Id = id.Int64
	cat.CatId = id_neural_cat.Int64
	cat.Title = title.String
	return cat,nil
}
// Получение категории ответа нейронной сети
func (self * SessionDb) SaveReply(catId int64,title string) (NeuralRandomReply,error) {
	cat := NeuralRandomReply{}
	cat.Title =title
	smtp,err := self.GetDb().Prepare("INSERT INTO neural_random_reply( id_neural_cat, title ) VALUES( ?,? )")

	defer smtp.Close()

	if err != nil {
		fmt.Println("SaveReply Prepare Error",err)
		return cat,err
	}
	res,err := smtp.Exec(catId,title)
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

func (self * SessionDb) UpdateReply(cat NeuralRandomReply) (NeuralRandomReply,error){

	smtp,err := self.GetDb().Prepare("UPDATE neural_random_reply SET id_neural_cat = ?,title = ? WHERE id = ?")

	defer smtp.Close()

	if err != nil {
		fmt.Println("UpdateReply Prepare Error",err)
		return cat,err
	}
	_,err = smtp.Exec(cat.CatId,cat.Title,cat.Id)
	if err != nil {
		fmt.Println("UpdateReply Exec Error",err)
		return cat,err
	}
	return cat,nil
}

func (self * SessionDb) DeleteReply(cat NeuralRandomReply) bool{
	smtp,err := self.GetDb().Prepare("DELETE FROM neural_random_reply WHERE id = ?")
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