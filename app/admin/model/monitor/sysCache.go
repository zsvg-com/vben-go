package monitor

type SysCache struct {
	CacheName  string `json:"cacheName,omitempty"`
	CacheKey   string `json:"cacheKey,omitempty"`
	CacheValue string `json:"cacheValue,omitempty"`
	Remark     string `json:"remark,omitempty"`
}

type LoginUserCache struct {
	TokenId       string `json:"tokenId"`
	DeptName      string `json:"deptName"`
	Ipaddr        string `json:"ipaddr"`
	LoginLocation string `json:"loginLocation"`
	Browser       string `json:"browser"`
	Os            string `json:"os"`
	LoginTime     int64  `json:"loginTime"`
	UserId        int    `json:"userId"`
	UserName      string `json:"userName"`
	Uuid          string `json:"uuid"`
	DeptId        int    `json:"deptId"`
	NickName      string `json:"nickName"`
}
