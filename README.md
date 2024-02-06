# debug插件

该插件可以使得pprof通过公共端口访问，并且调用profile可以将CPU分析文件保存为cpu.profile文件，然后自动打开分析UI界面。

## 插件地址

https://github.com/Monibuca/plugin-debug

## 插件引入
```go
import (
    _ "m7s.live/plugin/debug/v4"
)
```

## API


### GET `/debug/pprof`
打开pprof界面

### GET `/debug/profile`
默认30s采样，可以通过传入 `?seconds=xxx` 来指定采样时间长度，采样结束后将CPU分析文件保存为cpu.profile文件，并将启动采样展示 web 页面的命令输出到当前页面。比如 `go tool pprof -http :6060 ./monibuca cpu.profile`

在当前monibuca程序目录下，运行命令会自动打开浏览器，展示 profile 调用图。**注意需要先安装graphviz，并添加到环境变量PATH里。**

### GET `/debug/charts`
展示 CPU、内存、GC、pprof 实时变化可视化图表

### WS `/debug/charts/datafeed`
ws 接口，供`/debug/charts`页面前端拉取数据，不需关心

### GET `/debug/charts/data`
jquery 回调接口，供`/debug/charts`页面前端拉取数据，不需关心
