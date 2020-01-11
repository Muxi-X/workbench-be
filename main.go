package main

import (
	"fmt"
	"log"

	opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"workbench-status-service/config"
	"workbench-status-service/model"
	"workbench-status-service/pkg/tracer"
	pb "workbench-status-service/proto"
	s "workbench-status-service/service"

	"github.com/micro/go-micro"
	"github.com/opentracing/opentracing-go"
)

// type repository interface {
// 	Create(*pb.Consignment) (*pb.Consignment, error)
// 	GetAll() []*pb.Consignment
// }

// // Repository - Dummy repository, this simulates the use of a datastore
// // of some kind. We'll replace this with a real implementation later on.
// type Repository struct {
// 	consignments []*pb.Consignment
// }

// func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
// 	updated := append(repo.consignments, consignment)
// 	repo.consignments = updated
// 	return consignment, nil
// }

// func (repo *Repository) GetAll() []*pb.Consignment {
// 	return repo.consignments
// }

// // Service should implement all of the methods to satisfy the service
// // we defined in our protobuf definition. You can check the interface
// // in the generated code itself for the exact method signatures etc
// // to give you a better idea.
// type service struct {
// 	repo repository
// }

// // CreateConsignment - we created just one method on our service,
// // which is a create method, which takes a context and a request as an
// // argument, these are handled by the gRPC server.
// func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {

// 	// Save our consignment
// 	consignment, err := s.repo.Create(req)
// 	if err != nil {
// 		return err
// 	}

// 	// Return matching the `Response` message we created in our
// 	// protobuf definition.
// 	res.Created = true
// 	res.Consignment = consignment
// 	return nil
// }

// func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
// 	consignments := s.repo.GetAll()
// 	res.Consignments = consignments
// 	return nil
// }

func main() {

	t, io, err := tracer.NewTracer("workbench.service.status", "localhost:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	// set var t to Global Tracer (opentracing single instance mode)
	opentracing.SetGlobalTracer(t)

	// init config
	if err := config.Init(""); err != nil {
		panic(err)
	}

	// init db
	model.DB.Init()
	defer model.DB.Close()

	// repo := &Repository{}

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("workbench.service.status"),
		micro.WrapHandler(
			opentracingWrapper.NewHandlerWrapper(opentracing.GlobalTracer()),
		),
		// micro.WrapClient(
		// 	opentracingWrapper.NewClientWrapper(tracer),
		// ),
	)

	// Init will parse the command line flags.
	srv.Init()

	// Register handler
	pb.RegisterStatusServiceHandler(srv.Server(), &s.StatusService{})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
