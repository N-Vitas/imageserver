package models

import (
	"regexp"
	"strings"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"strconv"
)

// Абстрактная фабрика трансформации вопроса (Abstract Factory)

// Тип AbstractFactory, описывает интерфейс Абстрактной Фабрики
// Для трансформации предложения в набор классификации чисел.
// Конкретные фабрики должны его реализовать.
type AbstractFactoryF interface {
	CreateMaping(volume string) AbstractSendF
	CreateDump(volume string) AbstractDumpF
	CreateExecute(volume map[string]int64) AbstractExecuteF
	CreateBinId(integer int64) string
}

// Тип AbstractWater, описывает интерфейс предложения.
type AbstractSendF interface {
	GetVolume() []string // Возможность получения масива предложения из слов
}

// Тип AbstractWater, описывает интерфейс предложения.
type AbstractDumpF interface {
	GetVolume() map[string]int64 // Возможность получения масива предложения из слов
}

// Тип AbstractBottle, описывает интерфейс трансформации предложениея в массив чисел.
type AbstractExecuteF interface {
	InteractSend(maping AbstractSendF)  // Трансфигуратор взаимодействует с набором слов
	GetExecuteVolume() []int64           // Возможность получения трансформированных слов в числа
	GetSendVolume() []string            // Возможность получения массива слов из предложения
}

// Тип TransformFactory, реализует фабрику по трансформированию предложения в массив чисел.
type TransformFactoryF struct {
}

func (self *TransformFactoryF) CreateBinId(integer int64) string{
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

func (self *TransformFactoryF) TransformBinary(binary string) []int64{
	arbin := strings.Split(binary, "")
	newar := []int64{}
	for _,v := range arbin {
		n,e := strconv.ParseInt(v,10,64)
		if e != nil {
			newar = append(newar, 0)
		}
		newar = append(newar, n)
	}
	return newar
}

// Создает массив строк вырезая знаки припенания
func (self *TransformFactoryF) CreateMaping(volume string) *TransformTextF {
	pt, _ := regexp.Compile("[!/\".,!@#$%^&*()_=<>?-]")
	actual := string(pt.ReplaceAll([]byte(strings.ToLower(volume)), []byte("")))
	return &TransformTextF{volume: strings.Fields(actual)}
}

// Создает массив соотношений чисел к словам
func (self *TransformFactoryF) CreateExecute(volume map[string]int64) *TransformIntegerF {
	return &TransformIntegerF{volume: volume}
}
// Создается экземпляр дампа или обновляется
func (self *TransformFactoryF) CreateDump(path string,data []string) *TransformDumpF{
	return self.dumpFromFile(path)
}
// Обработанный массив строк
type TransformTextF struct {
	volume []string
}

// Получить созданный массив слов
func (self *TransformTextF) GetTransformValue() []string {
	return self.volume
}

// Реализации фабрики трансформации
type TransformIntegerF struct {
	maping  *TransformTextF // Трансформер должен иметь массив слов
	volume map[string]int64        // Массив соотношений слов к чиселу
	result []int64 // Трансформированный массив чисел
}

// Обнуляем результат. Затем трансформируем полученный массив в число и все записываем
func (self *TransformIntegerF) InteractSend(maping *TransformTextF) {
	self.maping = maping
	for _,value := range self.maping.volume{
		if self.volume[value] != 0{
			self.result = append(self.result,self.volume[value])
		}else{
			self.result = append(self.result,0)
		}
	}
}

// Получить массив чисел
func (self *TransformIntegerF) GetExecuteVolume() []int64 {
	return self.result
}

// Получить массив слов
func (self *TransformIntegerF) GetSendVolume() []string {
	return self.maping.GetTransformValue()
}

// Получить массив соответствия слов к числу
func (self *TransformDumpF) GetDumpVolume() map[string]int64 {
	return self.volume
}


// Обработанный массив строк
type TransformDumpF struct {
	path string
	volume map[string]int64
}
// Получаем дамп из файла
func (self *TransformFactoryF) dumpFromFile(path string) *TransformDumpF {
	b, err := ioutil.ReadFile(path)
	dump := &TransformDumpF{}
	dump.path = path
	if err != nil {
		fmt.Println("Пустой дамп!")
		return dump
	}

	err = json.Unmarshal(b, &dump.volume)
	if err != nil {
		panic(err)
	}
	return dump
}
// Обновляем дамп файла
func (self *TransformDumpF) DumpToFile(maping []string) {
	id := int64(0)
	for _,i := range self.GetDumpVolume() {
		if i > id {
			id = i
		}
	}
	dump := make(map[string]int64)
	for _,v := range maping{
		if self.volume[v] != 0 {
			dump[v] = self.volume[v]
			continue
		}
		id++
		dump[v] = int64(id)
	}
	self.volume = dump
	j, _ := json.Marshal(self.volume)

	err := ioutil.WriteFile(self.path, j, 0644)
	if err != nil {
		panic(err)
	}
}

// Создает массив строк вырезая знаки припенания только уже у существующей фабрики
func (self *TransformIntegerF) Execute(volume string) {
	pt, _ := regexp.Compile("[!/\".,!@#$%^&*()_=<>?-]")
	// Обновляем все кроме дампа
	self.maping.volume = strings.Fields(string(pt.ReplaceAll([]byte(strings.ToLower(volume)), []byte(""))))
	self.result = []int64{}
	for _,value := range self.maping.volume{
		if self.volume[value] != 0{
			self.result = append(self.result,self.volume[value])
		}else{
			self.result = append(self.result,0)
		}
	}
}