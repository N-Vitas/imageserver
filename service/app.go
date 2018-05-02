package service

import (
	. "imageserver/models"
	"github.com/emicklei/go-restful"
	"strings"
)
type App struct {
	Model *TransformFactory
	Neural *LearningNeural
	MiniNeural *LearningNeuralMini
	Training *Training
}

func (self *App) RegisterRoute() *restful.WebService {
	ws := new(restful.WebService)
	//search := ws.PathParameter("search", "search!").DataType("string")
	//id := ws.PathParameter("id", "indificator!").DataType("integer")
	ws.Path("/api")
	ws.Consumes("*/*")
	ws.Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/").To(self.defaultResponse))
	//ws.Route(ws.GET("/{search}").To(self.staticFromPathParam).Param(search))
	//ws.Route(ws.GET("/category").To(self.routCategory))
	//ws.Route(ws.GET("/category/{id}").To(self.routCategory).Param(id))
	//ws.Route(ws.POST("/category").To(self.createCategory))
	//ws.Route(ws.PUT("/category").To(self.updateCategory))
	//ws.Route(ws.DELETE("/category").To(self.delCategory))
	//ws.Route(ws.GET("/neuraldump").To(self.routNeuralDump))
	//ws.Route(ws.GET("/neuraldump/{id}").To(self.routNeuralDump).Param(id))
	//ws.Route(ws.POST("/neuraldump").To(self.createNeuralDump))
	//ws.Route(ws.PUT("/neuraldump").To(self.updateNeuralDump))
	//ws.Route(ws.DELETE("/neuraldump").To(self.delNeuralDump))
	//ws.Route(ws.GET("/neuralreply").To(self.routNeuralReply))
	//ws.Route(ws.GET("/neuralreply/{id}").To(self.routNeuralReply).Param(id))
	//ws.Route(ws.POST("/neuralreply").To(self.createNeuralReply))
	//ws.Route(ws.PUT("/neuralreply").To(self.updateNeuralReply))
	//ws.Route(ws.DELETE("/neuralreply").To(self.delNeuralReply))
	//ws.Route(ws.GET("/training").To(self.testFeedForward))
	return ws
}

func (self *App) defaultResponse(req *restful.Request, resp *restful.Response) {
	resp.WriteEntity(map[string]string{"server":"Файловый сервер изображений","version":"0.1"}) // Результат отдаем в ответе апи
}

func (self *App) staticFromPathParam(req *restful.Request, resp *restful.Response) {
	search := req.PathParameter("search") // Получаем вопрос
	//self.Model.CreateMaping(search) // Собираем масив вопросов
	//self.Model.TransformExecute() // Трансформируем массив для нейросети
	find,err := self.Model.Session.FindDictionary(search)
	if err != nil {
		resp.WriteEntity(err)
		return
	}
	self.Model.TransformBinary(find.Binary)

	result := self.MiniNeural.Execute(self.Model.GetValue()) // Отправляем запрос в нейросеть и получаем классификатор
	answer,err := self.Model.Session.GetDictionary(int64(result[0]))
	if err != nil {
		resp.WriteEntity(err)
		return
	}
	resp.WriteEntity(map[string]string{"search":find.Title,"answer":answer.Title}) // Результат отдаем в ответе апи
}
// Фильтр метода OPTION
func  (self *App) OptionFilter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	if "OPTIONS" != req.Request.Method && strings.Index(req.Request.URL.String(), "login") != -1 {
		chain.ProcessFilter(req, resp)
		return
	}
	chain.ProcessFilter(req, resp)
}

func (self * App) CorsRule( container *restful.Container,) restful.CrossOriginResourceSharing {
	corsRule := restful.CrossOriginResourceSharing{
		//ExposeHeaders: []string{"Content-Type"},
		AllowedDomains: []string{"http://test-react.app","http://localhost:* http://127.0.0.1:*","http://192.168.150.52"},
		AllowedHeaders: []string{"content-type", "Accept","X-Custom-Header", "Origin"},
		CookiesAllowed: true,
		Container: container,
	}
	return corsRule
}