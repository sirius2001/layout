package db

// 所有同步到数据库中数据
var models = []any{}

func AutoMerge() error {
	return db.AutoMigrate(models...)
}
