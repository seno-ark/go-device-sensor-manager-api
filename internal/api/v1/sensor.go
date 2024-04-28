package v1

import (
	"encoding/json"
	"go-api/internal/entities"
	"go-api/pkg/util"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// GetSensorTypes get sensor types handler
// @Summary			Get Sensor Types.
// @Description		Get Sensor Types.
// @Tags			Sensors
// @Produce			json
// @Success			200		{object}	util.Response
// @Failure			500		{object}	util.Response
// @Router	/v1/sensors/types [get]
func (h *Handler) GetSensorTypes(w http.ResponseWriter, r *http.Request) {
	resp := util.NewResponse()

	render.Status(r, http.StatusOK)
	render.JSON(w, r, resp.Set("success", entities.SensorTypesName))
}

// CreateSensor create sensor handler
// @Summary			Create Sensor.
// @Description		Create new Sensor.
// @Tags			Sensors
// @Accept			json
// @Param 			json	body		entities.CreateSensorPayload	true	"Sensor data"
// @Produce			json
// @Success			201		{object}	util.Response{data=entities.Sensor}
// @Failure			400		{object}	util.Response
// @Failure			500		{object}	util.Response
// @Router	/v1/sensors [post]
func (h *Handler) CreateSensor(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	resp := util.NewResponse()

	var body entities.CreateSensorPayload
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

	_, err = h.repo.GetDevice(ctx, body.DeviceID)
	if err != nil {
		status, msg := util.ErrStatusCode(err)
		render.Status(r, status)
		render.JSON(w, r, resp.Set(msg, nil))
		return
	}

	sensorID, err := h.repo.CreateSensor(ctx, entities.Sensor{
		DeviceID:    body.DeviceID,
		Type:        body.Type,
		Name:        body.Name,
		Description: body.Description,
	})
	if err != nil {
		status, msg := util.ErrStatusCode(err)
		render.Status(r, status)
		render.JSON(w, r, resp.Set(msg, nil))
		return
	}

	result, err := h.repo.GetSensor(ctx, sensorID)
	if err != nil {
		status, msg := util.ErrStatusCode(err)
		render.Status(r, status)
		render.JSON(w, r, resp.Set(msg, nil))
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, resp.Set("created", result))
}

// UpdateSensor update sensor handler
// @Summary			Update Sensor.
// @Description		Update existing Sensor.
// @Tags			Sensors
// @Accept			json
// @Param 			sensor_id	path	string							true	"Sensor ID" 	example(96a5ec77-9012-4bf3-b08e-39ef4c07fcce)
// @Param 			json		body	entities.UpdateSensorPayload	true	"Sensor data"
// @Produce			json
// @Success			200		{object}	util.Response{data=entities.Sensor}
// @Failure			400		{object}	util.Response
// @Failure			404		{object}	util.Response
// @Failure			500		{object}	util.Response
// @Router	/v1/sensors/{sensor_id} [put]
func (h *Handler) UpdateSensor(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	resp := util.NewResponse()

	sensorID := chi.URLParam(r, "sensor_id")
	if sensorID == "" {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, resp.Set("sensor not found", nil))
		return
	}

	var body entities.UpdateSensorPayload
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

	err = h.repo.UpdateSensor(ctx, sensorID, entities.Sensor{
		Name:        body.Name,
		Description: body.Description,
	})
	if err != nil {
		status, msg := util.ErrStatusCode(err)
		render.Status(r, status)
		render.JSON(w, r, resp.Set(msg, nil))
		return
	}

	result, err := h.repo.GetSensor(ctx, sensorID)
	if err != nil {
		status, msg := util.ErrStatusCode(err)
		render.Status(r, status)
		render.JSON(w, r, resp.Set(msg, nil))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, resp.Set("success", result))
}

// DeleteSensor delete sensor handler
// @Summary			Delete Sensor.
// @Description		Delete Sensor.
// @Tags			Sensors
// @Param			sensor_id		path			string	 true	"Sensor ID" example(96a5ec77-9012-4bf3-b08e-39ef4c07fcce)
// @Produce			json
// @Success			200 			{object}		util.Response
// @Failure			404				{object}		util.Response
// @Failure			500				{object}		util.Response
// @Router	/v1/sensors/{sensor_id} [delete]
func (h *Handler) DeleteSensor(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	resp := util.NewResponse()

	sensorID := chi.URLParam(r, "sensor_id")
	if sensorID == "" {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, resp.Set("sensor not found", nil))
		return
	}

	err := h.repo.DeleteSensor(ctx, sensorID)
	if err != nil {
		status, msg := util.ErrStatusCode(err)
		render.Status(r, status)
		render.JSON(w, r, resp.Set(msg, nil))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, resp.Set("success", nil))
}

// GetSensor get sensor handler
// @Summary			Get sensor by sensor ID.
// @Description		Get sensor by sensor ID.
// @Tags			Sensors
// @Param			sensor_id		path			string	 true	"Sensor ID"
// @Produce			json
// @Success			200 			{object}		util.Response{data=entities.Sensor}
// @Failure			404				{object}		util.Response
// @Failure			500				{object}		util.Response
// @Router	/v1/sensors/{sensor_id} [get]
func (h *Handler) GetSensor(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	resp := util.NewResponse()

	sensorID := chi.URLParam(r, "sensor_id")
	if sensorID == "" {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, resp.Set("sensor not found", nil))
		return
	}

	result, err := h.repo.GetSensor(ctx, sensorID)
	if err != nil {
		status, msg := util.ErrStatusCode(err)
		render.Status(r, status)
		render.JSON(w, r, resp.Set(msg, nil))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, resp.Set("success", result))
}

// GetSensorList get sensor list handler
// @Summary			Get list of Sensor.
// @Description		Get list of Sensor.
// @Tags			Sensors
// @Param			page			query			int	     false	"Pagination page number (default 1, max 500)"				example(1)
// @Param			count			query			int	     false	"Pagination data limit  (default 10, max 100)"				example(10)
// @Param			sort			query			string	 false	"Data sorting (value: name/created_at/updated_at). For desc order, use prefix '-'"	example(-created_at)
// @Param			device_id		query			string	 false	"Filter sensors by device ID"					 			example(96a5ec77-9012-4bf3-b08e-39ef4c07fcce)
// @Param			search			query			string	 false	"Keyword for searching sensors by name or description"		example(soil)
// @Produce			json
// @Success			200 			{object}		util.Response{data=[]entities.Sensor}
// @Failure			500				{object}		util.Response
// @Router	/v1/sensors [get]
func (h *Handler) GetSensorList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	resp := util.NewResponse()

	q := r.URL.Query()
	page, count := util.Pagination(q.Get("page"), q.Get("count"))

	params := entities.GetSensorListParams{
		DeviceID: q.Get("device_id"),
		Search:   q.Get("search"),
		Sort:     q.Get("sort"),
		Limit:    count,
		Offset:   (page - 1) * count,
	}

	results, total, err := h.repo.GetSensorList(ctx, params)
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
