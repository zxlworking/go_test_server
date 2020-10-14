package http_handle

import (
	"crypto/md5"
	"data"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	/*
		//获取 multi-part/form body中的form value
		for k, v := range form.Value {
			//fmt.Println("value,k,v = ", k, ",", v)
			fmt.Println("getFormData::form.Value::k = ", k, ",", len(v))

			if len(v) > 0 {
				dstFile, err := os.Create("./video_file/" + k)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				defer dstFile.Close()
				dstFile.WriteString(v[0])
			}
		}
		fmt.Println()
	*/

	r.ParseForm()
	fileName := r.URL.Query()["file_name"][0]
	fileRealName := r.URL.Query()["file_real_name"][0]
	fmt.Println("fileName", fileName)
	fmt.Println("fileRealName", fileRealName)

	/**
	底层通过调用multipartReader.ReadForm来解析
	如果文件大小超过maxMemory,则使用临时文件来存储multipart/form中文件数据
	*/
	r.ParseMultipartForm(32 << 20)
	//fmt.Println("r.Form:         ", r.Form)
	//fmt.Println("r.PostForm:     ", r.PostForm)
	//fmt.Println("r.MultiPartForm:", r.MultipartForm)

	form := r.MultipartForm

	for k, v := range form.File {
		//fmt.Println("value,k,v = ", k, ",", v)
		fmt.Println("getFormData::form.File::k = ", k, ",", len(v))

		for _, value := range v {
			fmt.Println("")
			fmt.Println("=====================getFormData::form.File::FileName = ", value.Filename)
			f, _ := value.Open()
			buf, _ := ioutil.ReadAll(f)
			fileContent := string(buf)
			fmt.Println("=====================getFormData::form.File::fileContent = ")

			md5hash := md5.New()
			io.WriteString(md5hash, fileContent)
			md5Buffer := md5hash.Sum(nil)
			fmt.Println("=================getFormData::form.File::MD5Str = %x", md5Buffer)
			md5Str := hex.EncodeToString(md5Buffer)
			fmt.Println("md5Str = " + md5Str)
			fileRealName = md5Str

			fmt.Println("=================getFormData::form.File::filepath = ", filepath.Dir(os.Args[0]))

			uploadFile, err := os.OpenFile(filepath.Dir(os.Args[0])+"/temp_file/"+fileRealName, os.O_WRONLY|os.O_CREATE, 0666)
			defer uploadFile.Close()
			if err != nil {
				fmt.Println("open new file error")
				continue
			} else {
				uploadFile.WriteString(fileContent)

				fmt.Println("getFormData::form.File::fileRealName = " + fileRealName)
			}

			if strings.Compare(fileRealName, value.Filename) == 0 {

			} /* else {
				md5hash := md5.New()
				io.WriteString(md5hash, fileContent)
				md5Buffer := md5hash.Sum(nil)
				fmt.Println("=================getFormData::form.File::MD5Str = %x", md5Buffer)
				md5Str := hex.EncodeToString(md5Buffer)
				fmt.Println("md5Str = " + md5Str)
				imgRealName = md5Str

				imgFile, err := os.OpenFile("./img_file/"+imgRealName, os.O_WRONLY|os.O_CREATE, 0666)
				defer imgFile.Close()
				if err != nil {
					fmt.Println("open new file error")
					continue
				}
				imgFile.WriteString(fileContent)

				fmt.Println("getFormData::form.File::imgRealName = " + imgRealName)
			}*/
		}
	}

	result := 0
	var response data.ResponseBaseBean
	if result == 0 {
		response = data.ResponseBaseBean{result, "success"}
	} else {
		response = data.ResponseBaseBean{result, "fail"}
	}

	responseResult, error := json.Marshal(response)
	if error != nil {

	}
	fmt.Println("main::add_video_file::success = ", string(responseResult))
	fmt.Fprint(w, string(responseResult))
}
