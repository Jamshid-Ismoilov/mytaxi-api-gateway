package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"

	_ "github.com/Jamshid-Ismoilov/mytaxi-api-gateway/api/handlers/models"
	pb "github.com/Jamshid-Ismoilov/mytaxi-api-gateway/genproto"
	l "github.com/Jamshid-Ismoilov/mytaxi-api-gateway/pkg/logger"
	// "github.com/Jamshid-Ismoilov/mytaxi-api-gateway/pkg/utils"
)

// CreateDriver ...
// @Summary CreateDriver
// @Description This API for creating a new driver
// @Tags driver
// @Accept  json
// @Produce  json
// @Param Driver request body models.CreateDriverBody true "driverCreateRequest"
// @Success 200 {object} models.Driver
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/drivers/ [post]
func (h *handlerV1) CreateDriver(c *gin.Context) {
	var (
		body        pb.Driver
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.TaxiService().CreateDriver(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create driver", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetDriver ...
// @Summary GetDriver
// @Description This API for getting driver detail
// @Tags driver
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.Driver
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/drivers/{id} [get]
func (h *handlerV1) GetDriver(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.TaxiService().GetDriver(
		ctx, &pb.ByIdReq{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateDriver ...
// @Summary UpdateDriver
// @Description This API for updating driver
// @Tags driver
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param Driver request body models.CreateDriverBody true "driverUpdateRequest"
// @Success 200 {object} models.Driver
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/drivers/{id} [put]
func (h *handlerV1) UpdateDriver(c *gin.Context) {
	var (
		body        pb.Driver
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	body.Id = c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.TaxiService().UpdateDriver(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteDriver ...
// @Summary DeleteDriver
// @Description This API for deleting driver
// @Tags driver
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/drivers/{id} [delete]
func (h *handlerV1) DeleteDriver(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.TaxiService().DeleteDriver(
		ctx, &pb.ByIdReq{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}



// CreateClient ...
// @Summary CreateClient
// @Description This API for creating a new client
// @Tags client
// @Accept  json
// @Produce  json
// @Param Client request body models.CreateClientBody true "clientCreateRequest"
// @Success 200 {object} models.Client
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/clients/ [post]
func (h *handlerV1) CreateClient(c *gin.Context) {
	var (
		body        pb.Client
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.TaxiService().CreateClient(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create driver", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetClient ...
// @Summary GetClient
// @Description This API for getting client detail
// @Tags client
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.Client
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/clients/{id} [get]
func (h *handlerV1) GetClient(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.TaxiService().GetClient(
		ctx, &pb.ByIdReq{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateClient ...
// @Summary UpdateClient
// @Description This API for updating client
// @Tags client
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param Client request body models.CreateClientBody true "clientUpdateRequest"
// @Success 200 {object} models.Client
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/clients/{id} [put]
func (h *handlerV1) UpdateClient(c *gin.Context) {
	var (
		body        pb.Client
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	body.Id = c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.TaxiService().UpdateClient(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteClient ...
// @Summary DeleteClient
// @Description This API for deleting client
// @Tags client
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/clients/{id} [delete]
func (h *handlerV1) DeleteClient(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.TaxiService().DeleteClient(
		ctx, &pb.ByIdReq{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}



// CreateOrder ...
// @Summary CreateOrder
// @Description This API for creating a new order
// @Tags order
// @Accept  json
// @Produce  json
// @Param Client request body models.CreateOrder true "clientCreateRequest"
// @Success 200 {object} models.FullOrder
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/orders/ [post]
func (h *handlerV1) CreateOrder(c *gin.Context) {
	var (
		body        pb.Order
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.TaxiService().CreateOrder(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create driver", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetOrder ...
// @Summary GetOrder
// @Description This API for getting order detail
// @Tags order
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.FullOrder
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/orders/{id} [get]
func (h *handlerV1) GetOrder(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.TaxiService().GetClient(
		ctx, &pb.ByIdReq{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateOrder ...
// @Summary UpdateOrder
// @Description This API for updating order
// @Tags order
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param Order request body models.UpdateOrder true "orderUpdateRequest"
// @Success 200 {object} models.FullOrder
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/orders/{id} [put]
func (h *handlerV1) UpdateOrder(c *gin.Context) {
	var (
		body        pb.Order
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	body.Id = c.Param("id")

	if body.Status != "accepted" && body.Status != "cancelled" && body.Status != "finished" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("cannot change standart status variables", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.TaxiService().UpdateOrder(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteOrder ...
// @Summary DeleteOrder
// @Description This API for deleting order
// @Tags order
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/orders/{id} [delete]
func (h *handlerV1) DeleteClient(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.TaxiService().DeleteClient(
		ctx, &pb.ByIdReq{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
