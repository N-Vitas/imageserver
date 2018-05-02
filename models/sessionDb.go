package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"unicode"
)

// Сессия с файловой базой данных
type SessionDb struct {
	db *sql.DB
}
// Переподключение или создание подключения к базе
func (self *SessionDb) GetDb() *sql.DB {
	if self.db != nil {
		return self.db
	}
	db, err := sql.Open("sqlite3", "./database/learnig.db")
	self.checkErr(err)
	self.db = db
	return self.db
}
// Поиск соответствия слов
func (self * SessionDb) FindAllId(maping []string) []float64{
	res := []float64{}
	for _,v := range maping{
		if v == ""{
			continue
		}

		res = append(res,self.findId(v))
	}
	return res
}

func (self * SessionDb) findId(search string) float64 {
	id := int64(0)
	err := self.GetDb().QueryRow("SELECT ID FROM sg_entry WHERE NAME LIKE '%"+search+"%'").Scan(&id)
	if err != nil {
		self.db.Close()
		self.db = nil
		return float64(self.createString(search))
	}
	return float64(id)
}
// Создание новой записи
func (self * SessionDb) createString(search string) int64 {
	if IsInt(search) {
		return 0
	}
	id := int64(0)
	self.GetDb().QueryRow("SELECT MAX(ID) FROM sg_entry").Scan(&id)
	smtp,err := self.GetDb().Prepare("INSERT INTO sg_entry( id, name, uname, id_class, freq, exportable, flags ) VALUES( ?, ?, 'НОВАЯ ЗАПИСЬ', 7, 0, 1, 0 )")
	self.checkErr(err)
	res,err := smtp.Exec(id+1,search)
	id,err = res.LastInsertId()
	self.checkErr(err)
	smtp.Close()
	return id
}
// Закрытие сессии
func (self * SessionDb) Close()  {
	if self.GetDb() != nil {
		self.GetDb().Close()
	}
}
// Обработка ошибки
func (self *SessionDb) checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func IsInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}