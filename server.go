package main

import (
	"fmt"
	"net/http"

	"github.com/ovasquezbrito/golang-rest-api/repository"

	"github.com/ovasquezbrito/golang-rest-api/service"

	"github.com/ovasquezbrito/golang-rest-api/controller"

	router "github.com/ovasquezbrito/golang-rest-api/http"
)

var (
	postRepository repository.PostRepository = repository.NewFirestoreRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewChiRouter()
)

func main() {
	const port string = ":8000"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(port)

}
