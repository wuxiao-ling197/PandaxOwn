package tool

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"pandax/pkg/global"
	"path"
	"regexp"
	"strings"
	"time"
)

type Local struct {
	Path string
}

const (
	ImagePath    = "image" // 图片资源,包括贴图
	OtaPath      = "ota"   // 固件包
	ModelPath    = "model" // 模型资源
	DataJsonPath = "data"  //组态，规则链json
	OtherPath    = "other" //其他
)

var PathMap = map[string]string{
	ImagePath:    "uploads/image",
	OtaPath:      "uploads/ota",
	ModelPath:    "uploads/model",
	DataJsonPath: "uploads/json",
	OtherPath:    "uploads/other",
}

func init() {
	for _, path := range PathMap {
		go os.MkdirAll(path, 0755)
	}
}

func GetFilePath(fielType string) string {
	if path, ok := PathMap[fielType]; ok {
		return path
	}
	return ""
}

//@object: *Local
//@function: UploadFile
//@description: 上传文件
//@param: file *multipart.FileHeader
//@return: string, string, error

func (local *Local) UploadFile(file *multipart.FileHeader) (string, string, error) {
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	name = MD5V([]byte(name))
	// 拼接新文件名
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	// 尝试创建此路径
	mkdirErr := os.MkdirAll(local.Path, os.ModePerm)
	if mkdirErr != nil {
		global.Log.Error("function os.MkdirAll() Filed", mkdirErr.Error())
		return "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	// 拼接路径和文件名
	p := local.Path + "/" + filename

	f, openError := file.Open() // 读取文件
	if openError != nil {
		global.Log.Error("function file.Open() Filed", openError.Error())
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	out, createErr := os.Create(p)
	if createErr != nil {
		global.Log.Error("function os.Create() Filed", createErr.Error())
		return "", "", errors.New("function os.Create() Filed, err:" + createErr.Error())
	}
	defer out.Close() // 创建文件 defer 关闭

	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		global.Log.Error("function io.Copy() Filed", copyErr.Error())
		return "", "", errors.New("function io.Copy() Filed, err:" + copyErr.Error())
	}
	return p, filename, nil
}

func (local *Local) Base64ToFile(name, base64Str string) (path string, fileName string, err error) {
	fileType := "jpg"
	if strings.Contains(base64Str, "data:image") {
		re := regexp.MustCompile(`^data:image/(\w+);base64,?`)
		matchedString := re.FindStringSubmatch(base64Str)
		global.Log.Info("re", matchedString)
		if len(matchedString) <= 1 {
			return "", "", errors.New("文件Base64格式错误")
		}
		base64Str = strings.TrimPrefix(base64Str, matchedString[0])
		fileType = matchedString[1]
	}

	imgData, err := base64.StdEncoding.DecodeString(base64Str)
	filename := name + "." + fileType
	path = local.Path + "/" + filename
	_, err = os.Stat(path)
	if err == nil {
		// 文件存在,删除原有文件
		os.Remove(fileName)
	}

	file, err := os.Create(path)
	if err != nil {
		return path, "", errors.New("function os.Create() Filed, err:" + err.Error())
	}
	defer file.Close() // 创建文件 defer 关闭

	_, err = file.Write(imgData)
	if err != nil {
		return path, "", errors.New("写文件错误:" + err.Error())
	}
	return path, filename, nil
}

func (local *Local) JsonToFile(name, jsonStr string) (path string, fileName string, err error) {
	var data map[string]interface{}
	err = json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return "", "", errors.New("json 格式错误:" + err.Error())
	}
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", "", errors.New("json 格式化错误:" + err.Error())
	}

	filename := name + ".json"
	path = local.Path + "/" + filename
	_, err = os.Stat(path)
	if err == nil {
		// 文件存在,删除原有文件
		os.Remove(fileName)
	}

	file, err := os.Create(path)
	if err != nil {
		return path, "", errors.New("function os.Create() Filed, err:" + err.Error())
	}
	defer file.Close() // 创建文件 defer 关闭

	_, err = file.Write(jsonData)
	if err != nil {
		return path, "", errors.New("写文件错误:" + err.Error())
	}
	return path, filename, nil
}

func (local *Local) GetJsonFile(name string) (map[string]any, error) {
	path := local.Path + "/" + name + ".json"
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	var data = make(map[string]any)
	err = decoder.Decode(&data)
	return data, err
}

//@object: *Local
//@function: DeleteFile
//@description: 删除文件
//@param: key string
//@return: error

func (local *Local) DeleteFile(key string) error {
	p := local.Path + "/" + key
	if err := os.Remove(p); err != nil {
		return errors.New("本地文件删除失败, err:" + err.Error())
	}
	return nil
}

func GetFileMd5(filePath string) (string, error) {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	// 创建MD5哈希对象
	hash := md5.New()
	// 将文件内容拷贝到哈希对象中
	_, err = io.Copy(hash, file)
	if err != nil {
		return "", err
	}
	// 计算MD5哈希值
	hashBytes := hash.Sum(nil)
	md5String := hex.EncodeToString(hashBytes)
	return md5String, nil
}

func MD5V(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}
