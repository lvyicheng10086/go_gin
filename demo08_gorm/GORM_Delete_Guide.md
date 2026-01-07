# GORM 删除操作实用指南

本指南总结了 GORM 中删除数据的核心方法，包括物理删除和软删除。

## 1. 删除单条记录
删除必须指定主键，否则会触发批量删除保护机制。

```go
var user User
// 方式1: 传入已赋值主键的结构体
user.ID = 10
db.Delete(&user)

// 方式2: 指定主键 (推荐)
db.Delete(&User{}, 10) // DELETE FROM users WHERE id = 10;
```

## 2. 批量删除
通过条件删除匹配的多条记录。

```go
// 删除名为 jinzhu 的所有用户
db.Where("name = ?", "jinzhu").Delete(&User{})

// 通过主键切片删除
db.Delete(&User{}, []int{1, 2, 3})
```

## 3. 软删除 (Soft Delete)
这是 GORM 最强大的特性之一。如果你的模型包含 `gorm.DeletedAt` 字段（`gorm.Model` 默认包含），调用 `Delete` 时不会真正从数据库移除数据，而是设置删除时间。

### A. 触发软删除
```go
// 假设 User 包含 gorm.DeletedAt
db.Delete(&user) 
// SQL: UPDATE users SET deleted_at="2023-10-29 10:00:00" WHERE id = 10;
```

### B. 查询软删除记录
默认情况下，普通查询会自动过滤掉被软删除的记录。

```go
db.Find(&users) // 查不到已删除的记录
```

### C. 查找被软删除的记录 (Unscoped)
如果你想找回被删除的数据（例如回收站功能）。

```go
db.Unscoped().Find(&users) // 包含被软删除的记录
```

### D. 永久删除 (Unscoped + Delete)
彻底从数据库物理删除记录，不可恢复。

```go
db.Unscoped().Delete(&user)
// SQL: DELETE FROM users WHERE id = 10;
```

## 4. 阻止全局删除 (Block Global Delete)
GORM 默认禁止没有任何条件的批量删除，以防删库跑路。

```go
db.Delete(&User{}) // ❌ 报错: GORM: missing WHERE clause
```

如果必须删除全表数据（慎用！），需要显式启用：
```go
// 方式1: 加个无意义条件
db.Where("1 = 1").Delete(&User{})

// 方式2: 原生 SQL
db.Exec("DELETE FROM users")

// 方式3: 开启 AllowGlobalUpdate
db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&User{})
```

## 总结速查表

| 方法 | 作用 | SQL 行为 | 是否可恢复 |
| :--- | :--- | :--- | :--- |
| `Delete` | 软删除 (若有 DeletedAt) | `UPDATE ... SET deleted_at = NOW()` | 是 (`Unscoped` 可查) |
| `Delete` | 物理删除 (若无 DeletedAt) | `DELETE FROM ...` | 否 |
| `Unscoped().Delete` | 强制物理删除 | `DELETE FROM ...` | 否 |
| `Unscoped().Find` | 查询所有 (含已删) | `SELECT * FROM ...` (无 deleted_at 条件) | - |
