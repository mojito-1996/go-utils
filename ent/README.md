## Ent 框架

### Field 字段属性
- GoType 
- Default 字段默认值
- Other
- Comment 字段注释
- Deprecated  将字段标记为已弃用。已弃用的字段在查询中默认不会被选中，并且其在生成代码中的结构体字段会被标注为 Deprecated
- Optional 可选字段
- Immutable 不可变字段
- Validate  字段验证，参数为一个验证函数
  ```go
  package schema
  
  type Group struct {
    ent.Schema
  }
  func MaxRuneCount(maxLen int) func(s string) error {
    return func(s string) error {
        if utf8.RuneCountInString(s) > maxLen {
            return errors.New("value is more than the max length")
        }
        return nil
    }
  }
  func (Group) Fields() []ent.Field {
    return []ent.Field{
      field.String("name").
        Match(regexp.MustCompile("[a-zA-Z_]+$")).
        Validate(func(s string) error {
          if strings.ToLower(s) == s { 
            return errors.New("group name must begin with uppercase")
          }
          return nil
        }),
      }
    }
  ```
- Annotations
- StorageKey
- Unique 唯一字段，不能有默认值
- Nillable 可以为空的字段