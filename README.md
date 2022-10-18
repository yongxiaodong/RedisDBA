# RedisDBA  
redis key在线扫描。直接链接redis进行在线扫描（默认配置下，qps大约4W，可根据redis性能自行调整配置）


- 未设置TTL的key扫描  

` ./bin/RedisDBA_linux_amd64 -a nottl`

- 大key前50名扫描

`./bin/RedisDBA_linux_amd64 -a bigkey`


## 输出目录

>> 程序运行目录中的result目录



## 配置文件

Addr: redis链接地址
Port: redis端口
Password: redis密码
DB: redis的DB，redis DB默认为0
PullKeysCount: 20000    一次拉取多少个key（根据redis配置自定义）
PipeQueryCount: 2000    一次查询多少个key（根据redis配置自定义）
ConsumerNum: 3  协程数量（根据redis配置自定义）
