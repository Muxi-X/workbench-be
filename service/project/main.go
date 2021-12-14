package main

import (
	"fmt"
	m "github.com/Muxi-X/workbench-be/model"
	"log"
	mm "muxi-workbench-project/model"

	// mm "muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	s "muxi-workbench-project/service"
	"muxi-workbench/config"
	logger "muxi-workbench/log"
	"muxi-workbench/model"
	// m "muxi-workbench/model"
	"muxi-workbench/pkg/handler"
	tracer "muxi-workbench/pkg/tracer"

	_ "github.com/micro/go-plugins/registry/kubernetes"

	"github.com/micro/go-micro"
	opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
)

func init() {
	s.UserInit()
	s.AttentionInit()
	s.TeamInit()
}

func main() {
	// init config
	if err := config.Init("", "WORKBENCH_PROJECT"); err != nil {
		panic(err)
	}

	t, io, err := tracer.NewTracer(viper.GetString("local_name"), viper.GetString("tracing.jager"))
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	defer logger.SyncLogger()

	// set var t to Global Tracer (opentracing single instance mode)
	opentracing.SetGlobalTracer(t)

	// init db
	model.DB.Init()
	defer model.DB.Close()
	// {
	// 	pIDs := []uint32{2}
	// 	var record []*mm.SearchResult
	// 	query := m.DB.Self.
	// 		Raw("SELECT d.id, filename, create_time, u.name, content, p.name project_name FROM docs d "+
	// 			"LEFT JOIN users u ON u.id = d.editor_id "+
	// 			"LEFT JOIN projects p ON p.id = project_id "+
	// 			"WHERE project_id in (?) AND (d.filename like ? OR d.content like ?) "+
	// 			"UNION ALL SELECT f.id, filename, create_time, u.name, url, p.name project_name FROM files f "+
	// 			"LEFT JOIN users u ON u.id = f.creator_id "+
	// 			"LEFT JOIN projects p ON p.id = project_id "+
	// 			"WHERE project_id in (?) AND f.filename like ?", pIDs, "%1%", "%1%", pIDs, "%1%")
	// 	if err := query.Scan(&record).Error; err != nil {
	// 		panic(err)
	// 	}
	// 	// db := m.DB.Self
	// 	// db.Raw("? UNION ?",
	// 	// 	db.Raw("SELECT d.id, filename, create_time, u.name, content, p.name project_name FROM docs d "+
	// 	// 		"LEFT JOIN users u ON u.id = d.editor_id "+
	// 	// 		"LEFT JOIN projects p ON p.id = project_id "+
	// 	// 		"WHERE project_id in (?) AND d.filename like ? ", pIDs, "%1%"),
	// 	// 	db.Raw("SELECT f.id, filename, create_time, u.name, url, p.name project_name FROM files f "+
	// 	// 		"LEFT JOIN users u ON u.id = f.creator_id "+
	// 	// 		"LEFT JOIN projects p ON p.id = project_id "+
	// 	// 		"WHERE project_id in (?) AND f.filename like ?", pIDs, "%1%"),
	// 	// ).Scan(&record)
	//
	// 	fmt.Printf("%+v\n", record[0])
	// 	fmt.Println()
	// 	fmt.Println()
	// 	fmt.Println()
	// 	fmt.Printf("%+v\n", record[1])
	// 	// for _, r := range record {
	// 	// if r.Content != "" {
	// 	// 	fmt.Println(i)
	// 	// 	continue
	// 	// } else {
	// 	// 	r.Type = 2
	// 	// }
	// 	// 	fmt.Printf("%+v\n", r)
	//
	// 	return
	// }
	// init redis
	model.RedisDB.Init()
	defer model.RedisDB.Close()

	// 同步 redis
	if err := s.SynchronizeTrashbinToRedis(); err != nil {
		log.Fatal(err)
	}

	// 定时任务
	go s.GoTidyTrashbin(model.DB.Self)

	srv := micro.NewService(
		micro.Name(viper.GetString("local_name")),
		micro.WrapHandler(
			opentracingWrapper.NewHandlerWrapper(opentracing.GlobalTracer()),
		),
		micro.WrapHandler(handler.ServerErrorHandlerWrapper()),
	)

	// Init will parse the command line flags.
	srv.Init()

	// Register handler
	pb.RegisterProjectServiceHandler(srv.Server(), &s.Service{})

	// Run the server
	if err := srv.Run(); err != nil {
		logger.Error(err.Error())
	}
}
