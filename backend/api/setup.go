package api

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
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

const authScheme = "auth-fwt"

func setup(ctx context.Context, db *gorm.DB, humaapi huma.API) (operationIdEnumMap map[string]int, err error) {
	// setup endpoints
	var ops []huma.Operation
	for _, op := range []endpoint.Register{
		auth.New(authScheme, db),
		file.New(authScheme, db),
		filecategory.New(authScheme, db),
		fileextension.New(authScheme, db),
		filegroup.New(authScheme, db),
		fileuser.New(authScheme, db),
		group.New(authScheme, db),
		grouprole.New(authScheme, db),
		groupuser.New(authScheme, db),
		operationidenumpair.New(authScheme, db),
		option.New(authScheme, db),
		permission.New(authScheme, db),
		role.New(authScheme, db),
		rolepermission.New(authScheme, db),
		user.New(authScheme, db),
		userpassword.New(authScheme, db),
		userrole.New(authScheme, db),
	} {
		ops = append(ops, op.Register(humaapi)...)
	}

	// finalize endpoints
	ops = append(ops, endpoints.New(authScheme, ops).Register(humaapi)...)

	operationIdEnumPairService, cancel := service.NewOperationIdEnumPairService(ctx, db)
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
