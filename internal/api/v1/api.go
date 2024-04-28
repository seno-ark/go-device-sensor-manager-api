package v1

import (
	"go-api/internal/repositories"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	repo     repositories.IRepository
	validate *validator.Validate
}

func NewHandler(validate *validator.Validate, repo repositories.IRepository) *Handler {
	return &Handler{
		repo:     repo,
		validate: validate,
	}
}

func (h *Handler) Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/devices", func(r chi.Router) {
		r.Post("/", h.CreateDevice)
		r.Put("/{device_id}", h.UpdateDevice)
		r.Delete("/{device_id}", h.DeleteDevice)
		r.Get("/", h.GetDeviceList)
		r.Get("/{device_id}", h.GetDevice)
	})

	r.Route("/sensors", func(r chi.Router) {
		r.Get("/types", h.GetSensorTypes)
		r.Post("/", h.CreateSensor)
		r.Put("/{sensor_id}", h.UpdateSensor)
		r.Delete("/{sensor_id}", h.DeleteSensor)
		r.Get("/", h.GetSensorList)
		r.Get("/{sensor_id}", h.GetSensor)
	})

	return r
}
