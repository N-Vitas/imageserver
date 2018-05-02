package models

import (
	"database/sql"
	"fmt"
)

type NeuralDump struct {
	Id int64 `json:"id"`
	CatId int64 `json:"catId"`
	Title string `json:"title"`
}


// Получение всех категории ответа нейронной сети
func (self * SessionDb) GetTrainings() ([]NeuralDump,error) {
	var (
		cat []NeuralDump
		id sql.NullInt64
		id_neural_cat sql.NullInt64
		title sql.NullString
	)
	rows,err := self.GetDb().Query("SELECT id,id_neural_cat,title FROM neural_learning_dump")
	if err != nil {
		fmt.Println("GetTrainings Query Error",err)
		return cat,err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id,&id_neural_cat,&title)
		if err != nil {
			fmt.Println("GetTrainings Scan Error",err)
			return cat,err
		}
		cat = append(cat,NeuralDump{ Id : id.Int64, CatId : id_neural_cat.Int64, Title : title.String })
	}
	return cat,nil
}
// Получение категории ответа нейронной сети
func (self * SessionDb) GetTraining(idcat int64) (NeuralDump,error) {
	var (
		cat NeuralDump
		id sql.NullInt64
		id_neural_cat sql.NullInt64
		title sql.NullString
	)
	err := self.GetDb().QueryRow("SELECT id,id_neural_cat,title FROM neural_learning_dump WHERE id = ? ",idcat).Scan(&id,&id_neural_cat,&title)
	if err != nil {
		fmt.Println("GetTraining Error",err)
		return cat,err
	}
	cat.Id = id.Int64
	cat.CatId = id_neural_cat.Int64
	cat.Title = title.String
	return cat,nil
}
// Получение категории ответа нейронной сети
func (self * SessionDb) SaveTraining(catId int64,title string) (NeuralDump,error) {
	cat := NeuralDump{}
	cat.Title =title
	smtp,err := self.GetDb().Prepare("INSERT INTO neural_learning_dump( id_neural_cat, title ) VALUES( ?,? )")

	defer smtp.Close()

	if err != nil {
		fmt.Println("SaveTraining Prepare Error",err)
		return cat,err
	}
	res,err := smtp.Exec(catId,title)
	if err != nil {
		fmt.Println("SaveTraining Exec Error",err)
		return cat,err
	}
	cat.Id,err = res.LastInsertId()
	if err != nil {
		fmt.Println("SaveTraining LastInsertId Error",err)
		return cat,err
	}
	return cat,nil
}

func (self * SessionDb) UpdateTraining(cat NeuralDump) (NeuralDump,error){

	smtp,err := self.GetDb().Prepare("UPDATE neural_learning_dump SET id_neural_cat = ?,title = ? WHERE id = ?")

	defer smtp.Close()

	if err != nil {
		fmt.Println("UpdateTraining Prepare Error",err)
		return cat,err
	}
	_,err = smtp.Exec(cat.CatId,cat.Title,cat.Id)
	if err != nil {
		fmt.Println("UpdateTraining Exec Error",err)
		return cat,err
	}
	return cat,nil
}

func (self * SessionDb) DeleteTraining(cat NeuralDump) bool{
	smtp,err := self.GetDb().Prepare("DELETE FROM neural_learning_dump WHERE id = ?")
	defer smtp.Close()
	if err != nil {
		fmt.Println("DeleteTraining Prepare Error",err)
		return false
	}
	_,err = smtp.Exec(cat.Id)
	if err != nil {
		fmt.Println("DeleteTraining Exec Error",err)
		return false
	}
	return true
}