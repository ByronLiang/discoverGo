# Golang 项目提交脚本

## 代码规范检验

将`.pre-commit` 与 `init.sh` 放置仓库根目录, 实现配置 Git Hook 提交代码检验

当对`.pre-commit`变更过, 需要重新执行`init.sh` 脚本, 确保下一次的 `git commit` 生效
