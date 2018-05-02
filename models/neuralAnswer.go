package models

import (
	"fmt"
	"database/sql"
)

type NeuralAnswer struct {
	Id int64 `json:"id"`
	Question string `json:"question"`
	Answer string `json:"answer"`
}


// Получение всех категории ответа нейронной сети
func (self * SessionDb) GetAnswers() ([]NeuralAnswer,error) {
	var (
		cat []NeuralAnswer
		id sql.NullInt64
		question sql.NullString
		answer sql.NullString
	)
	rows,err := self.GetDb().Query("SELECT id,question,answer FROM neural_answer")
	if err != nil {
		fmt.Println("GetReplys Query Error",err)
		return cat,err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id,&question,&answer)
		if err != nil {
			fmt.Println("GetReplys Scan Error",err)
			return cat,err
		}
		cat = append(cat,NeuralAnswer{ Id : id.Int64, Question : question.String, Answer : answer.String })
	}
	return cat,nil
}
// Получение категории ответа нейронной сети
func (self * SessionDb) GetAnswer(idcat int64) (NeuralAnswer,error) {
	var (
		cat NeuralAnswer
		id sql.NullInt64
		question sql.NullString
		answer sql.NullString
	)
	err := self.GetDb().QueryRow("SELECT id,question,answer FROM neural_answer WHERE id = ? ",idcat).Scan(&id,&question,&answer)
	if err != nil {
		fmt.Println("GetReply Error",err)
		return cat,err
	}
	cat.Id = id.Int64
	cat.Question = question.String
	cat.Answer = answer.String
	return cat,nil
}
// Поиск ответа нейронной сети
func (self * SessionDb) FindAnswer(search string) (NeuralAnswer,error) {
	var (
		cat NeuralAnswer
		id sql.NullInt64
		question sql.NullString
		answer sql.NullString
	)
	cat.Id = 2
	err := self.GetDb().QueryRow("SELECT id,question,answer FROM neural_answer WHERE question LIKE '%"+search+"%'").Scan(&id,&question,&answer)
	if err != nil {
		fmt.Println("FindAnswer Error",err)
		return cat,err
	}
	cat.Id = id.Int64
	cat.Question = question.String
	cat.Answer = answer.String
	return cat,nil
}
// Поиск ответа нейронной сети
func (self * SessionDb) FindQuestion(search string) (NeuralAnswer,error) {
	var (
		cat NeuralAnswer
		id sql.NullInt64
		question sql.NullString
		answer sql.NullString
	)
	cat.Id = 2
	err := self.GetDb().QueryRow("SELECT id,question,answer FROM neural_answer WHERE answer LIKE '%"+search+"%'").Scan(&id,&question,&answer)
	if err != nil {
		fmt.Println("FindAnswer Error",err)
		return cat,err
	}
	cat.Id = id.Int64
	cat.Question = question.String
	cat.Answer = answer.String
	return cat,nil
}
// Получение категории ответа нейронной сети
func (self * SessionDb) SaveAnswer(question string,answer string) (NeuralAnswer,error) {
	cat := NeuralAnswer{}
	cat.Answer = answer
	smtp,err := self.GetDb().Prepare("INSERT INTO neural_answer( question, answer ) VALUES( ?,? )")

	defer smtp.Close()

	if err != nil {
		fmt.Println("SaveReply Prepare Error",err)
		return cat,err
	}
	res,err := smtp.Exec(question,answer)
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

func (self * SessionDb) UpdateAnswer(cat NeuralAnswer) (NeuralAnswer,error){

	smtp,err := self.GetDb().Prepare("UPDATE neural_answer SET question = ?,answer = ? WHERE id = ?")

	defer smtp.Close()

	if err != nil {
		fmt.Println("UpdateReply Prepare Error",err)
		return cat,err
	}
	_,err = smtp.Exec(cat.Question,cat.Answer,cat.Id)
	if err != nil {
		fmt.Println("UpdateReply Exec Error",err)
		return cat,err
	}
	return cat,nil
}

func (self * SessionDb) DeleteAnswer(cat NeuralAnswer) bool{
	smtp,err := self.GetDb().Prepare("DELETE FROM neural_answer WHERE id = ?")
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
func (self * SessionDb) DeleteAnswerAll() bool{
	smtp,err := self.GetDb().Prepare("DELETE FROM neural_answer")
	defer smtp.Close()
	if err != nil {
		fmt.Println("DropAnswer Prepare Error",err)
		return false
	}
	_,err = smtp.Exec()
	if err != nil {
		fmt.Println("DropAnswer Exec Error",err)
		return false
	}
	return true
}