# Casbin 笔记

## Casbin 是什么？
Casbin 是一个授权库，可用于希望特定用户或 `subject` 访问特定 `object` 或实体的流程中。
访问类型即 `action`，可以是读取、写入、删除或开发人员设置的任何其他操作。
这就是 Casbin 最广泛的使用方式，被称为 "标准" 或经典的 { subject, object, action } 流程。
另外 Casbin 支持大部分的语言。

## 工作原理

在 Casbin 中，访问控制模型被抽象为基于 **PERM 元模型（Policy, Effect, Request, Matchers）** 的 CONF 文件。
切换或升级项目的授权机制就像修改配置一样简单。你可以通过组合可用的模型来定制自己的访问控制模型。
例如，你可以在一个模型内将 RBAC 角色和 ABAC 属性结合在一起，共享一套策略规则。

PERM 模型由四个基础部分组成：Policy、Effect、Request 和 Matchers。这些基础描述了资源和用户之间的关系。

### Request 请求
定义 request 参数。基本 request 是一个元组对象，至少需要一个 subject（访问实体）、object（访问资源）和 action（访问方法）。

例如， request 定义可能是这样的：`r={sub,obj,act}`。

该定义指定了访问控制匹配函数所需的参数名称和顺序。

### Policy 策略
定义访问 policy 的模型。它指定了 policy 规则文件中字段的名称和顺序。

例如： `p={sub, obj, act}` 或 `p={sub, obj, act, eft}`。

注意：如果未定义 eft（policy 结果），则不会读取 policy 文件中的结果字段，匹配的 policy 结果将默认 allowed。

### Matcher 匹配器
定义 Request 和 Policy 的匹配规则。

例如： `e = some(where(p.eft == allow))` 。

这句话的意思是，如果匹配 policy 结果 `p.eft` 某些为 `allow`，那么最终结果为 true。

让我们看另一个例子：

`e = some(where (p.eft == allow)) && !some(where (p.eft == deny))`

这个示例组合的逻辑含义是：如果有策略与 `allow` 的结果相匹配，而没有策略与 `deny` 的结果相匹配，则结果为 true。
换句话说，当匹配的策略都是 `allow` 时，结果为真。如果有任何 `deny`，则两者都为 false（更简单地说，当 allow 和 deny 同时存在时，deny 优先）。