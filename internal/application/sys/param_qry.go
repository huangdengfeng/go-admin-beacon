package sys

import (
	"context"
	"go-admin-beacon/internal/infrastructure/config"
	"go-admin-beacon/internal/infrastructure/response"
)

type ParamCO struct {
	// 上传文件服务器地址
	FileUploadUrl string `json:"fileUploadUrl"`
	// 文件服务器地址
	FileServerUrl string               `json:"fileServerUrl"`
	Dicts         map[string][]*DictCO `json:"dicts"`
}

type DictCO struct {
	Value    any    `json:"value"`
	Name     string `json:"name"`
	Disabled bool   `json:"disabled"`
}

type paramQryExe struct {
}

func NewParamQryExe() *paramQryExe {
	return &paramQryExe{}
}

func (e *paramQryExe) Execute(_ context.Context) (*response.Response, error) {
	instance := config.AppDictInstance
	dictsMap := make(map[string][]*DictCO, len(instance))
	for k, v := range instance {
		dictCOs := make([]*DictCO, 0, len(v))
		for _, dict := range v {
			dictCOs = append(dictCOs, &DictCO{
				Value:    dict.Value,
				Name:     dict.Name,
				Disabled: dict.Disabled,
			})
		}
		dictsMap[k] = dictCOs
	}
	return response.Success(&ParamCO{
		Dicts: dictsMap,
	}), nil
}
