package srv

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mi11/pe/pkg/config"
	"net/http"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	env    *config.Env
	router *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(env *config.Env) (*Server, error) {
	server := &Server{
		env: env,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.GET("/", server.HelloHandler)

	server.router = router
}

func (server *Server) HelloHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, fmt.Sprintf("Hello %s", server.env.Name))
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start() error {
	return http.ListenAndServe(fmt.Sprintf(":%s", server.env.Port), server.router)
}
