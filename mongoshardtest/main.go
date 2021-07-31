package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
		return
	}

	coll := func() *mongo.Collection {
		return client.Database("test").Collection("users")
	}

	tk := time.NewTicker(1 * time.Second)
	// 插入 3 条数据
	for i := 0; i < 100000; i++ {

		var users []interface{}
		for j := i; j < j+100; j++ {
			users = append(users, &User{
				Username:     fmt.Sprintf("user%d", i),
				Password:     "从结果 winningPlan 中可以看出执行计划是否高效，比如：\n\n未能命中索引的结果，会显示COLLSCAN\n命中索引的结果，使用IXSCAN\n出现了内存排序，显示为 SORT\n关于 explain 的结果说明，可以进一步参考文档：\n\nhttps://docs.mongodb.com/manual/reference/explain-results/index.html\n\n五、集群\n在大数据领域常常提到的4V特征中，Volume(数据量大)是首当其冲被提及的。\n由于单机垂直扩展能力的局限，水平扩展的方式则显得更加的靠谱。 MongoDB 自带了这种能力，可以将数据存储到多个机器上以提供更大的容量和负载能力。\n此外，同时为了保证数据的高可用，MongoDB 采用副本集的方式来实现数据复制。\n\n一个典型的MongoDB集群架构会同时采用分片+副本集的方式，如下图：",
				Description:  "从结果 winningPlan 中可以看出执行计划是否高效，比如：\n\n未能命中索引的结果，会显示COLLSCAN\n命中索引的结果，使用IXSCAN\n出现了内存排序，显示为 SORT\n关于 explain 的结果说明，可以进一步参考文档：\n\nhttps://docs.mongodb.com/manual/reference/explain-results/index.html\n\n五、集群\n在大数据领域常常提到的4V特征中，Volume(数据量大)是首当其冲被提及的。\n由于单机垂直扩展能力的局限，水平扩展的方式则显得更加的靠谱。 MongoDB 自带了这种能力，可以将数据存储到多个机器上以提供更大的容量和负载能力。\n此外，同时为了保证数据的高可用，MongoDB 采用副本集的方式来实现数据复制。\n\n一个典型的MongoDB集群架构会同时采用分片+副本集的方式，如下图：",
				Description1: "从结果 winningPlan 中可以看出执行计划是否高效，比如：",
				Description2: "数据分片（Shards）\n分片用于存储真正的集群数据，可以是一个单独的 Mongod实例，也可以是一个副本集。 生产环境下Shard一般是一个 Replica Set，以防止该数据片的单点故障。\n对于分片集合(sharded collection)来说，每个分片上都存储了集合的一部分数据(按照分片键切分)，如果集合没有分片，那么该集合的数据都存储在数据数据分片（Shards）\n分片用于存储真正的集群数据，可以是一个单独的 Mongod实例，也可以是一个副本集。 生产环境下Shard一般是一个 Replica Set，以防止该数据片的单点故障。\n对于分片集合(sharded collection)来说，每个分片上都存储了集合的一部分数据(按照分片键切分)，如果集合没有分片，那么该集合的数据都存储在数据数据分片（Shards）\n分片用于存储真正的集群数据，可以是一个单独的 Mongod实例，也可以是一个副本集。 生产环境下Shard一般是一个 Replica Set，以防止该数据片的单点故障。\n对于分片集合(sharded collection)来说，每个分片上都存储了集合的一部分数据(按照分片键切分)，如果集合没有分片，那么该集合的数据都存储在数据数据分片（Shards）\n分片用于存储真正的集群数据，可以是一个单独的 Mongod实例，也可以是一个副本集。 生产环境下Shard一般是一个 Replica Set，以防止该数据片的单点故障。\n对于分片集合(sharded collection)来说，每个分片上都存储了集合的一部分数据(按照分片键切分)，如果集合没有分片，那么该集合的数据都存储在数据数据分片（Shards）\n分片用于存储真正的集群数据，可以是一个单独的 Mongod实例，也可以是一个副本集。 生产环境下Shard一般是一个 Replica Set，以防止该数据片的单点故障。\n对于分片集合(sharded collection)来说，每个分片上都存储了集合的一部分数据(按照分片键切分)，如果集合没有分片，那么该集合的数据都存储在数据",
				Description3: "未能命中索引的结果，会显示COLLSCAN",
				Description4: "命中索引的结果，使用IXSCAN",
				Description5: "出现了内存排序，显示为 SORT",
				Description6: "关于 explain 的结果说明，可以进一步参考文档：",
				Description7: "数据分片（Shards）\n分片用于存储真正的集群数据，可以是一个单独的 Mongod实例，也可以是一个副本集。 生产环境下Shard一般是一个 Replica Set，以防止该数据片的单点故障。\n对于分片集合(sharded collection)来说，每个分片上都存储了集合的一部分数据(按照分片键切分)，如果集合没有分片，那么该集合的数据都存储在数据数据分片（Shards）\n分片用于存储真正的集群数据，可以是一个单独的 Mongod实例，也可以是一个副本集。 生产环境下Shard一般是一个 Replica Set，以防止该数据片的单点故障。\n对于分片集合(sharded collection)来说，每个分片上都存储了集合的一部分数据(按照分片键切分)，如果集合没有分片，那么该集合的数据都存储在数据数据分片（Shards）\n分片用于存储真正的集群数据，可以是一个单独的 Mongod实例，也可以是一个副本集。 生产环境下Shard一般是一个 Replica Set，以防止该数据片的单点故障。\n对于分片集合(sharded collection)来说，每个分片上都存储了集合的一部分数据(按照分片键切分)，如果集合没有分片，那么该集合的数据都存储在数据数据分片（Shards）\n分片用于存储真正的集群数据，可以是一个单独的 Mongod实例，也可以是一个副本集。 生产环境下Shard一般是一个 Replica Set，以防止该数据片的单点故障。\n对于分片集合(sharded collection)来说，每个分片上都存储了集合的一部分数据(按照分片键切分)，如果集合没有分片，那么该集合的数据都存储在数据",
			})
		}
		coll().InsertMany(context.Background(), users)
		select {
		case <-tk.C:
			fmt.Println(i)
		}
	}
}

type User struct {
	Username     string `bson:"username"`
	Password     string `bson:"password"`
	Description  string `bson:"description"`
	Description1 string `bson:"description1"`
	Description2 string `bson:"description2"`
	Description3 string `bson:"description3"`
	Description4 string `bson:"description4"`
	Description5 string `bson:"description5"`
	Description6 string `bson:"description6"`
	Description7 string `bson:"description7"`
}
