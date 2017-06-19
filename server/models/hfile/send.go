package hfile

import (
	. "github.com/yunkaiyueming/netburn/g"
)

type HfielServer struct{}

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

func (h *HfielServer) GetBucketFiles(bucketName string, items *[]HFile) error {
	if ret, ok := defaultBucketGit[bucketName]; ok {
		*items = ret
	} else {
		*items = nil
	}
	return nil
}

func (h *HfielServer) GetBucketNames(input interface{}, output *[]string) error {
	for name, _ := range defaultBucketGit {
		*output = append(*output, name)
	}
	return nil
}

func (h *HfielServer) GetAllBucketsFiles(input interface{}, output *interface{}) error {
	*output = defaultBucketGit
	return nil
}

func (h *HfielServer) GetFileMsg(input map[string]string, output *interface{}) error {
	items := defaultBucketGit[input["bucket_name"]]
	for _, item := range items {
		if item.Key == input["key"] {
			*output = item
			break
		}
	}
	return nil
}
