package cache

import (
	"encoding/json"
	"vben/app/admin/model/constants"
	"vben/app/admin/model/system"
	"vben/app/admin/model/tools"
	"vben/pkg/cache/redisCache"
)

// InitDicCache 初始化数据字典到redis
func InitDictCache() {
	var param = tools.SearchTableDataParam{
		Other: system.SysDictType{},
	}
	var typeResult = system.SelectSysDictTypeList(param, false)
	var dataParam = tools.SearchTableDataParam{
		Other: system.SysDictData{},
	}
	dataResult := system.SelectDictDataList(dataParam, false)
	dictMap := make(map[string][]system.SysDictData)
	for _, dictData := range dataResult.Rows.([]system.SysDictData) {
		dictMap[dictData.DictType] = append(dictMap[dictData.DictType], dictData)
	}
	for _, dictType := range typeResult.Rows.([]system.SysDictType) {
		SetDictDataCache(dictType.DictType, dictMap[dictType.DictType])
	}

}

// RefreshDictDataCache 刷新某个字典类型的所有值
func RefreshDictDataCache(dictType string) {
	dataResult := system.SelectDictDataByType(dictType)
	SetDictDataCache(dictType, dataResult)

}

// GetDictDataCacheByType 根据字典类型来获取字典数据
func GetDictDataCacheByType(dictType string) []system.SysDictData {
	dataJson, err := redisCache.NewRedisCache().Get(constants.SysDictCacheKey + dictType)
	if err != nil {
		return nil
	}
	var dictDataList []system.SysDictData
	err = json.Unmarshal([]byte(dataJson), &dictDataList)
	if err != nil || len(dictDataList) <= 0 {
		return nil
	}
	return dictDataList
}

// SetDictDataCache 设置字典类型 值
func SetDictDataCache(dictType string, dictDataList []system.SysDictData) {
	system.OrderByDictSortAsc(dictDataList)
	dataBytes, _ := json.Marshal(dictDataList)
	redisCache.NewRedisCache().Put(constants.SysDictCacheKey+dictType, string(dataBytes), -1)
}

// DeleteDictDataCache 删除字典
func DeleteDictDataCache(dictType string) {
	redisCache.NewRedisCache().Del(constants.SysDictCacheKey + dictType)

}

func DeleteAllDictCache() error {
	keys, _, err := redisCache.NewRedisCache().Scan(0, constants.SysDictCacheKey+"*", constants.ScanCountMax)
	if err != nil {
		return err
	}
	for _, key := range keys {
		redisCache.NewRedisCache().Del(key)
	}
	return nil
}
