package main

import (
	. "learnig-assistain/models"
	. "learnig-assistain/service"
	"fmt"
	"io/ioutil"
	"strings"
)

func main()  {
	app := App{Model:&TransformFactory{Session:&SessionDb{}},Neural:&LearningNeural{},Training:&Training{}}
	fmt.Println(app)
	dat, err := ioutil.ReadFile("answer_databse.bin")
	if err != nil {
		fmt.Println(err)
	}
	str := string(dat)
	ar := strings.SplitAfter(str, "\\0")
	var traning string
	var reply string
	if app.Model.Session.DeleteAnswerAll() {
		fmt.Println("База очищена")
	}
	for _,v := range ar {
		actual := strings.SplitAfter(v, "\\")
		traning = strings.TrimSpace(strings.Replace(resolve(actual,0),"\\","",-1))
		reply = strings.TrimSpace(strings.Replace(resolve(actual,1),"\\","",-1))
		fmt.Println("Вопрос",traning)
		fmt.Println("Ответ",reply)
		if len(traning) > 0 && len(reply) > 0 {
			app.Model.Session.SaveAnswer(traning,reply)
		}
	}
	//for _,v := range traning{
	//	app.Model.Session.SaveTraining(1,v)
	//}
	//for _,v := range reply{
	//	app.Model.Session.SaveReply(1,v)
	//}
}

func resolve(str []string,i int)string  {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	return str[i]
}
