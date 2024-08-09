package test

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"testing"
)

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func TestGEN(t *testing.T) {
	// 连接数据库
	dsn := "root:a111111@(47.109.31.106:3306)/redu?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}

	// 生成实例
	g := gen.NewGenerator(gen.Config{
		OutPath:      "../query",                                                         // 定义 dao 文件输出目录
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		ModelPkgPath: "./output/entity",                                                  // 定义 model 文件输出目录
	})

	//启用数据库连接
	g.UseDB(db) // reuse your gorm db

	//生成单个表的model 若只为了生成model则不需要接受参数，如下生成所有表model的示例
	//userModel := g.GenerateModel("menu_table")
	//生成所有表的model
	g.GenerateAllTable()

	//根据model生成dao文件
	// Generate basic type-safe DAO API for struct `model.User` following conventions
	//g.ApplyBasic(userModel)

	//根据所有model生成dao文件
	//g.ApplyBasic(g.GenerateAllTable()...)

	//对于已经存在model的情况，可以传入该model的实例，如下所示
	//g.ApplyBasic(model.User{})

	//根据接口生成自定义方法
	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	//g.ApplyInterface(func(Querier) {}, userModel)

	//执行
	// Generate the code
	g.Execute()
}
