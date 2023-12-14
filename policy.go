package dbresolver

import (
	"context"
	"math/rand"

	"gorm.io/gorm"
)

type Policy interface {
	Resolve(context.Context, map[string]gorm.ConnPool) gorm.ConnPool
}

type RandomPolicy struct {
}

const TargetDB = "targetDB"

func (RandomPolicy) Resolve(ctx context.Context, connPools map[string]gorm.ConnPool) gorm.ConnPool {
	u, ok := ctx.Value(TargetDB).(string)
	if ok {
		return connPools[u]
	}
	return randomGetMap(connPools)
}

func randomGetMap(connPools map[string]gorm.ConnPool) gorm.ConnPool {
	list := make([]gorm.ConnPool, 0, len(connPools))
	for _, connPool := range connPools {
		list = append(list, connPool)
	}
	return list[rand.Intn(len(connPools))]
}
