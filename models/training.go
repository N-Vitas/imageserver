package models

import (
	"time"
	"fmt"
)

/*
 * Модель тренировки для нейронной сети
 */
const speed  = 0.1

type AbstractTraining interface {
	GetIn() []float64
	GetOut() float64
	GetError(index int) float64
	GetErrors() []float64
	Clear()
	AddError(err float64)
	SetSpeed(speed float64)
	AddTrainingItem(in []float64,out float64)
	SetTraining(in [][]float64,out []float64)
	Start(db *SessionDb,limit int,n *LearningNeural)
	StartMini(db *SessionDb,limit int,n *LearningNeuralMini)
}
type Training struct {
	in []float64
	out float64
	error []float64
	speed float64
	traningIn [][]float64
	traningOut [][]float64
}
// Получение входящих данных для тренировки
func (self *Training) Ready() bool {
	return len(self.traningOut) > 0
}
// Получение входящих данных для тренировки
func (self *Training) GetIn() []float64 {
	return self.in
}
// Получение исходящих для тренировки
func (self *Training) GetOut() float64 {
	return self.out
}
// Установка входящих и ожидающих данных для тренировки
func (self *Training) AddTrainingItem(in []float64,out []float64) {
	self.traningIn = append(self.traningIn,in)
	self.traningOut = append(self.traningOut,out)
	self.in = in
	self.out = out[0]
}
// Установка исходящих для тренировки
func (self *Training) SetTraining(in [][]float64,out [][]float64) {
	cIn := len(in)
	cOut := len(out)
	if cIn == 0 || cOut == 0{
		return
	}
	for k,v := range in {
		if len(v) == 255 && k <= cOut{
			self.traningOut = append(self.traningOut,out[k])
			self.traningIn = append(self.traningIn,v)
		}
	}
	self.in = in[0]
	self.out = out[0][0]
}
// Установка скорости обучения нейросети
func (self *Training) SetSpeed(newspeed float64) {
	self.speed = newspeed
}
// Добавление погрешности в массив чтоб можно было выстроить график
func (self *Training) AddError(err float64) {
	self.error = append(self.error,err)
}
// Общий сброс тренировки
func (self *Training) Clear() {
	self.in = []float64{}
	self.out = 0
	self.error = []float64{}
	self.traningIn = [][]float64{}
	self.traningOut = [][]float64{}
}
// Получение массива погрешности нейросети
func (self *Training) GetError(index int) float64 {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	return self.error[index]
}
// Получение массива погрешности нейросети
func (self *Training) GetErrors() []float64 {
	return self.error
}
// Тренировка нейросети
func (self *Training) Start(db *SessionDb,limit int,n *LearningNeural){
	clock := time.Now().UnixNano() // Текущее время для показа времени прохождения тренировки
	// Начало тренировки
	fmt.Println("Старт тренировки")
	for i := 0; i < 10000;i++{
		n.Learning(self.traningIn,self.traningOut,1000)
		n.SaveNeuralDump(db,limit)
	}
	fmt.Printf("Тренировка завершена %d наносек.\n",time.Now().UnixNano()-clock)
}
// Тренировка нейросети
func (self *Training) StartMini(db *SessionDb,limit int,n *LearningNeuralMini){
	clock := time.Now().UnixNano() // Текущее время для показа времени прохождения тренировки
	// Начало тренировки
	fmt.Println("Старт тренировки")
	for i := 0; i < 10000;i++{
		n.Learning(self.traningIn,self.traningOut,1000)
		n.SaveNeuralDump(db,limit)
	}
	fmt.Printf("Тренировка завершена %d наносек.\n",time.Now().UnixNano()-clock)
}