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
 *   
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
	app.Post("/api/video",auth.Protect, auth.AdminOnly, controllers.CreateVideo)
	app.Get("/api/video",auth.Protect, controllers.GetAllVideos)
	app.Get("/api/video/:id",auth.Protect, controllers.GetVideo)
	app.Patch("/api/video/:id",auth.Protect,auth.AdminOnly, controllers.UpdateVideo)
	app.Delete("/api/video/:id",auth.Protect, auth.AdminOnly, controllers.DeleteVideo)
}

func setExerciseRoutes(app *fiber.App){
	app.Post("/api/exercise", auth.Protect, auth.AdminOnly, controllers.CreateExercise)
	app.Get("/api/exercise", auth.Protect, controllers.GetAllExercises)
    app.Get("/api/exercise/:id",auth.Protect, controllers.GetExercise)
    app.Patch("/api/exercise/:id",auth.Protect, auth.AdminOnly, controllers.UpdateExercise)
    app.Delete("/api/exercise/:id",auth.Protect, auth.AdminOnly, controllers.DeleteExercise)
}

func setProblemRoutes(app *fiber.App) {
	app.Post("/api/problem",auth.Protect, auth.AdminOnly, controllers.CreateProblem)
    app.Get("/api/problem",auth.Protect, controllers.GetAllProblems)
    app.Get("/api/problem/:id",auth.Protect, controllers.GetProblem)
    app.Patch("/api/problem/:id",auth.Protect, auth.AdminOnly, controllers.UpdateProblem)
    app.Delete("/api/problem/:id",auth.Protect, auth.AdminOnly, controllers.DeleteProblem)
}

func setTakeRoutes(app *fiber.App){
	app.Post("/api/take",auth.Protect, controllers.CreateTake)
    app.Get("/api/take",auth.Protect, auth.AdminOnly, controllers.GetAllTakes)
	app.Get("/api/take/:id",auth.Protect, controllers.GetTake)
    app.Patch("/api/take/:id",auth.Protect, controllers.UpdateTake)
    app.Delete("/api/take/:id",auth.Protect, auth.AdminOnly, controllers.DeleteTake)
}

func setClassRoutes(app *fiber.App) {
	app.Post("/api/class",auth.Protect, auth.TeacherAndAdminOnly, controllers.CreateClass)
    app.Get("/api/class",auth.Protect, controllers.GetAllClasses)
	app.Get("/api/class/:id",auth.Protect, controllers.GetClass)
    app.Patch("/api/class/:id",auth.Protect, auth.TeacherAndAdminOnly, controllers.UpdateClass)
    app.Delete("/api/class/:id",auth.Protect, auth.TeacherAndAdminOnly, controllers.DeleteClass)
}

func setPlaylistRoutes(app *fiber.App) {
	app.Post("/api/playlist",auth.Protect, controllers.CreatePlaylist)
    app.Get("/api/playlist",auth.Protect, controllers.GetAllPlaylists)
    app.Get("/api/playlist/:id",auth.Protect, controllers.GetPlaylist)
    app.Patch("/api/playlist/:id",auth.Protect, controllers.UpdatePlaylist)
    app.Delete("/api/playlist/:id",auth.Protect, controllers.DeletePlaylist)
}

//!!!!!!!!!!!!!!! remove comments later

func setUserRoutes(app *fiber.App) {
	app.Post("/api/user", auth.Protect, controllers.CreateUser)
    app.Get("/api/user", auth.Protect, auth.AdminOnly, controllers.GetAllUsers)
    app.Get("/api/user/:id", auth.Protect, controllers.GetUser)
    app.Patch("/api/user/:id", auth.Protect, controllers.UpdateUser)
    app.Delete("/api/user/:id",auth.Protect, auth.AdminOnly, controllers.DeleteUser)
}

func main() {
    app := fiber.New()
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