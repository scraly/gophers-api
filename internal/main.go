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

	api.GetGopherHandler = operations.GetGopherHandlerFunc(GetGopherByName)

	api.PostGopherHandler = operations.PostGopherHandlerFunc(CreateGopher)

	api.DeleteGopherHandler = operations.DeleteGopherHandlerFunc(DeleteGopher)

	api.PutGopherHandler = operations.PutGopherHandlerFunc(UpdateGopher)

	// Start server which listening
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

type gopher struct {
	Name        string `json:"name"`
	Displayname string `json:"displayname"`
	URL         string `json:"url"`
}

type allGophers []gopher

/*var gophers = allGophers{
	{
		Name:        "5th-element",
		Displayname: "5th Element",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/5th-element.png",
	},
}*/

var gophers = allGophers{
	{
		Name:        "5th-element",
		Displayname: "5th Element",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/5th-element.png",
	},
	{
		Name:        "arrow-gopher",
		Displayname: "Arrow Gopher",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/arrow-gopher.png",
	},
	{
		Name:        "back-to-the-future-v2",
		Displayname: "Back to the future",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/back-to-the-future-v2.png",
	},
	{
		Name:        "baywatch",
		Displayname: "Baywatch",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/baywatch.jpg",
	},
	{
		Name:        "bike-gopher",
		Displayname: "Bike Gopher",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/bike-gopher.png",
	},
	{
		Name:        "blues-gophers",
		Displayname: "Blues Gopher",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/blues-gophers.png",
	},
	{
		Name:        "buffy-the-gopher-slayer",
		Displayname: "Buffy the vampire slayer",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/buffy-the-gopher-slayer.png",
	},
	{
		Name:        "chandleur-gopher",
		Displayname: "Chandleur Gopher",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/chandleur-gopher.png",
	},
	{
		Name:        "cherry-gopher",
		Displayname: "Cherry gopher",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/cherry-gopher.png",
	},
	{
		Name:        "cloud-nord",
		Displayname: "Cloud Nord",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/cloud-nord.png",
	},
	{
		Name:        "devnation-france-gopher",
		Displayname: "Devnation France",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/devnation-france-gopher.png",
	},
	{
		Name:        "dr-who-gophers",
		Displayname: "Dr Who",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/dr-who-gophers.png",
	},
	{
		Name:        "dumbledore",
		Displayname: "Dumbledore",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/dumbledore.png",
	},
	{
		Name:        "fire-gopher",
		Displayname: "Fire",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/fire-gopher.png",
	},
	{
		Name:        "firefly-gopher",
		Displayname: "Firefly",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/firefly-gopher.png",
	},
	{
		Name:        "fort-boyard",
		Displayname: "Fort Boyard",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/fort-boyard.png",
	},
	{
		Name:        "friends",
		Displayname: "Friends",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/friends.png",
	},
	{
		Name:        "gandalf",
		Displayname: "Gandalf",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/gandalf",
	},
	{
		Name:        "gladiator-gopher",
		Displayname: "Gladiator",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/gladiator-gopher.png",
	},
	{
		Name:        "gopher-dead",
		Displayname: "Gopher Dead",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/gopher-dead.png",
	},
	{
		Name:        "gopher johnny",
		Displayname: "Gopher Johnny",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/gopher-johnny",
	},
	{
		Name:        "gopher-open",
		Displayname: "Gopher Open",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/gopher-open.png",
	},
	{
		Name:        "gopher-speaker",
		Displayname: "Gopher speaker",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/gopher-speaker.png",
	},
	{
		Name:        "gopher",
		Displayname: "Gopher",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/gopher.png",
	},
	{
		Name:        "graffiti-devfest-nantes-2021",
		Displayname: "DevFest Nantes 2021",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/graffiti-devfest-nantes-2021.png",
	},
	{
		Name:        "halloween-spider",
		Displayname: "halloween spider.png",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/halloween-spider.png",
	},
	{
		Name:        "happy-gopher",
		Displayname: "Happy Gopher",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/happy-gopher.png",
	},
	{
		Name:        "harry-gopher",
		Displayname: "Harry Gopher",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/harry-gopher.png",
	},
	{
		Name:        "idea-gopher",
		Displayname: "Idea Gopher",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/idea-gopher.png",
	},
	{
		Name:        "indiana-jones",
		Displayname: "indiana jones.png",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/indiana-jones.png",
	},
	{
		Name:        "jedi-gopher",
		Displayname: "Jedi Gopher",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/jedi-gopher.png",
	},
	{
		Name:        "jurassic-park",
		Displayname: "jurassic park",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/jurassic-park.png",
	},
	{
		Name:        "love-gopher",
		Displayname: "Love Gopher",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/love-gopher.png",
	},
	{
		Name:        "luigi-gopher",
		Displayname: "Luigi Gopher",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/luigi-gopher.png",
	},
	{
		Name:        "mac-gopher",
		Displayname: "Mac Gopher",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/mac-gopher.png",
	},

	{
		Name:        "mario-gopher",
		Displayname: "Mario Gopher",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/mario-gopher.png",
	},
	{
		Name:        "marshal",
		Displayname: "Marshal Gopher",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/marshal.png",
	},
	{
		Name:        "men-in-black-v2",
		Displayname: "Men in black",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/men-in-black-v2.png",
	},
	{
		Name:        "mojito-gopher",
		Displayname: "Mojito Gopher",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/mojito-gopher.png",
	},
	{
		Name:        "paris-gopher",
		Displayname: "Paris Gopher",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/paris-gopher.png",
	},
	{
		Name:        "pere-fouras.png",
		Displayname: "Pere Fouras",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/pere-fouras.png",
	},
	{
		Name:        "recipe",
		Displayname: "Recipe",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/recipe.png",
	},
	{
		Name:        "sandcastle-gopher",
		Displayname: "Sandcastle",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/sandcastle-gopher.png",
	},
	{
		Name:        "santa-gopher",
		Displayname: "Santa",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/santa-gopher.png",
	},
	{
		Name:        "saved-by-the-bell",
		Displayname: "saved by the bell",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/saved-by-the-bell.png",
	},
	{
		Name:        "sheldon",
		Displayname: "Sheldon",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/sheldon.png",
	},
	{
		Name:        "star-wars",
		Displayname: "Star Wars",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/star-wars.png",
	},
	{
		Name:        "stargate",
		Displayname: "Stargate",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/stargate.png",
	},
	{
		Name:        "tadx-gopher",
		Displayname: "Tadx Gopher",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/tadx-gopher.png",
	},
	{
		Name:        "unicorn",
		Displayname: "Unicorn",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/unicorn.png",
	},
	{
		Name:        "urgences",
		Displayname: "Urgences",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/urgences.png",
	},
	{
		Name:        "vampire-xmas",
		Displayname: "vampire Xmas",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/vampire-xmas.png",
	},
	{
		Name:        "wired-gopher",
		Displayname: "Wired Gopher",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/wired-gopher.png",
	},
	{
		Name:        "x-files",
		Displayname: "X Files",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/x-files.png",
	},
	{
		Name:        "yoda-gopher",
		Displayname: "Yoda Gopher",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/yoda-gopher.png",
	},
	{
		Name:        "zelda-gopher",
		Displayname: "Zelda Gopher",
		URL:         "https://raw.githubusercontent.com/scraly/gophers/main/zelda-gopher.png",
	},
}

// Health route returns OK
func Health(operations.CheckHealthParams) middleware.Responder {
	fmt.Println("[Health] Call method")
	return operations.NewCheckHealthOK().
		WithPayload("OK").
		WithAccessControlAllowOrigin("*")
}

// Returns a a list of Gophers
func GetGophers(gopher operations.GetGophersParams) middleware.Responder {
	fmt.Println("[GetGophers] Call method")

	var gophersList []*models.Gopher

	// Get all existing Gophers
	for _, myGopher := range gophers {
		gophersList = append(gophersList, &models.Gopher{Name: myGopher.Name, Displayname: myGopher.Displayname, URL: myGopher.URL})
	}

	return operations.NewGetGophersOK().
		WithPayload(gophersList).
		WithAccessControlAllowOrigin("*")
}

// Returns an object of type Gopher with a given name
func GetGopherByName(gopherParam operations.GetGopherParams) middleware.Responder {
	fmt.Println("[GetGopherByName] Call method")

	for _, myGopher := range gophers {
		if myGopher.Name == gopherParam.Name {
			fmt.Println("Gopher", gopherParam.Name, "found in DB")

			return operations.NewGetGopherOK().WithPayload(
				&models.Gopher{
					Name:        myGopher.Name,
					Displayname: myGopher.Displayname,
					URL:         myGopher.URL}).
				WithAccessControlAllowOrigin("*")
		}
	}

	//If gopher have not been found, returns a 404 HTTP Error Code
	return operations.
		NewGetGopherNotFound().
		WithAccessControlAllowOrigin("*")
}

// TODO: to finish
func getGopher(gopherName string) gopher {
	for _, myGopher := range gophers {
		if myGopher.Name == gopherName {
			return myGopher
		}
	}

	return gopher{}
}

func gopherExists(gopherName string) bool {
	for _, myGopher := range gophers {
		if myGopher.Name == gopherName {
			return true
		}
	}

	return false
}

// Add a new Gopher
func CreateGopher(gopherParam operations.PostGopherParams) middleware.Responder {
	fmt.Println("[CreateGopher] Call method")

	name := gopherParam.Gopher.Name
	displayname := gopherParam.Gopher.Displayname
	url := gopherParam.Gopher.URL

	fmt.Println("Try to create a Gopher with the parameters:", *name, *displayname, *url)

	// Check if a gopher not already exists
	if !gopherExists(*name) {
		// Add new gopher in the list of existing Gophers
		gophers = append(gophers, gopher{*name, *displayname, *url})

		fmt.Println("Gopher", *name, "created!")

		return operations.NewPostGopherCreated().WithPayload(&models.Gopher{Name: *name, Displayname: *displayname, URL: *url})
	} else {
		return operations.NewPostGopherConflict()
	}
}

// Delete a Gopher with a given name
func DeleteGopher(gopherParam operations.DeleteGopherParams) middleware.Responder {
	fmt.Println("[DeleteGopher] Call method")

	for i, myGopher := range gophers {
		if myGopher.Name == gopherParam.Name {
			fmt.Println("Gopher", gopherParam.Name, "found in DB, try to delete it")

			gophers = append(gophers[:i], gophers[i+1:]...)

			fmt.Println("Gopher", gopherParam.Name, "deleted!")

			return operations.NewDeleteGopherOK()
		}
	}

	fmt.Println("[DeleteGopher] End of the method")

	//If gopher have not been found, returns a 404 HTTP Error Code
	return operations.NewDeleteGopherNotFound()
}

// Update the displayname and the URL of an existing Gopher
func UpdateGopher(gopherParam operations.PutGopherParams) middleware.Responder {
	fmt.Println("[UpdateGopher] Call method")

	fmt.Println("Updating", *gopherParam.Gopher.Name, "with new values")

	for i := range gophers {
		if gophers[i].Name == *gopherParam.Gopher.Name {
			gophers[i].Displayname = *gopherParam.Gopher.Displayname
			gophers[i].URL = *gopherParam.Gopher.URL

			fmt.Println("Gopher updated!")

			return operations.NewPostGopherCreated().WithPayload(&models.Gopher{
				Name:        *gopherParam.Gopher.Name,
				Displayname: *gopherParam.Gopher.Displayname,
				URL:         *gopherParam.Gopher.URL})
		}
	}

	fmt.Println("[UpdateGopher] End of the method")

	return operations.NewPutGopherOK()
}

//TODO: Create Helper function in order to create a JSON with full existing Gophers in github.com/scraly/gophers
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

// 		arr = append(arr, &models.Gopher{name, *c.Displayname, *c.DownloadURL})

// 	}

// 	return arr
// }
