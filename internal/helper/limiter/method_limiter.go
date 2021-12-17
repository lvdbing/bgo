package limiter

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

type MethodLimit struct {
	*Limiter
}

func NewMethodLimiter() LimiterInterface {
	l := &Limiter{buckets: make(map[string]*ratelimit.Bucket)}
	return MethodLimit{
		Limiter: l,
	}
}

func (ml MethodLimit) Key(c *gin.Context) string {
	uri := c.Request.RequestURI
	index := strings.Index(uri, "?")
	if index == -1 {
		return uri
	}
	return uri[:index]
}

func (ml MethodLimit) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := ml.buckets[key]
	return bucket, ok
}

func (ml MethodLimit) AddBuckets(rules ...BucketRule) LimiterInterface {
	for _, rule := range rules {
		if _, ok := ml.buckets[rule.Key]; !ok {
			bucket := ratelimit.NewBucketWithQuantum(rule.FillInterval, rule.Capacity, rule.Quantum)
			ml.buckets[rule.Key] = bucket
		}
	}
	return ml
}
