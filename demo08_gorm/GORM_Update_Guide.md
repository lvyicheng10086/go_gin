# GORM 更新操作实用指南

本指南总结了 GORM 框架中日常开发最常用的更新数据方法。

## 1. 保存所有字段 (Save)

`Save` 是一个全能方法，既可用于**新增**也可用于**更新**（Upsert）。

- **场景**: 无论记录是否存在，都保存所有字段的值（包括零值）。
- **特点**: 必须包含主键，否则会变成创建新记录。

```go
var user User
db.First(&user) // 先查询
user.Name = "新名字"
user.Age = 100
db.Save(&user) // 更新所有字段
```

## 2. 更新单个字段 (Update)

当只想修改某个特定字段时使用。

- **场景**: 修改状态、更新计数器等。
- **特点**: 需要通过 `Model()` 指定对象（带主键）或 `Where()` 指定条件。

```go
// 方式1: 通过 Model 指定对象
db.Model(&user).Update("name", "hello")

// 方式2: 通过 Where 指定条件
db.Model(&User{}).Where("active = ?", true).Update("name", "hello")
```

## 3. 更新多个字段 (Updates)

这是最常用的批量更新方法，支持 Struct 和 Map。

### A. 使用 Struct 更新

- **特点**: **只更新非零值字段**。如果字段值为 0, "", false 等零值，会被 GORM 忽略，**不会更新到数据库**。
- **场景**: 只需要更新部分有值的字段。

```go
// 仅更新 Name 和 Age (如果 Age 不为 0)
db.Model(&user).Updates(User{Name: "hello", Age: 18})
```

### B. 使用 Map 更新

- **特点**: **更新所有指定的字段**，包括零值。
- **场景**: 需要将字段强制更新为 0 或 false 时必须使用 Map。

```go
// 即使 Active 为 false 也会被更新
db.Model(&user).Updates(map[string]interface{}{"name": "hello", "active": false})
```

## 4. 强制更新选定字段 (Select/Omit)

当使用 Struct 更新但又想更新零值，或者想忽略某些字段时使用。

- **Select**: 只更新选中的字段（即使是零值也会更新）。
- **Omit**: 忽略选中的字段。

```go
// 强制更新 Name 和 Age (即使 Age 为 0)
db.Model(&user).Select("Name", "Age").Updates(User{Name: "new_name", Age: 0})

// 更新除 Role 以外的所有字段
db.Model(&user).Omit("Role").Updates(User{Name: "jinzhu", Role: "admin"})
```

## 5. 批量更新 (Batch Updates)

如果不指定主键 ID，而是使用 Where 条件匹配多条记录，则会触发批量更新。

```go
// 将所有 role 为 admin 的用户 Age 改为 18
db.Model(User{}).Where("role = ?", "admin").Updates(User{Age: 18})
```

## 6. 获取更新结果

可以通过返回值的 `RowsAffected` 检查有多少条记录被修改。

```go
result := db.Model(&user).Update("name", "jinzhu")
fmt.Println("更新行数:", result.RowsAffected)
fmt.Println("错误信息:", result.Error)
```

## 总结速查表

| 方法               | 作用      | 是否更新零值      | 适用场景                                          |
| :----------------- | :-------- | :---------------- | :------------------------------------------------ |
| `Save`             | 保存/新增 | 是 (所有字段)     | 对象整体保存，或者不确定是新增还是更新            |
| `Update`           | 更新单列  | 是                | 只修改一个字段                                    |
| `Updates (Struct)` | 更新多列  | **否** (忽略零值) | 更新部分字段，且不需要把值改为 0/false            |
| `Updates (Map)`    | 更新多列  | 是                | 更新部分字段，且**包含**要把值改为 0/false 的情况 |
| `Select + Updates` | 指定字段  | 是                | 精确控制哪些字段必须更新（包括零值）              |

|      |      |      |      |
| :--- | :--- | :--- | :--- |
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |
