# n9e-mac-agent
只适用于V2版本，V3版本不支持

该代码由 https://github.com/didi/nightingale/tree/v2.8.0 和 https://github.com/GitHamburg/agent-machttps://github.com/GitHamburg/agent-mac 修改而来
由于公司内部有部分mac机器需要增加监控，只需要监控到硬盘跟机器的存活监控指标没有windows跟linux多请见谅，其他监控指标可通过修改代码显示。

配置GOPATH 此处忽略

./control build collector

最后提取n9e-collector 跟etc 文件即可使用
使用前需要修改 specify：“” 

./n9e-collector start
