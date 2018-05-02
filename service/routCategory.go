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
func (self *App) routCategory(req *restful.Request, resp *restful.Response) {
	id, err := strconv.ParseInt(req.PathParameter("id"), 10, 64)
	if err != nil {
		cat,err := self.Model.Session.GetCategories()
		if err != nil {
			resp.WriteEntity(map[string]interface{}{"Error":err})
			return
		}
		resp.WriteEntity(cat)
		return
	}

	cat,err := self.Model.Session.GetCategory(id)
	if err != nil {
		resp.WriteEntity(map[string]interface{}{"Error":err})
		return
	}
	resp.WriteEntity(cat)
}
// Создание категории
func (self *App) createCategory(req *restful.Request, resp *restful.Response) {
	formData := struct {Title string `json:"title"`}{}
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

	cat,err := self.Model.Session.SaveCategory(formData.Title)
	if err != nil {
		resp.WriteEntity(map[string]interface{}{"Error":err})
		return
	}
	resp.WriteEntity(cat)
}
// Обновление категории
func (self *App) updateCategory(req *restful.Request, resp *restful.Response) {
	formData := Category{}
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

	cat,err := self.Model.Session.UpdateCategory(formData)
	if err != nil {
		resp.WriteEntity(map[string]interface{}{"Error":err})
		return
	}
	resp.WriteEntity(cat)
}
// Удаление категории
func (self *App) delCategory(req *restful.Request, resp *restful.Response) {
	formData := Category{}
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
	resp.WriteEntity(map[string]interface{}{"succes":self.Model.Session.DeleteCategory(formData)})
}
