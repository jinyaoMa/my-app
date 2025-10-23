package api

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"majinyao.cn/my-app/backend/api/apir"
	"majinyao.cn/my-app/backend/api/endpoints"
	"majinyao.cn/my-app/backend/api/endpoints/auth"
	"majinyao.cn/my-app/backend/api/endpoints/file"
	"majinyao.cn/my-app/backend/api/endpoints/filecategory"
	"majinyao.cn/my-app/backend/api/endpoints/fileextension"
	"majinyao.cn/my-app/backend/api/endpoints/filegroup"
	"majinyao.cn/my-app/backend/api/endpoints/fileuser"
	"majinyao.cn/my-app/backend/api/endpoints/group"
	"majinyao.cn/my-app/backend/api/endpoints/grouprole"
	"majinyao.cn/my-app/backend/api/endpoints/groupuser"
	"majinyao.cn/my-app/backend/api/endpoints/operationidenumpair"
	"majinyao.cn/my-app/backend/api/endpoints/option"
	"majinyao.cn/my-app/backend/api/endpoints/permission"
	"majinyao.cn/my-app/backend/api/endpoints/role"
	"majinyao.cn/my-app/backend/api/endpoints/rolepermission"
	"majinyao.cn/my-app/backend/api/endpoints/user"
	"majinyao.cn/my-app/backend/api/endpoints/userpassword"
	"majinyao.cn/my-app/backend/api/endpoints/userrole"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/internal/service"
	"majinyao.cn/my-app/backend/pkg/api/endpoint"
	"majinyao.cn/my-app/backend/pkg/utils"
)

func setup(ctx context.Context, res *apir.Resources, humaapi huma.API) (operationIdEnumMap map[string]int, err error) {
	// setup endpoints
	var ops []huma.Operation
	for _, op := range []endpoint.Register{
		auth.New(res),
		file.New(res),
		filecategory.New(res),
		fileextension.New(res),
		filegroup.New(res),
		fileuser.New(res),
		group.New(res),
		grouprole.New(res),
		groupuser.New(res),
		operationidenumpair.New(res),
		option.New(res),
		permission.New(res),
		role.New(res),
		rolepermission.New(res),
		user.New(res),
		userpassword.New(res),
		userrole.New(res),
	} {
		ops = append(ops, op.Register(humaapi)...)
	}

	// finalize endpoints
	ops = append(ops, endpoints.New(ops).Register(humaapi)...)

	operationIdEnumPairService, cancel := service.NewOperationIdEnumPairService(ctx, res.Db)
	defer cancel()

	operationIdEnumPairs := entity.OperationIdEnumPairs(utils.SliceMap(ops, func(op huma.Operation) entity.OperationIdEnumPair {
		return entity.OperationIdEnumPair{
			OperationId: op.OperationID,
		}
	}))
	err = operationIdEnumPairService.LoadOrCreate(&operationIdEnumPairs)
	if err != nil {
		return
	}
	return operationIdEnumPairs.ToMap(), nil
}
