package dbresolver

import (
	"gorm.io/gorm"
	"regexp"
)

var fromTableRegexp = regexp.MustCompile("(?i)(?:FROM|UPDATE|MERGE INTO|INSERT [a-z ]*INTO) ['`\"]?([a-zA-Z0-9_]+)([ '`\",)]|$)")

func getTableFromRawSQL(sql string) string {
	if matches := fromTableRegexp.FindAllStringSubmatch(sql, -1); len(matches) > 0 {
		return matches[0][1]
	}

	return ""
}

func GetMapOnlyOne(connPools map[string]gorm.ConnPool) gorm.ConnPool {
	for _, connPool := range connPools {
		return connPool
	}
	return nil
}
