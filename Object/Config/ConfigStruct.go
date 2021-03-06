package Config

import (
	"github.com/minio/minio-go/v7"
	"time"
)

type Yaml struct {
	ServerList []Server `yaml:"server"`
}
type Server struct {
	Name            string `yaml:"name"`
	ListenPort      int    `yaml:"listenPort"`
	Enable          bool   `yaml:"enable"`
	Host            string `yaml:"host"`
	AccessKeyID     string `yaml:"accessKeyID"`
	SecretAccessKey string `yaml:"secretAccessKey"`
	Bucket          string `yaml:"bucket"`
	Options         struct {
		UseSSL           bool                   `yaml:"useSSL"`
		Region           string                 `yaml:"region"`
		BucketLookupType minio.BucketLookupType `yaml:"bucketLookupType"` //DNS,Path:1,Auto:0
	} `yaml:"options"`
}

type ObjectInfo struct {
	Key          string
	Size         int64
	ETag         string
	LsatModified time.Time
}

type ObjectInfoList struct {
	Info []ObjectInfo
}
