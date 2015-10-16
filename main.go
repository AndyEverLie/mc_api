package main
import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/AndyEverLie/mc_api/routes"
	"log"
	"net/http"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := getRouter()
	if err != nil {
		log.Fatal(err)
	}

	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func getRouter() (rest.App, error) {
	mcPlugins := routes.McPlugins{}

	router, err := rest.MakeRouter(
		rest.Get("/plugins", mcPlugins.GetAllPlugins),
		rest.Post("/plugins", mcPlugins.PostPlugin),
		rest.Get("/plugins/:id", mcPlugins.GetPlugin),
		rest.Put("/plugins/:id", mcPlugins.PutPlugin),
		rest.Delete("/plugins/:id", mcPlugins.DeletePlugin),
	)
	return router, err
}