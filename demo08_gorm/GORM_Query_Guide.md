# GORM 查询操作实用指南 (基础与高级)

本指南涵盖了 GORM 中从基础单条记录查询到复杂高级查询的核心用法，聚焦于日常开发高频场景。

## 1. 基础查询

### A. 检索单个对象 (First/Take/Last)
- **场景**: 根据主键或排序获取一条记录。
- **注意**: 如果找不到记录，会返回 `ErrRecordNotFound` 错误。建议总是检查错误。

```go
var user User

// 获取第一条记录（按主键升序）
db.First(&user)
// SELECT * FROM users ORDER BY id LIMIT 1;

// 获取最后一条记录（按主键降序）
db.Last(&user)

// 根据主键获取 (推荐用法)
db.First(&user, 10) // id = 10
db.First(&user, "10") // id = "10" (字符串主键)
```

### B. 检索全部对象 (Find)
- **场景**: 获取列表数据。
- **注意**: `Find` 不会报 `ErrRecordNotFound`，如果没有数据，切片为空。

```go
var users []User
result := db.Find(&users)
fmt.Println("记录数:", result.RowsAffected)
```

### C. 条件查询 (Where)
最常用的过滤方式，支持字符串、Struct 和 Map。

```go
// 1. 字符串条件 (推荐，防注入)
db.Where("name = ? AND age >= ?", "jinzhu", 22).Find(&users)

// 2. Struct 条件 (注意：零值字段会被忽略！)
// 只有 Name 会作为条件，Age=0 会被忽略
db.Where(&User{Name: "jinzhu", Age: 0}).Find(&users)

// 3. Map 条件 (包含零值)
// Name 和 Age=0 都会作为条件
db.Where(map[string]interface{}{"Name": "jinzhu", "Age": 0}).Find(&users)

// 4. 切片 (IN 查询)
db.Where("id IN ?", []int{1, 2, 3}).Find(&users)
```

---

## 2. 高级查询

### A. 智能选择字段 (Select)
- **场景**: API 接口只需要返回部分字段，减少流量和内存消耗。
- **技巧**: 可以配合定义的 API 专用结构体使用。

```go
type APIUser struct {
  Name string
}

// 自动选择 APIUser 中包含的字段
db.Model(&User{}).Limit(10).Find(&APIUser{})
// SELECT name FROM users LIMIT 10
```

### B. 排序 (Order)
```go
db.Order("age desc, name").Find(&users)
```

### C. 分页 (Limit & Offset)
```go
// page: 页码, pageSize: 每页数量
db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&users)
```

### D. 分组与聚合 (Group & Having)
```go
type Result struct {
    Date  time.Time
    Total int
}
db.Model(&User{}).Select("name, sum(age) as total").Group("name").Scan(&result)
```

### E. 连接查询 (Joins)
适用于简单的关联查询。

```go
db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&result)
```

### F. 锁 (Locking)
在高并发或事务中锁定记录，防止竞态条件。

```go
// SELECT * FROM users FOR UPDATE;
db.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&users)
```

### G. 原始 SQL (Raw)
当 GORM 语法无法满足复杂需求时使用。

```go
db.Raw("SELECT id, name, age FROM users WHERE name = ?", "jinzhu").Scan(&result)
```

## 3. 常见陷阱总结

| 问题 | 现象 | 解决方案 |
| :--- | :--- | :--- |
| **First 报错** | 找不到记录返回 error | 使用 `errors.Is(err, gorm.ErrRecordNotFound)` 判断，或改用 `Find` + `Limit(1)` |
| **Struct 查询** | 零值条件不生效 | 使用 `Map` 或字符串条件 `Where("age = ?", 0)` |
| **Find 性能** | `Find(&user)` 没加 limit | 单条查询务必用 `First` 或 `Take`，或者手动 `Limit(1)` |
