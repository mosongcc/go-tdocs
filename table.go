package tdocs

import (
	"reflect"
	"strings"
)

var (
	tables = make([]Table, 0) // 表数据缓存
)

type Table struct {
	Name        string
	Title       string
	Description string
	Fields      []Field
}

type Field struct {
	Name    string
	Type    string
	Comment string
}

type tableName interface {
	TableName() string
	TableTitle() string
	TableDescription() string
}

// Register 注册表结构视图
// 参数 tag  读取字段名字的标签
// 参数 models 数据库表结构体
func Register(tag Tag, models ...tableName) {
	for _, t := range models {
		table := Table{
			Name:        t.TableName(),
			Title:       t.TableTitle(),
			Description: t.TableDescription(),
			Fields:      make([]Field, 0),
		}
		typeOfTable := reflect.TypeOf(t)
		for i := 0; i < typeOfTable.NumField(); i++ {
			field := typeOfTable.Field(i)
			name := fieldName(tag, field.Tag.Get(string(tag)))
			if name != "" {
				table.Fields = append(table.Fields, Field{
					Name:    name,
					Type:    field.Type.Name(),
					Comment: field.Tag.Get("comment"),
				})
			}
		}
		tables = append(tables, table)
	}
}

type Tag string

const (
	Bson Tag = "bson" // 用于Mongodb
	Json Tag = "json" // 用于JSON
	Gorm Tag = "gorm" // 用于gorm.io/gorm
)

func fieldName(tag Tag, v string) string {
	if v == "" || v == "-" {
		return ""
	}
	if tag == Bson {
		return bsonName(v)
	}
	if tag == Json {
		return jsonName(v)
	}
	if tag == Gorm {
		return gormName(v)
	}
	return ""
}

func bsonName(v string) string {
	return strings.ReplaceAll(v, ",omitempty", "")
}

func jsonName(v string) string {
	return strings.ReplaceAll(v, ",omitempty", "")
}

func gormName(v string) string {
	if strings.HasPrefix(v, "column:") {
		v = v[7:]
	}
	return ""
}
