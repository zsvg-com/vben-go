package cache

import (
	"vben/app/admin/model/constants"
	"vben/app/admin/model/system"
	"vben/app/admin/model/tools"
	"vben/pkg/cache/redisCache"
)

// InitSysConfigCache 初始化参数配置到redis
func InitSysConfigCache() {
	var param = tools.SearchTableDataParam{
		Other: system.SysConfig{},
	}
	result := system.SelectConfigList(param, false)
	list := result.Rows.([]system.SysConfig)
	for _, config := range list {
		SetSysConfigCache(config.ConfigKey, config.ConfigValue)
	}
}

// RefreshSysConfigCache 刷新某个参数配置的所有值
func RefreshSysConfigCache(configKey string) {
	config := system.SelectConfigByKey(configKey)
	SetSysConfigCache(configKey, config)
}

// GetSysConfigCacheByKey 根据参数配置来获取字典数据
func GetSysConfigCacheByKey(configKey string) string {
	get, err := redisCache.NewRedisCache().Get(constants.SysConfigCacheKey + configKey)
	if err != nil {
		return ""
	}
	return get
}

// SetSysConfigCache 设置参数配置 值
func SetSysConfigCache(configKey string, configValue string) {
	redisCache.NewRedisCache().Put(constants.SysConfigCacheKey+configKey, configValue, -1)
}

// DeleteSysConfigCache 删除字典
func DeleteSysConfigCache(configKey string) {
	redisCache.NewRedisCache().Del(constants.SysConfigCacheKey + configKey)
}

func DeleteAllSysConfigCache() error {
	keys, _, err := redisCache.NewRedisCache().Scan(0, constants.SysConfigCacheKey+"*", constants.ScanCountMax)
	if err != nil {
		return err
	}
	for _, key := range keys {
		redisCache.NewRedisCache().Del(key)
	}
	return nil
}
