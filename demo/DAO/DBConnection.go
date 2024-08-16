package DAO

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var mongo_cnn *mongo.Client

func InitMongo() {
	// 设置MongoDB连接字符串
	connectionString := "mongodb://10.0.1.108:27017"
	// 创建客户端选项
	opts := options.Client().ApplyURI(connectionString).SetMaxPoolSize(50)
	// 创建一个新的客户端并连接到服务器
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")
	// ... 进行数据库操作 ...

	mongo_cnn = client
}
