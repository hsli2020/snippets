service GoReleaseService {
    rpc GetReleaseInfo(GetReleaseInfoRequest) returns (ReleaseInfo) {}
    rpc ListReleases(ListReleasesRequest) returns (ListReleasesResponse) {}
}

message GetReleaseInfoRequest {
    string version = 1;
}

message ListReleasesRequest {} //empty

message ListReleasesResponse {
    repeated ReleaseInfo releases = 1;
}

message ReleaseInfo {
    string version = 1;
    string release_date = 2;
    string release_notes_url = 3;
}

////////////////////////////////////////////////////////////////////////////////

type releaseInfo struct {
    ReleaseDate     string `json:"release_date"`
    ReleaseNotesURL string `json:"release_notes_url"`
}

/* goReleaseService implements GoReleaseServiceServer as defined in the generated code:
// Server API for GoReleaseService service
type GoReleaseServiceServer interface {
    GetReleaseInfo(context.Context, *GetReleaseInfoRequest) (*ReleaseInfo, error)
    ListReleases(context.Context, *ListReleasesRequest) (*ListReleasesResponse, error)
}
*/
type goReleaseService struct {
    releases map[string]releaseInfo
}

func (g *goReleaseService) GetReleaseInfo(ctx context.Context,
		r *pb.GetReleaseInfoRequest) (*pb.ReleaseInfo, error) {

    // lookup release info for version supplied in request
    ri, ok := g.releases[r.GetVersion()]
    if !ok {
        return nil, status.Errorf(codes.NotFound, "release verions %s not found", r.GetVersion())
    }

    // success
    return &pb.ReleaseInfo{
        Version:         r.GetVersion(),
        ReleaseDate:     ri.ReleaseDate,
        ReleaseNotesUrl: ri.ReleaseNotesURL,
    }, nil
}

func (g *goReleaseService) ListReleases(ctx context.Context, r *pb.ListReleasesRequest) (*pb.ListReleasesResponse, error) {
    var releases []*pb.ReleaseInfo

    // build slice with all the available releases
    for k, v := range g.releases {
        ri := &pb.ReleaseInfo{
            Version:         k,
            ReleaseDate:     v.ReleaseDate,
            ReleaseNotesUrl: v.ReleaseNotesURL,
        }

        releases = append(releases, ri)
    }

    return &pb.ListReleasesResponse{ Releases: releases }, nil
}

func main() {
    // code redacted
    lis, err := net.Listen("tcp", *listenPort)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    log.Println("Listening on ", *listenPort)
    server := grpc.NewServer()

    pb.RegisterGoReleaseServiceServer(server, svc)

    if err := server.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

////////////////////////////////////////////////////////////////////////////////

func main() {
    flag.Parse()

    conn, err := grpc.Dial(*target, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("grpc.Dial err: %v", err)
    }

    client := pb.NewGoReleaseServiceClient(conn)

    ctx := context.Background()
    rsp, err := client.ListReleases(ctx, &pb.ListReleasesRequest{})

    if err != nil {
        log.Fatalf("ListReleases err: %v", err)
    }

    releases := rsp.GetReleases()
    if len(releases) > 0 {
        sort.Sort(byVersion(releases))
        fmt.Printf("Version\tRelease Date\tRelease Notes\n")
    } else {
        fmt.Println("No releases found")
    }

    for _, ri := range releases {
        fmt.Printf("%s\t%s\t%s\n",
            ri.GetVersion(),
            ri.GetReleaseDate(),
            ri.GetReleaseNotesURL())
    }
}
