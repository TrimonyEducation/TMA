package main

import (
	"golang-crud/controllers"
	"golang-crud/utils"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

//TODO: QUERY PARAMS FOR ALL READER (SEARCH BY SUBJECT); ADD isPUBLISHED TO MODELS; isFinished




func init(){
	err:=godotenv.Load()
	if err!= nil {
        log.Println(err.Error())
	}
	utils.ConnectDB()
	
}

func setVideoRoutes(app *fiber.App){
	app.Post("/api/video", controllers.CreateVideo)
	app.Get("/api/video", controllers.GetAllVideos)
	app.Get("/api/video/:id", controllers.GetVideo)
	app.Patch("/api/video/:id", controllers.UpdateVideo)
	app.Delete("/api/video/:id", controllers.DeleteVideo)
}

func setExerciseRoutes(app *fiber.App){
	app.Post("/api/exercise", controllers.CreateExercise)
	app.Get("/api/exercise", controllers.GetAllExercises)
    app.Get("/api/exercise/:id", controllers.GetExercise)
    app.Patch("/api/exercise/:id", controllers.UpdateExercise)
    app.Delete("/api/exercise/:id", controllers.DeleteExercise)
}

func setProblemRoutes(app *fiber.App) {
	app.Post("/api/problem", controllers.CreateProblem)
    app.Get("/api/problem", controllers.GetAllProblems)
    app.Get("/api/problem/:id", controllers.GetProblem)
    app.Patch("/api/problem/:id", controllers.UpdateProblem)
    app.Delete("/api/problem/:id", controllers.DeleteProblem)
}

func setTakeRoutes(app *fiber.App){
	app.Post("/api/take", controllers.CreateTake)
    app.Get("/api/take", controllers.GetAllTakes)
	app.Get("/api/take/:id", controllers.GetTake)
    app.Patch("/api/take/:id", controllers.UpdateTake)
    app.Delete("/api/take/:id", controllers.DeleteTake)
}

func setClassRoutes(app *fiber.App) {
	app.Post("/api/class", controllers.CreateClass)
    app.Get("/api/class", controllers.GetAllClasses)
	app.Get("/api/class/:id", controllers.GetClass)
    app.Patch("/api/class/:id", controllers.UpdateClass)
    app.Delete("/api/class/:id", controllers.DeleteClass)
}

func setPlaylistRoutes(app *fiber.App) {
	app.Post("/api/playlist", controllers.CreatePlaylist)
    app.Get("/api/playlist", controllers.GetAllPlaylists)
    app.Get("/api/playlist/:id", controllers.GetPlaylist)
    app.Patch("/api/playlist/:id", controllers.UpdatePlaylist)
    app.Delete("/api/playlist/:id", controllers.DeletePlaylist)
}

func main() {
    app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	app.All("/api/" , func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "Welcome to the Trimony API 🥹",
		})
	})

	//USER ROUTES
	setVideoRoutes(app)

	//EXERCISE ROUTES
	setExerciseRoutes(app)

	//PROBLEM ROUTES
	setProblemRoutes(app)

	//TAKE ROUTES
	setTakeRoutes(app)

	//CLASS ROUTES
	setClassRoutes(app)

	//PLAYLIST ROUTES
	setPlaylistRoutes(app)

	//START SERVER
	port:=os.Getenv("PORT")
    app.Listen(":"+port)

}