package nim

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	BaseUrl = "https://api.netease.im/nimserver"
)

var jsonTool = jsoniter.ConfigCompatibleWithStandardLibrary

type ImClient struct {
	AppKey    string
	AppSecret string
	Nonce     string
	mutex     *sync.Mutex
	client    *resty.Client
}

//CreateImClient  创建im客户端
func CreateImClient(appKey string, appSecret string) *ImClient {
	c := &ImClient{AppKey: appKey, AppSecret: appSecret, Nonce: uuid.NewString(), mutex: new(sync.Mutex)}
	c.client = resty.New()
	c.client.SetHeader("Accept", "application/json;charset=utf-8")
	c.client.SetHeader("Content-Type", "application/x-www-form-urlencoded;charset=utf-8;")
	c.client.SetHeader("AppKey", c.AppKey)
	c.client.SetHeader("Nonce", c.Nonce)
	return c
}

func (c *ImClient) setCommonHead(req *resty.Request) {
	c.mutex.Lock() //多线程并发访问map导致panic
	defer c.mutex.Unlock()
	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	req.SetHeader("CurTime", timeStamp)
	req.SetHeader("CheckSum", getCheckSum(c.AppSecret, c.Nonce, timeStamp))
}

func getCheckSum(appSecret string, Noce string, curTime string) string {
	hasher := sha1.New()
	hasher.Write([]byte(appSecret + Noce + curTime))
	return strings.ToLower(hex.EncodeToString(hasher.Sum(nil)))
}
