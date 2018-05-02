package models

import (
	"regexp"
	"strings"
	"strconv"
)

// Абстрактная фабрика трансформации вопроса (Abstract Factory)

// Тип AbstractFactory, описывает интерфейс Абстрактной Фабрики
// Для трансформации предложения в набор классификации чисел.
// Конкретные фабрики должны его реализовать.

// Тип AbstractWater, описывает интерфейс предложения.
type AbstractFactory interface {
	CreateMaping(volume string)
	SetLimit(limit int)
	GetVolume() []string // Возможность получения масива предложения из слов
	GetValue() []float64 // Возможность получения масива предложения из чисел
	CreateBinId(integer int64) string
	TransformBinary(binary string)
}

// Тип TransformFactory, реализует фабрику по трансформированию предложения в массив чисел.
type TransformFactory struct {
	Session *SessionDb
	volume []string
	value []float64
	limit int
}

// Создает массив строк вырезая знаки припенания
func (self *TransformFactory) CreateMaping(volume string) {
	pt, _ := regexp.Compile("[!/\".,!@#$%^&*()_=<>?-]")
	actual := string(pt.ReplaceAll([]byte(strings.ToLower(volume)), []byte("")))
	self.volume = strings.Fields(actual)
}
// Получить созданный массив слов
func (self *TransformFactory) GetVolume() []string {
	return self.volume
}

func (self *TransformFactory) CreateBinId(integer int64) string{
	bin := strconv.FormatInt(integer, 2)
	arbin := strings.Split(bin, "")
	newar := []string{}
	for i := 0; i < 16 - len(arbin); i++ {
		newar = append(newar,"0")
	}
	for _,v := range arbin {
		newar = append(newar, v)
	}
	return strings.Join(newar,"")
}

func (self *TransformFactory) TransformBinary(binary string){
	arbin := strings.Split(binary, "")
	self.value = []float64{}
	for _,v := range arbin {
		n,e := strconv.ParseFloat(v,64)
		if e != nil {
			self.value = append(self.value, 0)
		}
		self.value = append(self.value, n)
		self.volume = append(self.volume, v)
	}
}
// Получить созданный массив чисел
func (self *TransformFactory) GetValue() []float64 {
	return self.value
}

// Поиск и трансформация данных предложения
func (self *TransformFactory) TransformExecute() {
	if self.Session == nil {
		self.Session = &SessionDb{}
	}
	self.value = self.Session.FindAllId(self.volume)
	if self.limit != 0{
		i := len(self.value)
		for i < self.limit {
			i++
			self.value = append(self.value,0)
		}
	}
}

func (self *TransformFactory) SetLimit(limit int)  {
	self.limit = limit
}

func (self *TransformFactory) GetLimit() int {
	return self.limit
}
