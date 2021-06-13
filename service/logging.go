package service

import (
	"willow/config"
	"willow/global"
	"willow/response"
)

type Log struct {
	Value string `json:"value"`
}

func (l *Log) Search() response.Response {
	if !config.Conf.ES.Enable {
		return response.Error(response.EsNotEnable)
	}
	ret, err := global.ES.Search(config.Conf.ES.Index, config.Conf.ES.Key, l.Value)
	if err != nil {
		return response.Error(response.EsSearchError)
	}
	return response.Success(ret)
}
