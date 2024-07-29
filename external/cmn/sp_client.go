package cmn

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/bnb-chain/greenfield-bundle-sdk/bundle"
	"greeenfield-bsc-archiver/models"
)

type SPClient struct {
	hc   *http.Client
	host string
}

func NewSPClient(host string) (*SPClient, error) {
	transport := &http.Transport{
		DisableCompression:  true,
		MaxIdleConnsPerHost: 1000,
		MaxConnsPerHost:     1000,
		IdleConnTimeout:     90 * time.Second,
	}
	client := &http.Client{
		Timeout:   10 * time.Minute,
		Transport: transport,
	}
	return &SPClient{hc: client, host: host}, nil
}

func (c *SPClient) GetBucketReadQuota(ctx context.Context, bucketName string) (QuotaInfo, error) {
	year, month, _ := time.Now().Date()
	var date string
	if int(month) < 10 {
		date = strconv.Itoa(year) + "-" + "0" + strconv.Itoa(int(month))
	} else {
		date = strconv.Itoa(year) + "-" + strconv.Itoa(int(month))
	}
	var urlStr string
	parts := strings.Split(c.host, "//")
	urlStr = parts[0] + "//" + bucketName + "." + parts[1] + "/"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlStr, nil)
	if err != nil {
		return QuotaInfo{}, err
	}
	// set query parameters
	q := req.URL.Query()
	q.Add("read-quota", "")
	q.Add("year-month", date)
	req.URL.RawQuery = q.Encode()
	resp, err := c.hc.Do(req)
	if err != nil {
		return QuotaInfo{}, err
	}
	defer resp.Body.Close()
	QuotaResult := QuotaInfo{}
	err = xml.NewDecoder(resp.Body).Decode(&QuotaResult)
	if err != nil {
		return QuotaInfo{}, err
	}
	return QuotaResult, nil
}

func (c *SPClient) GetBundleObject(ctx context.Context, bucketName, objectName string) (io.ReadCloser, error) {
	var urlStr string
	parts := strings.Split(c.host, "//")
	urlStr = parts[0] + "//" + bucketName + "." + parts[1] + "/" + objectName

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (c *SPClient) GetBundleBlocks(ctx context.Context, bucketName, objectName string) ([]*models.Block, error) {
	var urlStr string
	parts := strings.Split(c.host, "//")
	urlStr = parts[0] + "//" + bucketName + "." + parts[1] + "/" + objectName

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	tempFile, err := os.CreateTemp("", "bundle")
	if err != nil {
		fmt.Printf("Failed to create temporary file: %v\n", err)
		return nil, err
	}
	defer os.Remove(tempFile.Name())

	// Write the content to the temporary file
	_, err = tempFile.Write(body)
	if err != nil {
		fmt.Printf("Failed to write downloaded bundle to file: %v\n", err)
		return nil, err
	}
	defer tempFile.Close()

	bundleObjects, err := bundle.NewBundleFromFile(tempFile.Name())
	if err != nil {
		fmt.Printf("Failed to create bundle from file: %v\n", err)
		return nil, err
	}
	var blocksInfo []*models.Block
	for _, objMeta := range bundleObjects.GetBundleObjectsMeta() {
		objFile, _, err := bundleObjects.GetObject(objMeta.Name)
		if err != nil {
			return nil, err
		}

		var objectInfo []byte
		objectInfo, err = io.ReadAll(objFile)
		if err != nil {
			objFile.Close()
			return nil, err
		}
		objFile.Close()

		var blockInfo *models.Block
		err = json.Unmarshal(objectInfo, &blockInfo)
		if err != nil {
			return nil, err
		}
		blocksInfo = append(blocksInfo, blockInfo)
	}

	return blocksInfo, nil
}
