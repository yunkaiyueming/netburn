package hfile

import (
	"fmt"

	. "github.com/yunkaiyueming/netburn/g"
)

type HfileServer struct{}

type BucketItem map[string][]HFile

var defaultBucketGit = BucketItem{
	"img": []HFile{
		{"img_abcd1", "abcd1", 25, "2006-02-26 23:23:46", "ab"},
		{"img_abcd2", "abcd2", 25, "2006-02-26 23:23:46", "cd"},
		{"img_abcd2", "abcd2", 25, "2006-02-26 23:23:46", "ab"},
	},
	"ios": []HFile{
		{"ios_abcd1", "abcd1", 25, "2006-02-26 23:23:46", "ab"},
	},
	"tab": []HFile{
		{"tab_abcd1", "abcd1", 25, "2006-02-26 23:23:46", "ab"},
	},
	"android": {
		{"android_abcd1", "abcd1", 25, "2006-02-26 23:23:46", "ab"},
		{"android_abcd2", "abcd2", 25, "2006-02-26 23:23:46", "ab"},
	},
}

func (h *HfileServer) GetBucketFiles(bucketName string, output *[]HFile) error {
	if ret, ok := defaultBucketGit[bucketName]; ok {
		*output = ret
	} else {
		*output = nil
	}
	return nil
}

func (h *HfileServer) GetBucketNames(input interface{}, output *[]string) error {
	for name, _ := range defaultBucketGit {
		*output = append(*output, name)
	}
	return nil
}

func (h *HfileServer) GetAllBucketsFiles(input interface{}, output *interface{}) error {
	*output = defaultBucketGit
	return nil
}

func (h *HfileServer) GetFileMsg(input map[string]string, output *interface{}) error {
	items := defaultBucketGit[input["bucket_name"]]
	fmt.Println(input)
	for _, item := range items {
		if item.Key == input["key"] {
			*output = item
			break
		}
	}
	return nil
}
