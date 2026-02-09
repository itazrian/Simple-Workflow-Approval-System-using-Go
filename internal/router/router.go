package router

import (
	"workflow-approval/internal/handler"

	"github.com/labstack/echo/v4"
)

func Register(e *echo.Echo, h *handler.RequestHandler) {

	e.POST("/workflows", h.CreateWorkflow)
	e.GET("/workflows", h.GetAllWorkflows)
	e.GET("/workflows/:id", h.GetWorkflowByID)

	e.POST("/workflows/:id/steps", h.CreateStep)
	e.GET("/workflows/:id/steps", h.GetSteps)

	e.POST("/requests", h.CreateRequest)
	e.GET("/requests/:id", h.GetRequestByID)
	e.POST("/requests/:id/approve", h.Approve)
	e.POST("/requests/:id/reject", h.Reject)
}
