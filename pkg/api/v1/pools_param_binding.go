package api

import (
	"errors"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"go.infratographer.com/load-balancer-api/internal/models"
)

// poolsParamsBinding return a set of query mods based
// on the query parameters and path parameters
func (r *Router) poolsParamsBinding(c echo.Context) ([]qm.QueryMod, error) {
	mods := []qm.QueryMod{}

	// optional tenant_id in the request path
	if tenantID, err := r.parseUUID(c, "tenant_id"); err != nil {
		if !errors.Is(err, ErrUUIDNotFound) {
			return nil, err
		}
	} else {
		// found tenant_id in path so add to query mods
		mods = append(mods, models.PoolWhere.TenantID.EQ(tenantID))
		r.logger.Debugw("path param", "tenant_id", tenantID)
	}

	poolID := c.Param("pool_id")
	if poolID != "" {
		if _, err := uuid.Parse(poolID); err != nil {
			return nil, ErrInvalidUUID
		}

		mods = append(mods, models.PoolWhere.PoolID.EQ(poolID))
		r.logger.Debugw("path param", "pool_id", poolID)
	}

	queryParams := []string{"slug", "protocol"}

	qpb := echo.QueryParamsBinder(c)

	for _, qp := range queryParams {
		mods = queryParamsToQueryMods(qpb, qp, mods)

		if len(c.QueryParam(qp)) > 0 {
			r.logger.Debugw("query param", "query_param", qp, "param_vale", c.QueryParam(qp))
		}
	}

	if err := qpb.BindError(); err != nil {
		return nil, err
	}

	return mods, nil
}
