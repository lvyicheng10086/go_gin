# Gin 中间件与控制器数据共享流程

这个流程图展示了从 HTTP 请求发起，经过中间件设置数据，最后在控制器中获取数据并返回的完整过程。

```mermaid
sequenceDiagram
    participant Client as 客户端 (Browser/Postman)
    participant Router as 路由 (Router Group)
    participant Middleware as 中间件 (middlewares.SetValue)
    participant Context as 上下文 (gin.Context)
    participant Controller as 控制器 (AdminController.User)

    Note over Client, Router: 1. 发起 GET /admin/user 请求

    Client->>Router: HTTP Request
    Router->>Middleware: 匹配路由组，进入中间件

    Note over Middleware, Context: 2. 中间件执行

    activate Middleware
    Middleware->>Context: c.Set("name", "张三")
    Note right of Context: Context 内部存储: <br/>Key: "name"<br/>Value: "张三"
    Middleware->>Router: 中间件执行完毕 (隐式/显式 Next)
    deactivate Middleware

    Router->>Controller: 转发到最终处理函数

    Note over Controller, Context: 3. 控制器获取数据

    activate Controller
    Controller->>Context: c.Get("name")
    Context-->>Controller: 返回 interface{} ("张三")
    
    Controller->>Controller: 类型断言 name.(string)
    
    Controller-->>Client: c.JSON 响应
    deactivate Controller

    Note left of Client: 4. 收到 JSON 响应: <br/>{"msg": "用户列表", "name": "张三"}
```

## 关键步骤说明

1.  **路由匹配**：请求 `/admin/user` 命中 `adminRouters` 路由组。
2.  **中间件执行**：
    *   在进入控制器之前，先执行 `middlewares.SetValue`。
    *   **关键动作**：`c.Set("name", "张三")` 将数据挂载到本次请求的上下文 `gin.Context` 中。就像给这个请求贴了一个标签。
3.  **控制器执行**：
    *   `AdminController.User` 被调用。
    *   **关键动作**：`c.Get("name")` 从上下文中把刚才贴的标签取下来。
4.  **响应**：控制器将取到的数据通过 `c.JSON` 返回给客户端。
