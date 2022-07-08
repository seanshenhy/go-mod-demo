package demo

import (
	"github.com/go-redis/redis/v7"
)

type Demo struct {
	redis *redis.Client
}

func NewDemo(rds *redis.Client) *Demo {
	return &Demo{
		redis: rds,
	}
}

func TestDemo1(a int) int {
	return a + 12
}
