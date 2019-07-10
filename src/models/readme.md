// # 数据库建表遵循 全小写原则 
// 结构体遵循 首字母大写 的 驼峰命名法
// xorm 映射规则 --> NameZh:name_zh

```
// TableName 表名重命名，有这个方法，自动匹配数据库表 star_info
// 不用调用，xorm 会自动识别
func (StarInfo) TableName() string {
	return "star_info"
}
```