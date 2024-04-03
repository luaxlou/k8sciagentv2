# k8sciagent


k8s CI 代理工具 避免直接调用k8s获得过高权限,v2剔除其他无关操作
 
Install 

```bash
go install github.com/luaxlou/k8sciagentv2@latest
```

Usage
```bash
k8sciagent deploy -a appName -i imageName -v version -e env
```