package main

import (
	"google.golang.org/grpc"
	"log"
	pb "moviesapp.com/grpc/protos"
	"net"
)

const (
	port = ":50051"
)

var movies []*pb.MovieInfo

type movieServer struct {
	pb.UnimplementedMovieServer
}

func main() {
	initMovies()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to Listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterMovieServer(s, &movieServer{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to Listen: %v", err)
	}
}

func initMovies() {
	movie1 := &pb.MovieInfo{Id: "1", Isbn: "00251452", Title: "The Supermen", Director: &pb.Director{Firstname: "John", Lastname: "Doe"}}
	movie2 := &pb.MovieInfo{Id: "2", Isbn: "00212233", Title: "The Big Pig", Director: &pb.Director{Firstname: "Voi", Lastname: "TV"}}
	movie3 := &pb.MovieInfo{Id: "3", Isbn: "01231453", Title: "The Little Star", Director: &pb.Director{Firstname: "Poo", Lastname: "Doo"}}

	movies = append(movies, movie1)
	movies = append(movies, movie2)
	movies = append(movies, movie3)
}

func (s *movieServer) GetMovies(in *pb.Empty, stream pb.Movie_GetMoviesServer) error {
	log.Printf("Received: %v", in)
	for _, movie := range movies {
		if err := stream.Send(movie); err != nil {
			return err
		}
	}
	return nil
}

//func (s *movieServer) GetMovie(ctx context.Context,
//	in *pb.Id) (*pb.MovieInfo, error) {
//	log.Printf("Received: %v", in)
//
//	res := &pb.MovieInfo{}
//
//	for _, movie := range movies {
//		if strconv.Itoa(movie.GetId()) == in.GetValue() {
//			res = movie
//			break
//		}
//	}
//
//	return res, nil
//}
