package cowtransfer

import (
	cmap "github.com/orcaman/concurrent-map"
	"net/http"
	"sync"
	"time"
)

type requestConfig struct {
	debug    bool
	retry    int
	timeout  time.Duration
	modifier func(r *http.Request)
}

type uploadPart struct {
	content []byte
	count   int64
}

type uploadConfig struct {
	wg      *sync.WaitGroup
	token   string
	hashMap *cmap.ConcurrentMap
}

type cowOptions struct {
	Parallel   int
	token      string
	interval   int
	singleMode bool
	blockSize  int
	hashCheck  bool
	passCode   string
	DebugMode  bool
}

type prepareSendResp struct {
	UploadToken  string `json:"uptoken"`
	TransferGUID string `json:"transferguid"`
	UniqueURL    string `json:"uniqueurl"`
	Prefix       string `json:"prefix"`
	QRCode       string `json:"qrcode"`
	Error        bool   `json:"error"`
	ErrorMessage string `json:"error_message"`
}

type uploadResponse struct {
	Ticket string `json:"ctx"`
	Hash   int    `json:"crc32"`
}

type finishResponse struct {
	TempDownloadCode string `json:"tempDownloadCode"`
	Status           bool   `json:"complete"`
}

type downloadDetailsResponse struct {
	GUID         string                 `json:"guid"`
	DownloadName string                 `json:"downloadName"`
	Deleted      bool                   `json:"deleted"`
	Uploaded     bool                   `json:"uploaded"`
	Details      []downloadDetailsBlock `json:"transferFileDtos"`
}

type downloadDetailsBlock struct {
	GUID     string `json:"guid"`
	FileName string `json:"fileName"`
	Size     string `json:"size"`
}

type downloadConfigResponse struct {
	Link string `json:"link"`
}