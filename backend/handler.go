package backend

import (
	"context"
	"github.com/m03ed/gozargah-node/common"
)

type Handler interface {
	GetSysStats(context.Context) (*common.BackendStatsResponse, error)
	GetUsersStats(context.Context, bool) (*common.StatResponse, error)
	GetOutboundsStats(context.Context, bool) (*common.StatResponse, error)
	GetInboundsStats(context.Context, bool) (*common.StatResponse, error)
}
