    ```mermaid
    graph TD
    A["copilot/handler/handler.go<br/>接口处理层"] --> B["copilot/service/suggestion.go<br/>建议服务"]
    A --> C["copilot/service/completion.go<br/>补全服务"]
    B --> D["copilot/model/suggestion.go<br/>建议模型"]
    C --> D
    B --> E["copilot/dal/suggestion_query.go<br/>数据访问层"]
    C --> E
    F["copilot/pack/response.go<br/>响应打包"] --> A
    G["api/kitex/copilot.thrift<br/>接口定义"] --> A
    H["copilot/test/handler_test.go<br/>单元测试"] --> A

    click A "https://code.byted.org/vecode/vecode/merge_requests/4272?to_version=4&dv_filepath=module%2Fcopilot%2Fhandler%2Fhandler.go"
    ```
