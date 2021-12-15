// Unary RPC

// Server

type SearchService struct{}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
    return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

const PORT = "9001"

func main() {
    server := grpc.NewServer()
    pb.RegisterSearchServiceServer(server, &SearchService{})

    lis, err := net.Listen("tcp", ":"+PORT)
    ...

    server.Serve(lis)
}

// Client

func main() {
    conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
    ...
    defer conn.Close()

    client := pb.NewSearchServiceClient(conn)
    resp, err := client.Search(context.Background(), &pb.SearchRequest{
        Request: "gRPC",
    })
    ...
}

// Server-side streaming RPC

// Server

func (s *StreamService) List(r *pb.StreamRequest, stream pb.StreamService_ListServer) error {
    for n := 0; n <= 6; n++ {
        stream.Send(&pb.StreamResponse{
            Pt: &pb.StreamPoint{
                ...
            },
        })
    }

    return nil
}

// Client

func printLists(client pb.StreamServiceClient, r *pb.StreamRequest) error {
    stream, err := client.List(context.Background(), r)
    ...

    for {
        resp, err := stream.Recv()
        if err == io.EOF {
            break
        }
        ...
    }

    return nil
}

// Client-side streaming RPC

// Server

func (s *StreamService) Record(stream pb.StreamService_RecordServer) error {
    for {
        r, err := stream.Recv()
        if err == io.EOF {
            return stream.SendAndClose(&pb.StreamResponse{Pt: &pb.StreamPoint{...}})
        }
        ...

    }

    return nil
}

// Client

func printRecord(client pb.StreamServiceClient, r *pb.StreamRequest) error {
    stream, err := client.Record(context.Background())
    ...

    for n := 0; n < 6; n++ {
        stream.Send(r)
    }

    resp, err := stream.CloseAndRecv()
    ...

    return nil
}

// Bidirectional streaming RPC

// Server

func (s *StreamService) Route(stream pb.StreamService_RouteServer) error {
    for {
        stream.Send(&pb.StreamResponse{...})
        r, err := stream.Recv()
        if err == io.EOF {
            return nil
        }
        ...
    }

    return nil
}

// Client

func printRoute(client pb.StreamServiceClient, r *pb.StreamRequest) error {
    stream, err := client.Route(context.Background())
    ...

    for n := 0; n <= 6; n++ {
        stream.Send(r)
        resp, err := stream.Recv()
        if err == io.EOF {
            break
        }
        ...
    }

    stream.CloseSend()

    return nil
}
