package main

import (
	"fmt"
	"log"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/scraly/gophers-api/pkg/swagger/server/models"
	"github.com/scraly/gophers-api/pkg/swagger/server/restapi"

	"github.com/scraly/gophers-api/pkg/swagger/server/restapi/operations"
)

func main() {

	// Initialize Swagger
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewGophersAPIAPI(swaggerSpec)
	server := restapi.NewServer(api)

	defer func() {
		if err := server.Shutdown(); err != nil {
			// error handle
			log.Fatalln(err)
		}
	}()

	server.Port = 8080

	api.CheckHealthHandler = operations.CheckHealthHandlerFunc(Health)

	api.GetGophersHandler = operations.GetGophersHandlerFunc(GetGophers)

	api.PostGopherHandler = operations.PostGopherHandlerFunc(PostGopher)

	// Start server which listening
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

type gopher struct {
	Name string `json:"Name"`
	Path string `json:"Path"`
	URL  string `json:"URL"`
}

type allGophers []gopher

var gophers = allGophers{
	{
		Name: "5th-element",
		Path: "5th-element.png",
		URL:  "https://raw.githubusercontent.com/scraly/gophers/main/5th-element.png",
	},
}

// Health route returns OK
func Health(operations.CheckHealthParams) middleware.Responder {
	return operations.NewCheckHealthOK().WithPayload("OK")
}

// Display Gopher list with optional filter
func GetGophers(gopher operations.GetGophersParams) middleware.Responder {

	var gophersList []*models.Gopher

	// get Gophers by name
	if gopher.Name != nil {

		for _, myGopher := range gophers {
			//TODO: faire un starting with?
			if myGopher.Name == *gopher.Name {
				fmt.Println("name", *gopher.Name, "name found in DB", myGopher.Name)
				gophersList = append(gophersList, &models.Gopher{Name: myGopher.Name, Path: myGopher.Path, URL: myGopher.URL})
			}
		}
	} else {
		// get all existing Gophers
		for _, myGopher := range gophers {
			gophersList = append(gophersList, &models.Gopher{Name: myGopher.Name, Path: myGopher.Path, URL: myGopher.URL})
		}
	}

	return operations.NewGetGophersOK().WithPayload(gophersList)
}

// Add new Gopher
func PostGopher(gopherParam operations.PostGopherParams) middleware.Responder {

	name := gopherParam.Gopher.Name
	path := gopherParam.Gopher.Path
	url := gopherParam.Gopher.URL

	// Add new gopher in the list of existing Gophers
	gophers = append(gophers, gopher{*name, *path, *url})

	return operations.NewPostGopherOK().WithPayload(&models.Gopher{Name: *name, Path: *path, URL: *url})
}

//Helper function?
// /*
// *
// Get Gophers List from Scraly repository
// */
// func GetGophersList() []*models.Gopher {

// 	client := github.NewClient(nil)

// 	// list public repositories for org "github"
// 	ctx := context.Background()
// 	// list all repositories for the authenticated user
// 	_, directoryContent, _, err := client.Repositories.GetContents(ctx, "scraly", "gophers", "/", nil)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	var arr []*models.Gopher

// 	for _, c := range directoryContent {
// 		if *c.Name == ".gitignore" || *c.Name == "README.md" {
// 			continue
// 		}

// 		var name string = strings.Split(*c.Name, ".")[0]

// 		arr = append(arr, &models.Gopher{name, *c.Path, *c.DownloadURL})

// 	}

// 	return arr
// }
