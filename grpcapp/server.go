package grpcapp

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/jnst/x-go/grpcapp/internal/gen/lol"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedLeagueOfLegendsServer
}

func (s Server) GetChampion(ctx context.Context, request *pb.GetChampionRequest) (*pb.GetChampionResponse, error) {
	panic("implement me")
}

func (s Server) ListChampions(ctx context.Context, request *pb.ListChampionsRequest) (*pb.ListChampionsResponse, error) {
	return &pb.ListChampionsResponse{
		Champions: []*pb.Champion{
			{
				ChampionId: 1,
				Type:       pb.Champion_MARKSMAN,
				Name:       "ASHE",
				Message:    "THE FROST ARCHER",
			},
			{
				ChampionId: 2,
				Type:       pb.Champion_MAGE,
				Name:       "BRAND",
				Message:    "THE BURNING VENGEANCE",
			},
			{
				ChampionId: 3,
				Type:       pb.Champion_ASSASSIN,
				Name:       "FIZZ",
				Message:    "THE TIDAL TRICKSTER",
			},
			{
				ChampionId: 4,
				Type:       pb.Champion_TANK,
				Name:       "THE MIGHT OF DEMACIA",
				Message:    "GAREN",
			},
			{
				ChampionId: 5,
				Type:       pb.Champion_FIGHTER,
				Name:       "IRELIA",
				Message:    "THE BLADE DANCER",
			},
			{
				ChampionId: 6,
				Type:       pb.Champion_SUPPORT,
				Name:       "SONA",
				Message:    "MAVEN OF THE STRINGS",
			},
		},
	}, nil
}

func (s Server) GetBattleField(ctx context.Context, request *pb.GetBattleFieldRequest) (*pb.GetBattleFieldResponse, error) {
	panic("implement me")
}

func StartServer() {
	const port = 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(measurementInterceptor()))
	pb.RegisterLeagueOfLegendsServer(s, Server{})
	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func measurementInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		now := time.Now()

		res, err := handler(ctx, req)
		if err != nil {
			return nil, err
		}

		log.Printf("[%s] request size: %s, response size: %s, elapsed time: %dms\n",
			info.FullMethod,
			binarySize(req),
			binarySize(res),
			time.Since(now).Milliseconds())

		return res, nil
	}
}
