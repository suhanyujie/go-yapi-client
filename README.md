# go yapi client
* Go 实现的 yapi 提交客户端，用于将本地的 markdown 文档提交到公司/组织的 yapi 文档管理中心

## Usage（todo）
[ ] 列举出所有已配置的 token flag
[ ] 参数个数为 2 个时，表示直接提交某个md文档到 yapi 上，如 yc /some/path/a.md，token 从环境变量总获取 YC_TOKEN
[ ] 参数个数为 3 个时，如 yc file /some/path/a.md 同第一种功能
[ ] 参数个数为 3 个时，如 yc dir /some/path 表示使用环境变量中的 token 提交 path 目录下的所有 markdown 文档
[ ] 参数个数为 4 个时，如 yc file someTokenFlag /some/path/a.md 表示使用 someTokenFlag 对应的 token 提交该 markdown 文档
[ ] 参数个数为 4 个时，如 yc dir someTokenFlag /some/path 表示使用 someTokenFlag 对应的 token 提交目录 /some/path 下的 markdown 文档

## reference
* https://github.com/urfave/cli
* cli 介绍 https://juejin.im/post/5ef293ccf265da02ee5ef187
