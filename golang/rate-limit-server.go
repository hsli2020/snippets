//go:generate protoc -I ../token-bucket --go_out=plugins=grpc:../token-bucket ../token-bucket/token-bucket.proto
// https://github.com/sam09/rate-limiter
// Package main implements a server for TokenBucket service.
package main

import (
	"container/list"
	"context"
	"errors"
	"log"
	"net"
	"time"

	pb "github.com/sam09/rate-limiter/token-bucket"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type tokenBucket struct {
	name         string
	refillTime   int64
	refillAmount int64
	maxAmount    int64
	value        int64
	lastUpdated  time.Time
	tokens       *list.List
	maxID        int64
}

func newTokenBucket(name string, maxAmount int64, refillAmount int64, refillTime int64) *tokenBucket {
	bucket := tokenBucket{name: name, maxAmount: maxAmount, value: 0,
		refillAmount: refillAmount, refillTime: refillTime, lastUpdated: time.Now()}
	bucket.maxID = 0
	bucket.tokens = list.New()
	return &bucket
}

// rateLimitServer is used to implement TokenBucket.
type rateLimitServer struct {
	pb.UnimplementedTokenBucketServer
}

var buckets = make(map[string]*tokenBucket)

func getNewToken(value int64) int64 {
	return value + 1
}

func addToBucket(bucket *tokenBucket) error {
	var err error
	if bucket.value < bucket.maxAmount {
		token := getNewToken(bucket.maxID)
		bucket.tokens.PushBack(token)
		bucket.value++
		bucket.maxID++
		err = nil
		bucket.lastUpdated = time.Now()
	} else {
		log.Printf("Bucket Full")
		err = errors.New("Bucket Full")
	}
	return err
}

func removeFromBucket(bucket *tokenBucket) (int64, error) {
	var err error
	var id int64
	if bucket.value > 0 {
		tokenCounter := bucket.tokens.Back()
		bucket.tokens.Remove(tokenCounter)
		bucket.value--
		err = nil
		id = tokenCounter.Value.(int64)
		bucket.lastUpdated = time.Now()
	} else {
		log.Printf("No tokens available")
		err = errors.New("No tokens available")
	}
	return id, err
}

func (s *rateLimitServer) CreateBucket(ctx context.Context,
	in *pb.CreateBucketRequest) (*pb.CreateBucketResponse, error) {
	bucket := newTokenBucket(in.GetName(), in.GetMaxAmount(), in.GetRefillAmount(), in.GetRefillTime())
	buckets[bucket.name] = bucket
	log.Printf("Created bucket with name %v", bucket.name)
	return &pb.CreateBucketResponse{BucketName: bucket.name}, nil
}

func (s *rateLimitServer) AddToken(ctx context.Context, in *pb.AddTokenRequest) (*pb.AddTokenResponse, error) {
	log.Printf("Received request to add token: %v", in.GetBucketName())
	bucket := buckets[in.GetBucketName()]
	err := addToBucket(bucket)
	return &pb.AddTokenResponse{}, err
}

func (s *rateLimitServer) ConsumeToken(ctx context.Context, in *pb.ConsumeTokenRequest) (*pb.ConsumeTokenResponse, error) {
	log.Printf("Received request to consume token: %v", in.GetBucketName())
	bucket := buckets[in.GetBucketName()]
	var token *pb.Token
	id, err := removeFromBucket(bucket)
	if err == nil {
		token.Id = id
	}
	return &pb.ConsumeTokenResponse{Token: token}, err
}

func (s *rateLimitServer) Refill(ctx context.Context, in *pb.RefillTokenRequest) (*pb.RefillTokenResponse, error) {
	log.Printf("Received request to consume token: %v", in.GetBucketName())
	bucket := buckets[in.GetBucketName()]
	refillCount := bucket.refillAmount * (time.Now().Unix() - bucket.lastUpdated.Unix()) / bucket.refillTime
	if refillCount+bucket.value > bucket.maxAmount {
		refillCount = bucket.maxAmount - bucket.value
	}
	var err error
	for refillCount > 0 {
		err = addToBucket(bucket)
		if err != nil {
			break
		}
		refillCount--
	}
	return &pb.RefillTokenResponse{}, err
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTokenBucketServer(s, &rateLimitServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
