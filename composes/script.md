
### 在单机部署3个mongoshard 节点

1. 运行 container

```shell
k run -it --image=mongo:4.4 mongo bash
```

2. 进入容器，启动3个configsvr节点

```shell
# 创建各自的目录. 
mkdir /data/c{1,2,3}
mkdir /data/shard{1,2,3}

# 启动3个configsvr节点.
for i in {1..3}; do
   nohup mongod --configsvr --replSet="configsrv" --bind_ip_all --port="2700${i}" --dbpath="/data/c${i}" > "/tmp/c${i}.log" 2>&1 &
done

# 打印相关节点. 
root@mongo:/data# ps -axf
    PID TTY      STAT   TIME COMMAND
      1 pts/0    Ss     0:00 bash
     31 pts/0    Sl     0:00 mongod --configsvr --replSet=configsrv --bind_ip_all --port=27001
     32 pts/0    Sl     0:00 mongod --configsvr --replSet=configsrv --bind_ip_all --port=27002
     33 pts/0    Sl     0:00 mongod --configsvr --replSet=configsrv --bind_ip_all --port=27003
    178 pts/0    R+     0:00 ps -axf

```

3. 相互加入节点

```
echo "rs.initiate({_id: 'configsrv',configsvr: true,members: [{ _id : 0, host : '172.254.0.218:20001' },{ _id : 1, host : '172.254.0.218:20002' },{ _id : 2, host : '172.254.0.218:20003' }]})" > init.js

for i in {1..3}; do
    mongo admin --port="2700${i}" init.js
done
rm init.js

# 查看各个节点的状态
mongo admin --port 27001 --eval 'rs.status()'
mongo admin --port 27002 --eval 'rs.status()'
mongo admin --port 27003 --eval 'rs.status()'
```

4. 部署shard节点

```

# 部署节点
for i in {1..3}; do
   nohup mongod --shardsvr --replSet="shardsvr" --bind_ip_all --port="2701${i}" --dbpath="/data/shard${i}" > "/tmp/s${i}.log" 2>&1 &
done

# 加入节点
echo "rs.initiate({_id: 'shardsvr',members: [{ _id : 0, host : '172.254.0.218:27011' },{ _id : 1, host : '172.254.0.218:27012' },{ _id : 2, host : '172.254.0.218:27013' }]})" > init.js

for i in {1..3}; do
    mongo admin --port="2701${i}" init.js
done
rm init.js

```

5. 部署 mongos 节点

```

nohup mongos --configdb configsvr/172.254.0.218:27001,172.254.0.218:27002,172.254.0.218:27003 --bind_ip_all > /tmp/mongos.log 2>&1 &

```

6. 添加节点
```

mongo admin --port="21001" --eval 'sh.addShard( "shardsvr/172.254.0.218:21001,172.254.0.218:21002,172.254.0.218:21003")'



```

7. 对某个数据库/collection启用分片

```

# 对某个数据库
sh.enableSharding("<database>")
sh.shardCollection("<database>.<collection>", { <shard key field> : "hashed" } )
sh.shardCollection("<database>.<collection>", { <shard key field> : 1, ... } )

```


8. 对分片集群进行测试. 
```
sh.enableSharding("test")
sh.enableSharding("test.users",{_id: "hashed"})
k port-forward mongo 27017
```