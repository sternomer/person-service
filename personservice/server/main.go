package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	pb "github.com/sternomer/go-grpc-mongodb-master/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PersonServiceServer struct {
	pb.UnimplementedPersonServiceServer
}

func (s *PersonServiceServer) ReadPerson(ctx context.Context, req *pb.ReadPersonReq) (*pb.ReadPersonRes, error) {

	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	result := persondb.FindOne(ctx, bson.M{"_id": oid})

	data := PersonItem{}

	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find Person with Object Id %s: %v", req.GetId(), err))
	}

	response := &pb.ReadPersonRes{
		Person: &pb.Person{
			Id:        oid.Hex(),
			Birthdate: data.BirthDate,
			FirstName: data.FirstName,
			LastName:  data.LastName,
		},
	}
	return response, nil
}

func (s *PersonServiceServer) CreatePerson(ctx context.Context, req *pb.CreatePersonReq) (*pb.CreatePersonRes, error) {

	person := req.GetPerson()

	data := PersonItem{

		BirthDate: person.GetBirthdate(),
		FirstName: person.GetFirstName(),
		LastName:  person.GetLastName(),
	}

	result, err := persondb.InsertOne(mongoCtx, data)

	if err != nil {

		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	oid := result.InsertedID.(primitive.ObjectID)
	person.Id = oid.Hex()

	return &pb.CreatePersonRes{Person: person}, nil
}

func (s *PersonServiceServer) UpdatePerson(ctx context.Context, req *pb.UpdatePersonReq) (*pb.UpdatePersonRes, error) {

	person := req.GetPerson()

	oid, err := primitive.ObjectIDFromHex(person.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Could notupdate current person bacuse id not found: %v", err),
		)
	}

	update := bson.M{
		"birthdate": person.GetBirthdate(),
		"firstName": person.GetFirstName(),
		"lastName":  person.GetLastName(),
	}

	filter := bson.M{"_id": oid}

	result := persondb.FindOneAndUpdate(ctx, filter, bson.M{"$set": update}, options.FindOneAndUpdate().SetReturnDocument(1))

	decoded := PersonItem{}
	err = result.Decode(&decoded)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Could not find Person with supplied ID: %v", err),
		)
	}
	return &pb.UpdatePersonRes{
		Person: &pb.Person{
			Id:        decoded.ID.Hex(),
			Birthdate: decoded.BirthDate,
			FirstName: decoded.FirstName,
			LastName:  decoded.LastName,
		},
	}, nil
}

func (s *PersonServiceServer) DeletePerson(ctx context.Context, req *pb.DeletePersonReq) ( *pb.DeletePersonRes, error) {
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}

	_, err = persondb.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find/delete person with id %s: %v", req.GetId(), err))
	}
	return &pb.DeletePersonRes{
		Success: true,
	}, nil
}

func (s *PersonServiceServer) ListPersons(req *pb.ListPersonsReq, stream pb.PersonService_ListPersonsServer) error {

	data := &PersonItem{}

	cursor, err := persondb.Find(context.Background(), bson.M{})
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {

		err := cursor.Decode(data)

		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}

		stream.Send(&pb.ListPersonsRes{
			Person: &pb.Person{
				Id:        data.ID.Hex(),
				Birthdate: data.BirthDate,
				FirstName: data.FirstName,
				LastName:  data.LastName,
			},
		})
	}

	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unkown cursor error: %v", err))
	}
	return nil
}

type PersonItem struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	BirthDate string             `bson:"birthdate"`
	FirstName string             `bson:"firstName"`
	LastName  string             `bson:"lastName"`
}

var db *mongo.Client
var persondb *mongo.Collection
var mongoCtx context.Context

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Starting server on port :50051...")

	listener, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Unable to listen on port :50051: %v", err)
	}

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)

	srv := &PersonServiceServer{}

	pb.RegisterPersonServiceServer(s, srv)

	fmt.Println("Connecting to MongoDB...")
	mongoCtx = context.Background()
	db, err = mongo.Connect(mongoCtx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping(mongoCtx, nil)
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v\n", err)
	} else {
		fmt.Println("Connected to Mongodb")
	}

	persondb = db.Database("mydb").Collection("person")

	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	fmt.Println("Server succesfully started on port :50051")

	c := make(chan os.Signal)

	signal.Notify(c, os.Interrupt)

	<-c

	fmt.Println("\nStopping the server...")
	s.Stop()
	listener.Close()
	fmt.Println("Closing MongoDB connection")
	db.Disconnect(mongoCtx)
	fmt.Println("Done.")
}
