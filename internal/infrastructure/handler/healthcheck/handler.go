package healthcheck

import (
	"github.com/go-openapi/runtime/middleware"

	"github/user-manager/internal/generated/server/models"
	"github/user-manager/internal/generated/server/restapi/operations/healthcheck"
)

type Handler struct {
	healthCheckers []func() (string, bool)
}

func NewHandler(healthCheckers ...func() (string, bool)) *Handler {
	return &Handler{
		healthCheckers: healthCheckers,
	}
}

func (h *Handler) HealthCheck(_ healthcheck.GetHealthCheckParams) middleware.Responder {
	fail := false
	resources := make([]*models.HealthCheckResource, 0, len(h.healthCheckers))
	for _, healthChecker := range h.healthCheckers {
		resource, isAvailable := healthChecker()
		resources = append(resources, &models.HealthCheckResource{
			Resource:    &resource,
			IsAvailable: &isAvailable,
		})

		if !isAvailable {
			fail = true
		}
	}

	if fail {
		return healthcheck.NewGetHealthCheckInternalServerError().WithPayload(
			&healthcheck.GetHealthCheckInternalServerErrorBody{
				Resources: resources,
			},
		)
	}

	return healthcheck.NewGetHealthCheckOK().WithPayload(
		&healthcheck.GetHealthCheckOKBody{
			Resources: resources,
		},
	)
}
