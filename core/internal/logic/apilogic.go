package logic

import (
	"context"

	"github.com/pinkhello/hathor/core/internal/svc"
	"github.com/pinkhello/hathor/core/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) ApiLogic {
	return ApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApiLogic) Api(req types.Request) (*types.Response, error) {
	// add your logic here and delete this line

	return &types.Response{}, nil
}
