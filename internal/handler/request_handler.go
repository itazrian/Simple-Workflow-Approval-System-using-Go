package handler

import (
	"net/http"

	"workflow-approval/internal/model"
	"workflow-approval/internal/service"

	"github.com/labstack/echo/v4"
)

type RequestHandler struct {
	WorkflowService *service.WorkflowService
	StepService     *service.StepService
	RequestService  *service.RequestService
}

func NewRequestHandler(ws *service.WorkflowService, ss *service.StepService, rs *service.RequestService) *RequestHandler {
	return &RequestHandler{
		WorkflowService: ws,
		StepService:     ss,
		RequestService:  rs,
	}
}

func (h *RequestHandler) CreateWorkflow(c echo.Context) error {
	wf := new(model.Workflow)
	if err := c.Bind(wf); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "error": "invalid request"})
	}
	saved, err := h.WorkflowService.CreateWorkflow(wf)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "data": saved, "error": nil})
}

func (h *RequestHandler) GetAllWorkflows(c echo.Context) error {
	data, _ := h.WorkflowService.GetAllWorkflows()
	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "data": data, "error": nil})
}

func (h *RequestHandler) GetWorkflowByID(c echo.Context) error {
	id := c.Param("id")
	data, err := h.WorkflowService.GetWorkflowByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{"success": false, "error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "data": data, "error": nil})
}

func (h *RequestHandler) CreateStep(c echo.Context) error {
	step := new(model.WorkflowStep)
	if err := c.Bind(step); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "error": "invalid request"})
	}
	step.WorkflowID = c.Param("id")
	saved, err := h.StepService.CreateStep(step)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "data": saved, "error": nil})
}

func (h *RequestHandler) GetSteps(c echo.Context) error {
	workflowID := c.Param("id")
	data, err := h.StepService.GetSteps(workflowID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "data": data, "error": nil})
}

func (h *RequestHandler) CreateRequest(c echo.Context) error {
	req := new(model.Request)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "error": "invalid request"})
	}
	saved, err := h.RequestService.CreateRequest(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "data": saved, "error": nil})
}

func (h *RequestHandler) GetRequestByID(c echo.Context) error {
	id := c.Param("id")
	data, err := h.RequestService.GetRequestByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{"success": false, "error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "data": data, "error": nil})
}

func (h *RequestHandler) Approve(c echo.Context) error {
	id := c.Param("id")

	if err := h.RequestService.Approve(id); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": true})
}

func (h *RequestHandler) Reject(c echo.Context) error {
	id := c.Param("id")

	if err := h.RequestService.Reject(id); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": true})
}
