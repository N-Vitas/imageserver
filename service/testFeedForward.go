package service

import (
	"github.com/emicklei/go-restful"
	"fmt"
)

func (self *App) testFeedForward(req *restful.Request, resp *restful.Response) {
	answers,_ := self.Model.Session.GetDictionaryes()
	if !self.Training.Ready() {
		for _, answer := range answers {
			self.Model.TransformBinary(answer.Binary)
			self.Training.AddTrainingItem(self.Model.GetValue(),R(int(answer.Id)))
			fmt.Println(self.Model.GetValue())
		}
	}
	self.Training.StartMini(self.Model.Session,self.Model.GetLimit(),self.MiniNeural)

	self.Model.TransformBinary(answers[len(answers)-1].Title) // Трансформируем массив для нейросети
	resp.WriteEntity([]interface{}{answers[len(answers)-1].Title,self.MiniNeural.HasError(),self.MiniNeural.Execute(self.Model.GetValue())})
}

func Q(L []int64 ) []float64{
	tmp := make([]float64,10)
	for _,x := range(L){
		tmp[x] +=1.0
	}
	return tmp
}

func R(F int) []float64{
	return []float64{float64(F)}
}

func round(X float64) int {
	return int(X+0.5)
}