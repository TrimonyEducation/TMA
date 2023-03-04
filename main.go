package main

import (
	"golang-crud/auth"
	"golang-crud/controllers"
	"golang-crud/utils"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

/**------------------------------------------------------------------------
 * todo                             TODO
 *   QUERY PARAMS FOR ALL READER (SEARCH BY SUBJECT);
 *   ADD isPUBLISHED, isFinished, IsPublic, uuid TO MODELS
 *   Join tables
 *   READ ALL CONTROLLERS
 *   CHECK ALL CONTROLLERS
 *   ADD ROLE BASED AUTH TO CONTROLLERS
 *   INPUT VALIDATION
 *------------------------------------------------------------------------**/
//!!!!!!!!!!!!!!!!! output for getAll should contain results (len of records) and data should be called data

func init(){
	err:=godotenv.Load()
	if err!= nil {
        log.Println(err.Error())
	}
	utils.ConnectDB()
	auth.InitFireBase()
}

func setVideoRoutes(app *fiber.App){
	app.Post("/api/video", controllers.CreateVideo)
	app.Get("/api/video", controllers.GetAllVideos)
	app.Get("/api/video/:id", controllers.GetVideo)
	app.Patch("/api/video/:id", controllers.UpdateVideo)
	app.Delete("/api/video/:id", controllers.DeleteVideo)
}

func setExerciseRoutes(app *fiber.App){
	app.Post("/api/exercise",  controllers.CreateExercise)
	app.Get("/api/exercise",  controllers.GetAllExercises)
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
	app.Post("/api/class", auth.TeacherAndAdminOnly, controllers.CreateClass)
    app.Get("/api/class", controllers.GetAllClasses)
	app.Get("/api/class/:id", controllers.GetClass)
    app.Patch("/api/class/:id", auth.TeacherAndAdminOnly, controllers.UpdateClass)
    app.Delete("/api/class/:id", auth.TeacherAndAdminOnly, controllers.DeleteClass)
}

func setPlaylistRoutes(app *fiber.App) {
	app.Post("/api/playlist", controllers.CreatePlaylist)
    app.Get("/api/playlist", controllers.GetAllPlaylists)
    app.Get("/api/playlist/:id", controllers.GetPlaylist)
    app.Patch("/api/playlist/:id", controllers.UpdatePlaylist)
    app.Delete("/api/playlist/:id", controllers.DeletePlaylist)
}

//!!!!!!!!!!!!!!! remove comments later

func setUserRoutes(app *fiber.App) {
	app.Post("/api/user",  controllers.CreateUser)
    app.Get("/api/user",  controllers.GetAllUsers)
    app.Get("/api/user/:id",  controllers.GetUser)
    app.Patch("/api/user/:id",  controllers.UpdateUser)
    app.Delete("/api/user/:id", controllers.DeleteUser)
}

func main() {
    app := fiber.New(fiber.Config{
		BodyLimit: 4 * 1024 * 1024,
	})
	app.Use(helmet.New())
	app.Use(limiter.New(limiter.Config{
		Max: 50,
		Expiration: 60 * time.Second,
		LimitReached: func(c *fiber.Ctx) error{
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{"msg": "limit reached"})
		},
	}))
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

	//USER ROUTES
	setUserRoutes(app)

	//START SERVER
	port:=os.Getenv("PORT")
    app.Listen(":"+port)

}