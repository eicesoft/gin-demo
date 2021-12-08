package config_service

import (
	"eicesoft/web-demo/internal/model"
	"eicesoft/web-demo/internal/model/sys_configs"
	"eicesoft/web-demo/pkg/core"
	"eicesoft/web-demo/pkg/db"
)

var _ ConfigService = (*configService)(nil)

type ConfigService interface {
	List(ctx core.Context) []*sys_configs.Config
	Create(ctx core.Context, configInfo *ConfigInfo) (id int32, err error)
	Update(ctx core.Context, configInfo *ConfigInfo) (err error)
}

type ConfigInfo struct {
	Title string `form:"title"` //配置的介绍
	Key   string `form:"key"`   //配置的Key
	Value string `form:"value"` //配置的值
	Type  int32  `form:"type"`  //配置类型[1=>数值,2=>文本,3=>JSON]
}

type configService struct {
	db db.Repo
}

func NewConfigService(db db.Repo) *configService {
	return &configService{
		db: db,
	}
}

func (o *configService) Update(ctx core.Context, configInfo *ConfigInfo) (err error) {
	data := map[string]interface{}{
		"status": 100,
	}

	err = sys_configs.NewQueryBuilder().
		WhereKey(model.EqualPredicate, configInfo.Key).
		Updates(o.db.GetDbW().WithContext(ctx.RequestContext()), data)

	return
}

// Create 创建配置
func (o *configService) Create(ctx core.Context, configInfo *ConfigInfo) (id int32, err error) {
	configModel := sys_configs.NewModel()
	configModel.Assign(configInfo)
	configModel.IsDelete = 0
	id, err = configModel.Create(o.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		panic(err)
	}
	return
}

func (o *configService) List(ctx core.Context) []*sys_configs.Config {
	configList, err := sys_configs.
		NewQueryBuilder().
		Limit(10).
		WhereIsDelete(model.EqualPredicate, 0).
		QueryAll(o.db.GetDbR().WithContext(ctx.RequestContext()))

	if err != nil {
		ctx.Logger().Error(err.Error())
	}

	return configList
}
