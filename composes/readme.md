# 执行的命令: 

1. 启动 configsvr
```shell
docker-compose -f docker-compose-configsvr-configsvr-repl-1.yaml up -d
```
2. 加入 configsvr 节点
```shell
docker exec -it xxx bash
echo "rs.initiate({_id: 'configsvr-repl-1',configsvr: true,members: [{ _id : 0, host : '172.254.0.218:20001' },{ _id : 1, host : '172.254.0.218:20002' },{ _id : 2, host : '172.254.0.218:20003' }]})" > init.js
mongo admin --port="20001" init.js
```

3. 启动 shard 节点
```shell
docker-compose -f docker-compose-sharding-repl-1.yaml up -d 
```

4. 加入shard节点
```shell
docker exec xxx -it bash
echo "rs.initiate({_id: 'shard-repl-1',members: [{ _id : 0, host : '172.254.0.218:21001' },{ _id : 1, host : '172.254.0.218:21002' },{ _id : 2, host : '172.254.0.218:21003' }]})" > init.js
mongo admin --port="21001" init.js
```