package main

import (
	. "imageserver/models"
	. "imageserver/service"
	"github.com/emicklei/go-restful"
	"net/http"
	"log"
	"github.com/bmizerany/pat"
	"fmt"
)


func main() {
	app := App{Model:&TransformFactory{Session:&SessionDb{}},Neural:&LearningNeural{},MiniNeural:&LearningNeuralMini{},Training:&Training{}}
	app.Model.SetLimit(16)
	app.Neural.Create(app.Model.Session,50)
	app.MiniNeural.Create(app.Model.Session,16)

	//app.FeedForward.Init(2,2,1)
	// accept and respond in JSON unless told otherwise
	restful.DefaultRequestContentType(restful.MIME_JSON)
	restful.DefaultResponseContentType(restful.MIME_JSON)
	// gzip if accepted
	restful.DefaultContainer.EnableContentEncoding(true)
	// faster router
	restful.DefaultContainer.Router(restful.CurlyRouter{})
	restful.Filter(app.CorsRule(restful.DefaultContainer).Filter)
	restful.Filter(app.OptionFilter)

	restful.Add(app.RegisterRoute())

	// Для отдачи сервером статичных файлов из папки public/static
	fs := http.FileServer(http.Dir("./public/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	mux := pat.New()
	mux.Get("/:page", http.HandlerFunc(app.PostHandler))
	mux.Get("/:page/", http.HandlerFunc(app.PostHandler))
	mux.Get("/", http.HandlerFunc(app.PostHandler))
	http.Handle("/", mux)

	fmt.Println("[go-restful] Сервис нейросети бота. Работает на http://localhost:4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
