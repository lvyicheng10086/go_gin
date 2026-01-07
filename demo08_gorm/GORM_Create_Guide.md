# GORM 创建操作实用指南

本指南涵盖了 GORM 中创建（插入）数据的核心方法，包括单条插入、批量插入、指定字段以及关联创建。

## 1. 基础创建

### A. 创建单条记录
- **注意**: 必须传入**指针**，以便 GORM 回填主键 ID。
- **返回值**: `result.RowsAffected` (影响行数), `result.Error` (错误信息)。

```go
user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

// 必须传指针 &user
result := db.Create(&user) 

fmt.Println(user.ID)             // 创建成功后 ID 会被自动回填
fmt.Println(result.Error)        // nil 表示成功
fmt.Println(result.RowsAffected) // 1
```

### B. 创建多条记录 (Slice)
直接传入切片即可。

```go
users := []*User{
    {Name: "Jinzhu", Age: 18},
    {Name: "Jackson", Age: 19},
}

result := db.Create(users) // 传入切片
// users[0].ID 也会被自动回填
```

---

## 2. 字段控制

### A. 指定字段创建 (Select)
只插入指定的字段，其他字段使用数据库默认值或零值。

```go
// 只插入 Name 和 Age，忽略 CreatedAt 等其他字段
db.Select("Name", "Age").Create(&user)
// INSERT INTO users (name, age) VALUES ("Jinzhu", 18);
```

### B. 忽略字段创建 (Omit)
插入除指定字段外的所有字段。

```go
// 忽略 Age，插入其他所有字段
db.Omit("Age").Create(&user)
```

---

## 3. 批量插入 (Batch Insert)

当数据量非常大时，直接 `Create` 可能会生成一条超长的 SQL。使用 `CreateInBatches` 可以分批次执行。

```go
var users = []User{{Name: "u1"}, ..., {Name: "u10000"}}

// 每次插入 100 条
db.CreateInBatches(users, 100)
```

---

## 4. 高级创建

### A. 关联创建 (Associations)
如果在创建主表数据时，结构体里包含了关联对象（如 `CreditCard`），GORM 默认会**自动创建**这些关联数据。

```go
type User struct {
  gorm.Model
  Name       string
  CreditCard CreditCard // 关联结构体
}

// 会自动向 users 表和 credit_cards 表都插入数据
db.Create(&User{
  Name: "jinzhu",
  CreditCard: CreditCard{Number: "411111111111"},
})
```

如果不希望自动创建关联，可以使用 `Omit`:
```go
// 跳过 CreditCard 关联的创建
db.Omit("CreditCard").Create(&user)
```

### B. 使用 Map 创建
当没有定义 Struct，或者只想快速测试时，可以使用 Map。
- **缺点**: 不会执行 Hook 方法，也不会自动回填 ID。

```go
db.Model(&User{}).Create(map[string]interface{}{
  "Name": "jinzhu",
  "Age":  18,
})
```

### C. 默认值 (Default Value)
在 Struct Tag 中定义默认值。
- **注意**: 只有当字段为**零值**（0, "", false）时，才会使用 Tag 中的默认值。

```go
type User struct {
  ID   int64
  Name string `gorm:"default:galeone"` // 默认值 galeone
  Age  int64  `gorm:"default:18"`      // 默认值 18
}
```

## 总结速查表

| 方法 | 场景 | 注意事项 |
| :--- | :--- | :--- |
| `Create(&user)` | 标准单条插入 | 必传指针，会自动回填 ID |
| `Create(users)` | 批量插入 | 传入 Slice |
| `CreateInBatches` | 大数据量批量插入 | 推荐用于上千条数据，避免 SQL 过长 |
| `Select(...).Create` | 指定字段插入 | 安全性更高，防止误写入 |
| `Omit(...).Create` | 忽略字段插入 | 常用于跳过关联创建 |
