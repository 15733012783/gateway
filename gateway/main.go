package main

import (
	"context"
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"

	gw "test/example/gen"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:8050", "gRPC server endpoint")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gw.RegisterMyServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", mux)
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		grpclog.Fatal(err)
	}
}

//func Updated(stu DujiaweiReq) {
//	res := httplib.Post(ESURL + "dujiawei" + "/_doc/" + "3")
//
//	res.JSONBody(stu)
//
//	dawei := DujiaweiRes{}
//	s, err := res.String()
//	if err != nil {
//		logs.Info(err)
//		return
//	}
//
//	json.Unmarshal([]byte(s), &dawei)
//	fmt.Println(dawei)
//
//#添加数据
//POST /fitness/_doc/1
//{
//"id": 2,
//"name":"高伟明",
//"price": 200.30,
//"times":500,
//"pro_level":0,
//"img":"www.baidu.com",
//"type":1
//}
//
//POST /fitness/_search
//{
//"query": {
//"match": {
//"name": "高"
//}
//},
//"from": 0,
//"size": 1
//}
//
//#不同商品类别下的商品数量
//POST /google_bd/_search
//{
//"size": 0,
//"aggs": {
//"categories": {
//"terms": {
//"field": "id",
//"size": 10, // 返回前10个类别
//"order": {
//"_count": "desc" // 按数量降序排列
//}
//}
//}
//}
//}
//#删除数据
//DELETE /dujiawei/_doc/2
//#查询(高亮+分页)
//POST /new_gitee_dir/_search
//{
//"query": {
//"multi_match": {
//"query": "go-zero",
//"fields": [
//"title",
//"readme"
//]
//}
//},
//"from": 0,
//"size": 10,
//"highlight": {
//"fields": {
//"readme": {
//"pre_tags": "<a color='red'>",
//"post_tags": "</a>"
//}
//}
//}
//}
