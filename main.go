package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/tv2169145/cart/domain/repository"
	"github.com/tv2169145/cart/domain/service"
	"github.com/tv2169145/cart/handler"
	cart "github.com/tv2169145/cart/proto/cart"
	"github.com/tv2169145/common"
)

var QPS = 100 // 限流數(每秒幾個request)

func main() {
	// 配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "micro/config")
	if err != nil {
		log.Error(err)
	}

	// 註冊中心
	consulRegister := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:8500"}
	})

	// 鏈路追蹤 jaeger
	t, io, err := common.NewTracer("go.micro.service.cart", "localhost:6831")
	if err != nil {
		log.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// mysql
	mysqlConfig := common.GetMysqlFromConsul(consulConfig, "mysql")
	mysqlConnectionInfo := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		mysqlConfig.User,
		mysqlConfig.Pwd,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.Database,
		mysqlConfig.Charset,
	)
	fmt.Println("mysql connection:", mysqlConnectionInfo)

	// db初始化
	db, err := gorm.Open("mysql", mysqlConnectionInfo)
	if err != nil {
		log.Error(err)
	}
	defer db.Close()
	// 禁止副表
	db.SingularTable(true)

	// micro service
	microService := micro.NewService(
		micro.Name("go.micro.service.cart"),
		micro.Version("latest"),
		// 暴露的服務端口
		micro.Address("127.0.0.1:8084"),
		// 註冊中心
		micro.Registry(consulRegister),
		// 鏈路追蹤
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		// 添加限流
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
	)
	microService.Init() // 初始化

	// cartRepository
	cartRepository := repository.NewCartRepository(db)

	// 只跑一次 建表
	//if err := cartRepository.InitTable(); err != nil {
	//	log.Error(err)
	//}

	// cartService
	cartDataService := service.NewCartDataService(cartRepository)

	// handler
	err = cart.RegisterCartHandler(microService.Server(), &handler.Cart{CartDateService: cartDataService})
	if err != nil {
		log.Error(err)
	}

	// Run
	if err := microService.Run(); err != nil {
		log.Fatal(err)
	}
}
