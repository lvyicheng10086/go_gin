# GORM 多对多查询实战：查询学生及其选修课程

本文档总结了使用 GORM 实现“查询学生信息时显示其选修的所有课程”的完整操作步骤。这是一个典型的 **多对多 (Many-to-Many)** 关联查询场景。

## 1. 核心概念

*   **业务场景**：一个学生可以选多门课，一门课也可以被多个学生选。
*   **数据库设计**：需要三张表：
    *   `students` (学生表)
    *   `courses` (课程表)
    *   `student_courses` (中间表，存储关联关系)
*   **GORM 机制**：通过 `Preload` 预加载关联数据，或通过 `many2many` 标签自动处理中间表连接。

---

## 2. 操作步骤详解

### 第一步：定义模型 (Model)

这是最关键的一步，必须正确配置 `many2many` 标签和外键关系。

#### A. 学生表 (Student) - 主表
在学生结构体中定义一个课程切片 `[]Courses`，并配置 Tag。

```go
// model/student/student.go
type Student struct {
    StudentId   int32     `gorm:"primaryKey;column:student_id;type:int"` // 强制 int 类型
    StudentName string    `gorm:"column:student_name;size:255"`
    
    // 关键配置：
    // 1. many2many: 指定中间表名为 student_courses
    // 2. foreignKey: 本表主键 (StudentId)
    // 3. joinForeignKey: 中间表中指向本表的列名 (student_id)
    // 4. References: 对方表主键 (CourseId)
    // 5. joinReferences: 中间表中指向对方表的列名 (course_id)
    Courses     []Courses `gorm:"many2many:student_courses;foreignKey:StudentId;joinForeignKey:student_id;References:CourseId;joinReferences:course_id"`
}
```

#### B. 课程表 (Courses) - 关联表
在课程结构体中也可以定义学生切片（双向关联），Tag 配置逻辑与上面对称。

```go
// model/student/courses.go
type Courses struct {
    CourseId   int32     `gorm:"primaryKey;column:course_id;type:int"`
    CourseName string    `gorm:"column:course_name;size:255"`
    
    // 反向关联配置
    Student    []Student `gorm:"many2many:student_courses;foreignKey:CourseId;joinForeignKey:course_id;References:StudentId;joinReferences:student_id"`
}
```

#### C. 中间表 (StudentCourse) - 连接桥梁
定义中间表结构，确保字段名和类型与主表一致。

```go
// model/student/student_course.go
type StudentCourse struct {
    // ID        int32 `gorm:"primaryKey"` // 可选，中间表通常不需要独立主键
    StudentID int32 `gorm:"column:student_id;type:int"`
    CourseID  int32 `gorm:"column:course_id;type:int"`
}

// 必须实现 TableName 接口，确保表名统一为 student_courses
func (s *StudentCourse) TableName() string {
    return "student_courses"
}
```

---

### 第二步：执行查询 (Controller)

在 Controller 中使用 `Preload` 来加载关联数据。

```go
// controllers/apis/apiController.go

func (u *CourseController) GetCourse(c *gin.Context) {
    // 1. 定义接收结果的变量
    var studentList []student.Student

    // 2. 执行查询
    // Preload("Courses"): 告诉 GORM 在查询学生时，顺便把 Courses 字段填充好
    // Find(&studentList): 查询所有学生
    model.DB.Debug().Preload("Courses").Find(&studentList)

    // 3. 返回结果
    c.JSON(200, gin.H{
        "message": "查询成功",
        "result":  studentList,
    })
}

### 进阶：带条件的关联查询

#### 场景 1：过滤主表数据
**需求**：只查询名字叫“王娜”的学生，并显示她的选修课。

// Where 条件作用于主表 (students)
model.DB.Debug().
    Preload("Courses").                  // 预加载课程
    Where("student_name = ?", "王娜").   // 过滤学生
    Find(&studentList)
```
*   **原理**：先执行 `SELECT * FROM students WHERE student_name = '王娜'`，然后再去查该学生的课程。

#### 场景 2：过滤关联表数据 (Preload 条件)
**需求**：查询所有学生，但只显示他们选修的“Python编程”课（没选这门课的学生，Courses 字段为空）。

```go
// Preload 中的条件作用于关联表 (courses)
model.DB.Debug().
    Preload("Courses", "course_name = ?", "Python编程"). 
    Find(&studentList)
```
*   **原理**：主查询查所有学生。关联查询 SQL 会变成 `SELECT * FROM courses ... WHERE course_name = 'Python编程'`。

---

## 3. 底层原理 (SQL 执行流程)

当你执行 `Preload("Courses").Find(&studentList)` 时，GORM 实际上在后台执行了两条 SQL：

1.  **查询主表**：
    ```sql
    SELECT * FROM students;
    ```
2.  **查询关联表 (自动 Join 中间表)**：
    ```sql
    SELECT courses.*, student_courses.student_id 
    FROM courses 
    INNER JOIN student_courses ON student_courses.course_id = courses.course_id 
    WHERE student_courses.student_id IN (1, 2, 3...); -- 这里填入第一步查出的学生ID
    ```
3.  **组装数据**：GORM 将第二步查出的课程数据，根据 ID 匹配填入到 `studentList` 的 `Courses` 字段中。

## 4. 避坑指南

1.  **Tag 拼写**：`many2many`、`foreignKey`、`joinForeignKey` 等 Tag 必须严格拼写正确，且不能有多余空格。
2.  **非标准主键**：如果主键不是 `ID`（如 `StudentId`），必须显式指定所有 Key，否则 GORM 自动推断会失败（报错 `Unknown column ...`）。
3.  **类型一致性**：结构体中的 ID 类型（如 `int32`）必须与数据库中的列类型（`INT`）一致，防止外键创建失败。
4.  **Preload 参数**：`Preload("Courses")` 里的参数必须是 **Go 结构体中的字段名**（Courses），而不是数据库表名。
