package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"strings"
	"time"
)

type LimitConfRules struct {
	Rules map[string]*LimitOpt `mapstructure:"rules"`
}

type LimitOpt struct {
	Interval int64 `mapstructure:"interval"`
	Capacity int64 `mapstructure:"capacity"`
	Quantum  int64 `mapstructure:"quantum"`
}

type LimiterIface interface {
	Key(context *gin.Context) string
	GetBucket(key string) (*ratelimit.Bucket, bool)
	AddBucketsByUri(uri string, fillInterval, capacity, quantum int64) LimiterIface
	AddBucketByConf() LimiterIface
}

type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

type UriLimiter struct {
	*Limiter
	Rule *LimitConfRules
}

func NewUriLimiter() LimiterIface {
	return &UriLimiter{
		Limiter: &Limiter{
			limiterBuckets: make(map[string]*ratelimit.Bucket),
		},
	}
}

func (l *UriLimiter) Key(context *gin.Context) string {
	uri := context.Request.RequestURI
	index := strings.Index(uri, "?")
	if index == -1 {
		return uri
	}
	return uri[:index]
}

func (l *UriLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := l.limiterBuckets[key]
	return bucket, ok
}

func (l *UriLimiter) AddBucketsByUri(uri string, fillInterval, capacity, quantum int64) LimiterIface {
	bucket := ratelimit.NewBucketWithQuantum(time.Second*time.Duration(fillInterval), capacity, quantum)
	l.limiterBuckets[uri] = bucket
	return l
}

func (l *UriLimiter) getConf() *LimitConfRules {

	rule := &LimitConfRules{Rules: map[string]*LimitOpt{
		"/": {
			Interval: 10, //多长时间添加令牌
			Capacity: 1,  //令牌桶的容量
			Quantum:  1,  //到达定时器指定的时间，往桶里面加多少令牌
		},
	}}

	return rule
}

func (l *UriLimiter) AddBucketByConf() LimiterIface {
	rule := l.getConf()
	for k, v := range rule.Rules {
		l.AddBucketsByUri(k, v.Interval, v.Capacity, v.Quantum)
	}
	return l
}
