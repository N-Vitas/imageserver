package service

import (
	. "imageserver/models"
	"github.com/emicklei/go-restful"
	"strconv"
	"encoding/json"
	"io/ioutil"
	"net/http"
)
// Запросы категорий
func (self *App) routNeuralDump(req *restful.Request, resp *restful.Response) {
	id, err := strconv.ParseInt(req.PathParameter("id"), 10, 64)
	if err != nil {
		cat,err := self.Model.Session.GetTrainings()
		if err != nil {
			resp.WriteEntity(map[string]interface{}{"Error":err})
			return
		}
		resp.WriteEntity(cat)
		return
	}

	cat,err := self.Model.Session.GetTraining(id)
	if err != nil {
		resp.WriteEntity(map[string]interface{}{"Error":err})
		return
	}
	resp.WriteEntity(cat)
}
// Создание категории
func (self *App) createNeuralDump(req *restful.Request, resp *restful.Response) {
	formData := struct {
		Title string `json:"title"`
		CatId int64 `json:"catId"`
	}{}
	body, err := ioutil.ReadAll(req.Request.Body)
	if err != nil {
		resp.WriteHeaderAndEntity(http.StatusBadRequest,err)
		return
	}
	err = json.Unmarshal(body, &formData)
	if err != nil {
		resp.WriteHeaderAndEntity(http.StatusBadRequest,err)
		return
	}

	cat,err := self.Model.Session.SaveTraining(formData.CatId,formData.Title)
	if err != nil {
		resp.WriteEntity(map[string]interface{}{"Error":err})
		return
	}
	resp.WriteEntity(cat)
}
// Обновление категории
func (self *App) updateNeuralDump(req *restful.Request, resp *restful.Response) {
	formData := NeuralDump{}
	body, err := ioutil.ReadAll(req.Request.Body)
	if err != nil {
		resp.WriteHeaderAndEntity(http.StatusBadRequest,err)
		return
	}
	err = json.Unmarshal(body, &formData)
	if err != nil {
		resp.WriteHeaderAndEntity(http.StatusBadRequest,err)
		return
	}

	cat,err := self.Model.Session.UpdateTraining(formData)
	if err != nil {
		resp.WriteEntity(map[string]interface{}{"Error":err})
		return
	}
	resp.WriteEntity(cat)
}
// Удаление категории
func (self *App) delNeuralDump(req *restful.Request, resp *restful.Response) {
	formData := NeuralDump{}
	body, err := ioutil.ReadAll(req.Request.Body)
	if err != nil {
		resp.WriteHeaderAndEntity(http.StatusBadRequest,err)
		return
	}
	err = json.Unmarshal(body, &formData)
	if err != nil {
		resp.WriteHeaderAndEntity(http.StatusBadRequest,err)
		return
	}
	resp.WriteEntity(map[string]interface{}{"succes":self.Model.Session.DeleteTraining(formData)})
}
