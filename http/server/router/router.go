package route

import (
	"chat/ctx"
	"chat/model"
	"chat/service"
	"github.com/gin-gonic/gin"
)

type Router struct {
	serviceContext              ctx.ServiceContext
	transcriptTranslatorService *service.TranscriptTranslatorService
}

func NewRouter(serviceContext ctx.ServiceContext) *Router {
	return &Router{
		serviceContext:              serviceContext,
		transcriptTranslatorService: service.NewTranscriptTranslatorService(serviceContext),
	}
}

func (r *Router) Install(engine *gin.Engine) {
	prefix := "/api/v1/sales"
	router := engine.Group(prefix)
	router.POST("/translate", r.translate)
}

func (r *Router) translate(c *gin.Context) {
	req := new(model.TranslateTranscriptRequest)
	err := c.BindJSON(req)
	if err != nil {
		r.newBadRequestResponse(c, model.InvalidRequest)
		return
	}

	res, err := r.transcriptTranslatorService.Translate(req)
	if err != nil {
		r.newBadRequestResponse(c, err.Error())
		return
	}

	r.newSuccessResponse(c, res)
}
