# Shell Tool 简化状态流转图

## 🎯 核心状态流转

```mermaid
graph TD
    A[🚀 程序启动] --> B[⚙️ INIT 初始化]
    B --> C[✅ READY 就绪]
    
    C --> D{👤 用户操作}
    D -->|📋 查看帮助/状态| C
    D -->|⚙️ 配置管理| C
    D -->|📝 创建任务| C
    D -->|▶️ 执行任务| E[🔄 BUSY 忙碌]
    D -->|🚪 退出程序| F[👋 EXIT 退出]
    
    E --> G{🔄 任务状态}
    G -->|✅ 成功完成| C
    G -->|❌ 执行失败| H[⚠️ ERROR 错误]
    G -->|⏸️ 用户中断| I[⏸️ PAUSED 暂停]
    G -->|🛑 系统终止| F
    
    I --> J{⏸️ 暂停操作}
    J -->|▶️ 恢复执行| E
    J -->|❌ 取消任务| C
    J -->|🛑 强制退出| F
    
    H --> K{⚠️ 错误处理}
    K -->|🔄 可恢复| C
    K -->|💀 致命错误| F
    
    F --> L[🧹 清理资源]
    L --> M[💾 保存配置]
    M --> N[👋 程序结束]
    
    style A fill:#e1f5fe
    style C fill:#c8e6c9
    style E fill:#fff3e0
    style I fill:#e3f2fd
    style H fill:#ffebee
    style F fill:#f3e5f5
```

## 🔄 状态转换事件

```mermaid
flowchart LR
    subgraph "状态转换事件"
        A1[system_start] --> B1[初始化完成]
        A2[task_start] --> B2[开始任务]
        A3[task_complete] --> B3[任务完成]
        A4[interrupt] --> B4[用户中断]
        A5[resume] --> B5[恢复执行]
        A6[cancel] --> B6[取消任务]
        A7[user_exit] --> B7[用户退出]
        A8[terminate] --> B8[系统终止]
        A9[recover] --> B9[错误恢复]
    end
    
    subgraph "状态"
        S1[INIT]
        S2[READY]
        S3[BUSY]
        S4[PAUSED]
        S5[ERROR]
        S6[EXIT]
    end
    
    S1 -->|A1| S2
    S2 -->|A2| S3
    S3 -->|A3| S2
    S3 -->|A4| S4
    S3 -->|A8| S6
    S4 -->|A5| S3
    S4 -->|A6| S2
    S4 -->|A8| S6
    S2 -->|A7| S6
    S5 -->|A9| S2
    S5 -->|A8| S6
```

## 🎮 用户交互流程

```mermaid
sequenceDiagram
    participant U as 👤 用户
    participant S as 🖥️ Shell Tool
    participant T as 📋 任务系统
    participant C as ⚙️ 配置系统
    
    Note over U,S: 程序启动
    S->>S: 初始化系统
    S->>C: 加载配置
    S->>U: 显示欢迎信息
    
    Note over U,S: 主交互循环
    loop 用户输入循环
        U->>S: 输入命令
        S->>S: 解析命令
        
        alt 帮助命令
            S->>U: 显示帮助信息
        else 状态命令
            S->>U: 显示当前状态
        else 配置命令
            S->>C: 处理配置
            C->>U: 返回配置结果
        else 任务命令
            alt 创建任务
                S->>T: 创建新任务
                T->>U: 返回任务ID
            else 列出任务
                S->>T: 获取任务列表
                T->>U: 显示任务列表
            else 执行任务
                S->>T: 开始执行任务
                T->>S: 任务执行中
                S->>U: 显示执行状态
                T->>S: 任务完成
                S->>U: 显示执行结果
            end
        else 退出命令
            S->>C: 保存配置
            S->>U: 显示再见信息
            S->>S: 退出程序
        end
    end
```

## 🎯 关键状态特征

### 📊 状态特征表

| 状态 | 图标 | 特征 | 可执行操作 | 用户交互 |
|------|------|------|-----------|----------|
| **INIT** | ⚙️ | 系统初始化 | 无 | 显示初始化进度 |
| **READY** | ✅ | 等待用户输入 | 所有命令 | 完整交互 |
| **BUSY** | 🔄 | 任务执行中 | 中断操作 | 显示进度/中断 |
| **PAUSED** | ⏸️ | 任务暂停 | 恢复/取消 | 选择操作 |
| **ERROR** | ⚠️ | 错误状态 | 恢复/退出 | 错误信息 |
| **EXIT** | 👋 | 程序退出 | 清理操作 | 再见信息 |

### 🔄 状态转换规则

```mermaid
graph TD
    subgraph "状态转换规则"
        R1[严格验证] --> R2[记录历史]
        R2 --> R3[执行退出动作]
        R3 --> R4[更新状态]
        R4 --> R5[执行进入动作]
        R5 --> R6[记录日志]
    end
    
    subgraph "转换类型"
        T1[直接转换] --> T2[当前状态:目标状态]
        T3[通配符转换] --> T4[*:目标状态]
    end
    
    subgraph "验证逻辑"
        V1[检查直接转换] --> V2{存在?}
        V2 -->|是| V3[允许转换]
        V2 -->|否| V4[检查通配符]
        V4 --> V5{存在?}
        V5 -->|是| V3
        V5 -->|否| V6[拒绝转换]
    end
```

## 🎨 状态可视化

### 状态指示器
```bash
# 不同状态的提示符样式
INIT:   "⚙️  正在初始化..."
READY:  "✅ [READY] shell> "
BUSY:   "🔄 [BUSY] 执行中... "
PAUSED: "⏸️ [PAUSED] 已暂停 "
ERROR:  "⚠️ [ERROR] 发生错误 "
EXIT:   "👋 正在退出..."
```

### 状态颜色编码
- 🟢 **READY**: 绿色 - 系统就绪
- 🟡 **BUSY**: 黄色 - 正在工作
- 🔵 **PAUSED**: 蓝色 - 暂停状态
- 🔴 **ERROR**: 红色 - 错误状态
- 🟣 **EXIT**: 紫色 - 退出状态

## 🚀 实际运行示例

```bash
# 启动程序
$ ./shell_tool.sh

# 状态流转示例
✅ [READY] shell> task create "测试任务" "sleep 5"
📝 任务已创建: 测试任务 (ID: task_1)

✅ [READY] shell> task start task_1
🔄 [BUSY] 开始执行任务: task_1
🔄 [BUSY] 正在执行任务...
✅ [READY] 任务执行成功: task_1

✅ [READY] shell> task create "长时间任务" "sleep 30"
📝 任务已创建: 长时间任务 (ID: task_2)

✅ [READY] shell> task start task_2
🔄 [BUSY] 开始执行任务: task_2
🔄 [BUSY] 正在执行任务...
# 用户按 Ctrl+C
⏸️ [PAUSED] 任务已暂停

⏸️ [PAUSED] shell> resume
🔄 [BUSY] 任务已恢复
🔄 [BUSY] 正在执行任务...
✅ [READY] 任务执行完成

✅ [READY] shell> exit
👋 正在退出...
💾 正在保存配置...
👋 再见！
```

这个状态流转设计确保了：
- 🎯 **清晰的状态边界**: 每个状态都有明确的职责
- 🔄 **可控的状态转换**: 只有预定义的转换才被允许
- 🛡️ **完善的错误处理**: 各种异常情况都有处理机制
- 👤 **友好的用户体验**: 状态变化有清晰的视觉反馈
- 🔧 **易于扩展**: 可以轻松添加新状态和转换规则
