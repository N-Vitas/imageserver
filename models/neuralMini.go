package models

import (
	"fmt"
	"encoding/json"
	"github.com/fxsjy/gonn/gonn"
)

type LearningNeuralMini struct {
	Neural *gonn.NeuralNetwork
}
/* Создание нейросети
 * 50 входящих нейронов
 * 10 скрытых слоев
 * 1 нейрон результата классификации ответа
 */
func (self *LearningNeuralMini) Create(db *SessionDb,limit int)  {
	if !self.LoadNeuralDump(db,limit) {
		fmt.Println("Create NewNetwork")
		//self.Neural = neural.NewNetwork(255,[]int{500,255,100,10,1})
		//self.Neural = neural.NewNetwork(50,[]int{100,50,25,1})
		self.Neural = gonn.NewNetwork(limit,10,1,true,0.1,0.01)// neural.NewNetwork(10,[]int{10,3,1})
		//self.Neural.RandomizeSynapses()
		//self.dump = persist.ToDump(self.Neural)

		self.SaveNeuralDump(db,limit)
		return
	}
	fmt.Println("Loaded NewNetwork")
}
// Создание новой записи
func (self *LearningNeuralMini) SaveNeuralDump(db *SessionDb,limit int) {
	if db == nil{
		return
	}
	id := int64(0)
	//Enters":8,"Weights
	err := db.GetDb().QueryRow("SELECT Enters FROM learning WHERE Enters = ?",limit).Scan(&id)
	if err != nil {
		smtp,err := db.GetDb().Prepare("INSERT INTO learning( Enters, Weights) VALUES( ?,?)")
		db.checkErr(err)
		blop,err := json.Marshal(self.Neural)
		db.checkErr(err)
		res,err := smtp.Exec(limit,string(blop))
		db.checkErr(err)
		id,err = res.LastInsertId()
		db.checkErr(err)
		smtp.Close()
		fmt.Println("NewNetwork success create!")
	}
	if int(id) == limit{
		smtp,err := db.GetDb().Prepare("UPDATE learning SET Enters=?, Weights=? WHERE Enters = ?")
		db.checkErr(err)
		blop,err := json.Marshal(self.Neural)
		db.checkErr(err)
		res,err := smtp.Exec(limit,string(blop),limit)
		db.checkErr(err)
		id,err = res.LastInsertId()
		db.checkErr(err)
		smtp.Close()
		fmt.Println("NewNetwork success update!")
	}
	return
}
// Создание новой записи
func (self *LearningNeuralMini) LoadNeuralDump(db *SessionDb,limit int) bool {
	if db == nil {
		return false
	}
	weights := ""
	//Enters":8,"Weights
	err := db.GetDb().QueryRow("SELECT Weights FROM learning WHERE Enters = ?",limit).Scan(&weights)
	if err != nil {
		fmt.Println("LoadNeuralDump Error",err)
		return false
	}
	err = json.Unmarshal([]byte(weights),&self.Neural)
	if err != nil {
		fmt.Println("Unmarshal Error",err)
		return false
	}
	fmt.Println("LoadNeuralDump success",self.Neural.ErrOutput)
	return true
}
// Запрос результата от нейронной сети
func (self *LearningNeuralMini) Execute(in []float64) []float64 {
	return self.Neural.Forward(in)
}
// Запрос погрешности нейронной сети
func (self *LearningNeuralMini) HasError() []float64 {
	return self.Neural.ErrOutput
}
// Обучение нейронной сети
func (self *LearningNeuralMini) Learning(inputs [][]float64,out [][]float64,iteration int) {
	if iteration == 0 {
		iteration = 10000
	}
	self.Neural.Train(inputs,out,iteration)
}