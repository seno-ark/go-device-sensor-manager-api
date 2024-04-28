package v1

import (
	"encoding/json"
	"go-api/internal/entities"
	"go-api/pkg/util"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// CreateDevice create device handler
// @Summary			Create Device.
// @Description		Create new Device.
// @Tags			Devices
// @Accept			json
// @Produce			json
// @Param 			json	body		entities.CreateUpdateDevicePayload	true	"Device data"
// @Success			201		{object}	util.Response{data=entities.Device}
// @Failure			400		{object}	util.Response
// @Failure			500		{object}	util.Response
// @Router	/v1/devices [post]
func (h *Handler) CreateDevice(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	resp := util.NewResponse()

	var body entities.CreateUpdateDevicePayload
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, resp.Set("invalid data", nil))
		return
	}

	err = h.validate.Struct(body)
	if err != nil {
		errs := util.ParseValidatorErr(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, resp.Set("invalid data", nil).AddErrValidation(errs))
		return
	}

	deviceID, err := h.repo.CreateDevice(ctx, entities.Device{
		Name:        body.Name,
		Description: body.Description,
		Status:      body.Status,
	})
	if err != nil {
		status, msg := util.ErrStatusCode(err)
		render.Status(r, status)
		render.JSON(w, r, resp.Set(msg, nil))
		return
	}

	result, err := h.repo.GetDevice(ctx, deviceID)
	if err != nil {
		status, msg := util.ErrStatusCode(err)
		render.Status(r, status)
		render.JSON(w, r, resp.Set(msg, nil))
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, resp.Set("created", result))
}

// UpdateDevice update device handler
// @Summary			Update Device.
// @Description		Update existing Device.
// @Tags			Devices
// @Accept			json
// @Param 			device_id	path	string								true	"Device ID" example(01HQSH92SNYQVCBDSD38XNBRYM)
// @Param 			json		body	entities.CreateUpdateDevicePayload	true	"Device data"
// @Produce			json
// @Success			200		{object}	util.Response{data=entities.Device}
// @Failure			400		{object}	util.Response
// @Failure			404		{object}	util.Response
// @Failure			500		{object}	util.Response
// @Router	/v1/devices/{device_id} [put]
func (h *Handler) UpdateDevice(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	resp := util.NewResponse()

	deviceID := chi.URLParam(r, "device_id")
	if deviceID == "" {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, resp.Set("device not found", nil))
		return
	}

	var body entities.CreateUpdateDevicePayload
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, resp.Set("invalid data", nil))
		return
	}

	err = h.validate.Struct(body)
	if err != nil {
		errs := util.ParseValidatorErr(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, resp.Set("invalid data", nil).AddErrValidation(errs))
		return
	}

	err = h.repo.UpdateDevice(ctx, deviceID, entities.Device{
		Name:        body.Name,
		Description: body.Description,
		Status:      body.Status,
	})
	if err != nil {
		status, msg := util.ErrStatusCode(err)
		render.Status(r, status)
		render.JSON(w, r, resp.Set(msg, nil))
		return
	}

	result, err := h.repo.GetDevice(ctx, deviceID)
	if err != nil {
		status, msg := util.ErrStatusCode(err)
		render.Status(r, status)
		render.JSON(w, r, resp.Set(msg, nil))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, resp.Set("success", result))
}

// DeleteDevice delete device handler
// @Summary			Delete Device.
// @Description		Delete Device.
// @Tags			Devices
// @Param			device_id		path			string	 true	"Device ID" example(01HQSH92SNYQVCBDSD38XNBRYM)
// @Produce			json
// @Success			200 			{object}		util.Response
// @Failure			404				{object}		util.Response
// @Failure			500				{object}		util.Response
// @Router	/v1/devices/{device_id} [delete]
func (h *Handler) DeleteDevice(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	resp := util.NewResponse()

	deviceID := chi.URLParam(r, "device_id")
	if deviceID == "" {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, resp.Set("device not found", nil))
		return
	}

	err := h.repo.DeleteDevice(ctx, deviceID)
	if err != nil {
		status, msg := util.ErrStatusCode(err)
		render.Status(r, status)
		render.JSON(w, r, resp.Set(msg, nil))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, resp.Set("success", nil))
}

// GetDevice get device handler
// @Summary			Get device by device ID.
// @Description		Get device by device ID.
// @Tags			Devices
// @Param			device_id		path			string	 true	"Device ID"
// @Produce			json
// @Success			200 			{object}		util.Response{data=entities.Device}
// @Failure			404				{object}		util.Response
// @Failure			500				{object}		util.Response
// @Router	/v1/devices/{device_id} [get]
func (h *Handler) GetDevice(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	resp := util.NewResponse()

	deviceID := chi.URLParam(r, "device_id")
	if deviceID == "" {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, resp.Set("device not found", nil))
		return
	}

	result, err := h.repo.GetDevice(ctx, deviceID)
	if err != nil {
		status, msg := util.ErrStatusCode(err)
		render.Status(r, status)
		render.JSON(w, r, resp.Set(msg, nil))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, resp.Set("success", result))
}

// GetDeviceList get device list handler
// @Summary			Get list of Device.
// @Description		Get list of Device.
// @Tags			Devices
// @Produce			json
// @Param			page			query			int	     false	"Pagination page number (default 1, max 500)"				example(1)
// @Param			count			query			int	     false	"Pagination data limit  (default 10, max 100)"				example(10)
// @Param			sort			query			string	 false	"Data sorting (value: name/created_at/updated_at). For desc order, use prefix '-'"	example(-created_at)
// @Param			search			query			string	 false	"Keyword for searching device by title or content" 			example(raspi)
// @Success			200 			{object}		util.Response{data=[]entities.Device}
// @Failure			500				{object}		util.Response
// @Router	/v1/devices [get]
func (h *Handler) GetDeviceList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	resp := util.NewResponse()

	q := r.URL.Query()
	page, count := util.Pagination(q.Get("page"), q.Get("count"))

	params := entities.GetDeviceListParams{
		Search: q.Get("search"),
		Sort:   q.Get("sort"),
		Limit:  count,
		Offset: (page - 1) * count,
	}

	results, total, err := h.repo.GetDeviceList(ctx, params)
	if err != nil {
		status, msg := util.ErrStatusCode(err)
		render.Status(r, status)
		render.JSON(w, r, resp.Set(msg, nil))
		return
	}

	resp.AddMeta(page, count, total)
	render.Status(r, http.StatusOK)
	render.JSON(w, r, resp.Set("success", results))
}
