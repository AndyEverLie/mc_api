package routes
import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/AndyEverLie/mc_api/utils"
	"github.com/AndyEverLie/mc_api/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type McPlugin struct {
	Id     bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name   string `json:"name" bson:"name"`
	Detail string `json:"detail" bson:"detail"`
}
type McPlugins struct {
	Store map[string]*McPlugin
}

func (this *McPlugins) GetAllPlugins(w rest.ResponseWriter, req *rest.Request) {
	var plugins []McPlugin
	query := func(c *mgo.Collection) error {
		return c.Find(nil).All(&plugins)
	}
	if err := db.Query("plugins", query); err != nil {
		w.WriteJson(utils.Error(1, err.Error()))
	}

	w.WriteJson(utils.Success(&plugins))
}

func (this *McPlugins) GetPlugin(w rest.ResponseWriter, req *rest.Request) {
	mcPlugin := McPlugin{}
	id := req.PathParam("id")
	objId := bson.ObjectIdHex(id)

	query := func(c *mgo.Collection) error {
		return c.FindId(objId).One(&mcPlugin)
	}
	if err := db.Query("plugins", query); err != nil {
		w.WriteJson(utils.Error(404, err.Error()))
		return
	}
	w.WriteJson(utils.Success(&mcPlugin))
}

func (this *McPlugins) PostPlugin(w rest.ResponseWriter, req *rest.Request) {
	mcPlugin := McPlugin{}
	if err := req.DecodeJsonPayload(&mcPlugin); err != nil {
		w.WriteJson(utils.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	mcPlugin.Id = bson.NewObjectId()
	query := func(c *mgo.Collection) error {
		return c.Insert(mcPlugin)
	}
	if err := db.Query("plugins", query); err != nil {
		w.WriteJson(utils.Error(1, err.Error()))
		return
	}
	w.WriteJson(utils.Success(&mcPlugin))
}

func (this *McPlugins) PutPlugin(w rest.ResponseWriter, req *rest.Request) {
	mcPlugin := McPlugin{}
	if err := req.DecodeJsonPayload(&mcPlugin); err != nil {
		w.WriteJson(utils.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	id := req.PathParam("id")
	objId := bson.ObjectIdHex(id)
	src := bson.M{"_id": objId}
	tar := bson.M{
		"name": mcPlugin.Name,
		"detail": mcPlugin.Detail,
	}
	query := func(c *mgo.Collection) error {
		return c.Update(src, tar)
	}
	err := db.Query("plugins", query)
	if err != nil {
		w.WriteJson(utils.Error(1, err.Error()))
		return
	}

	w.WriteJson(utils.Success(&mcPlugin))
}

func (this *McPlugins) DeletePlugin(w rest.ResponseWriter, req *rest.Request) {
	id := req.PathParam("id")
	objId := bson.ObjectIdHex(id)

	query := func(c *mgo.Collection) error {
		return c.RemoveId(objId)
	}
	err := db.Query("plugins", query)
	if err != nil {
		w.WriteJson(utils.Error(1, err.Error()))
		return
	}

	w.WriteJson(utils.Success(objId))
}
