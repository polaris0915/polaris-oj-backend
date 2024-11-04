package main

import (
	"polaris-oj-backend/constant"
	"polaris-oj-backend/database/mysql"

	// "polaris-oj-backend/polaris_oj_backend/allModels"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

// TODO config: 自动生成数据库模型
func exportModels() {
	// 初始化代码生成器
	g := gen.NewGenerator(gen.Config{
		OutPath:      constant.OutPath,      // 输出路径
		ModelPkgPath: constant.ModelPkgPath, // 模型包路径
		// Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})

	// 使用数据库结构生成模型
	g.UseDB(mysql.DB)

	// 自动生成所有表的模型
	// g.ApplyBasic(g.GenerateAllTable()...)

	g.ApplyBasic(
		// ============每次添加外键需要在这里自定义外键关联字段===========
		g.GenerateModel("User"),
		g.GenerateModelAs(
			"question",
			"Question",
			gen.FieldNew(
				"User",
				"*User",
				field.Tag{
					"gorm": "foreignKey:UserID;references:Identity",
					"json": "user",
				},
			),
		),
		g.GenerateModelAs(
			"question_submit",
			"QuestionSubmit",
			gen.FieldNew(
				"User",
				"*User",
				field.Tag{
					"gorm": "foreignKey:UserID;references:Identity",
					"json": "user",
				},
			),
			gen.FieldNew(
				"Question",
				"*Question",
				field.Tag{
					"gorm": "foreignKey:QuestionID;references:Identity",
					"json": "question",
				},
			)),
	)
	// ============每次添加外键需要在这里自定义外键关联字段===========

	// 生成代码
	g.Execute()

}

// 自动生成数据表的CRUD，暂时没用到，也不知道有啥用
// func GenerateAllCRUD() {
// 	g := gen.NewGenerator(gen.Config{
// 		OutPath: "./dao", // 生成代码的输出目录
// 	})
// 	g.UseDB(mysql.DB)

// 	g.ApplyBasic(new(allModels.Question)) // 生成User模型的基本操作代码
// 	g.Execute()
// }

func main() {
	exportModels()
	// GenerateAllCRUD()
}
