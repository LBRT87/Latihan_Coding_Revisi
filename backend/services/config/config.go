package config

import (
	"context"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Config struct {
	GRPCPORT  string
	HTTPPORT  string
	DBHOST    string
	DBUSER    string
	DBNAME    string
	DBPORT    string
	DBPASS    string
	REDISHOST string
	REDISPASS string
	REDISPORT string
	JWTSECRET string
	S3ENDPOINT string
	PUBLICSEAWEEDENDPOINT string
	ACCESSKEYID string
	SECRETACCESSKEY string
	BUCKETNAME string
}

func Load() *Config {
	return &Config{
		GRPCPORT:  GetEnv("GRPC_PORT", "9001"),
		HTTPPORT:  GetEnv("HTTP_PORT", "8001"),
		DBUSER:    GetEnv("DB_USER", ""),
		DBHOST:    GetEnv("DB_HOST", ""),
		DBNAME:    GetEnv("DB_NAME", ""),
		DBPORT:    GetEnv("DB_PORT", ""),
		DBPASS:    GetEnv("DB_PASS", ""),
		REDISHOST: GetEnv("REDIS_HOST", ""),
		REDISPASS: GetEnv("REDIS_PASS", ""),
		REDISPORT: GetEnv("REDIS_PORT", ""),
		JWTSECRET: GetEnv("JWT_SECRET", ""),
		S3ENDPOINT: GetEnv("S3_ENPOINT", "seaweedfs:8333"),
		PUBLICSEAWEEDENDPOINT: GetEnv("S3_PUBLIC_ENDPOINT", "localhost:8333"),
		ACCESSKEYID: GetEnv("ACCESS_KEY_ID", "tpa_onsite"),
		SECRETACCESSKEY : GetEnv("SECRET_ACCESS_KEY", "tpa_onsite"),
		BUCKETNAME: GetEnv("BUCKET_NAME", "eskrim"),
	}
}

func (c *Config) NewSeaweedClient() *minio.Client {
	client, err := minio.New(c.S3ENDPOINT, &minio.Options{
		Secure: false,
		Creds: credentials.NewStaticV4(c.ACCESSKEYID, c.SECRETACCESSKEY, ""),
	})

	if err != nil {
		log.Fatalf("error while establihing connection to seaweed  : %v", err)
	}

	return client
}

func (c *Config) InitSeaweed(client *minio.Client) {

	ctx := context.Background()

	ok, err := client.BucketExists(ctx, c.BUCKETNAME)

	if err != nil {
		log.Fatalf("error while pinging the seaweed : %v", err)
	}

	if !ok {
		err := client.MakeBucket(ctx, c.BUCKETNAME, minio.MakeBucketOptions{})
		if err != nil {
			log.Fatalf("error while making the seaweed bucket : %v", err)
		}
	}
}

func GetEnv(code, fallback string) string {
	if cfg := os.Getenv(code); cfg != "" {
		return cfg
	}
	return fallback
}