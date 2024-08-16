package DAO

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type BaseDAO struct {
	Data interface{}
	Uid  int
}

// 创建时 查询 数据构造原始数据代理
func (a *BaseDAO) BaseDAO() {
	// 选择数据库和集合
	db := mongo_cnn.Database("dnf_server_zj")
	collection := db.Collection("roles")
	// 查询数据
	var result bson.M
	collection.FindOne(context.TODO(), bson.M{"uid": a.Uid}).Decode(&result)

	a.Data = result
}
