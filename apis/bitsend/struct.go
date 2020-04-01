package bitsend

import (
	"github.com/cheggaaa/pb/v3"
)

type wssOptions struct {
	token     string
	parallel  int
	interval  int
	prefix    string
	forceMode bool
	debugMode bool
	passCode  string
}

type uploadConfig struct {
	debug      bool
	fileName   string
	fileReader *pb.Reader
	fileSize   int64
}

type uploadResp struct {
	Files []uploadRespBlock `json:"files"`
}

type uploadRespBlock struct {
	Name       string `json:"name"`
	Size       int    `json:"size"`
	Type       string `json:"type"`
	NiceSize   string `json:"niceSize"`
	RealName   string `json:"realName"`
	FileKey    string `json:"fileKey"`
	DelFileKey string `json:"delFileKey"`
	DeleteUrl  string `json:"delete_url"`
	DeleteType string `json:"delete_type"`
}
