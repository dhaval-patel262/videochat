package server

import(
	"flag"
	"os"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/midlleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/template/html"
	"github.com/gofiber/websocket/v2"
)

var (
	addre = flag.String("addre,":"",os.Geten("PORT"),"")
	cert = flag.String("cert","","")
	key = flag.String("key","","")
)

func Run() error{

	flag.Parse()

	if *addr == ":"{
		*addre = ":8080"
	}

	engine := html.New("./views",".html")
	app := fiber.New(fiber.Config{Views : engine})
	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/", handlers.Welcome)
	app.Get("/room/create", handlers.RoomCreate)
	app.Get("/room/:uuid", handlers.Room)
	app.Get("/room/:uuid/websocket", websocket.New(handlers.RoomWebsocket, Websocket.Config{
		HandleshakeTimeout : 10*time.Second,
		
	}))
	app.Get("/room/:uuid/chat", handlers.RoomChat)
	app.Get("/room/:uuid/chat/websocket", websocket.New(Handlers.RoomChatWebsocket))
	app.Get("/room/:uuid/viewer/websocket", websocket.New(handlers.RoomViewerWebsocket))
	app.Get("/stream/:ssuid",handlers.Stream)
	app.Get("/stream/:ssuid/websocket",)
	app.Get("/stream/:ssuid/chat/websocket", )
	app.Get("/stream/:ssuid/viewer/websocket")


}