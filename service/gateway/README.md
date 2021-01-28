## Gateway

工作台 Gateway 服务

### Blacklist

由于 token 中包含用户所在团队和权限 role 的信息，所以需要设置一个黑名单让 token 无效。

token 无效的情况：
1. 移出团队
2. 更改用户权限（role）

#### Usage

Functions are in package `muxi-workbench-gateway/service`

将 token 放入黑名单

```go
// @token: token
// @expiresAt: 过期时间（10位的时间戳）
func AddToBlacklist(token string, expiresAt int64) error
```
