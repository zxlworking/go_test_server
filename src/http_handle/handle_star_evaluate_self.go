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

	base_db "db"

	"github.com/huaweicloud/huaweicloud-sdk-go-frs/client/param"
	"github.com/huaweicloud/huaweicloud-sdk-go-frs/client/result/common"
	"github.com/huaweicloud/huaweicloud-sdk-go-frs/client/service"
)

func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

func compareFace(imgPath1, imgPath2 string) {
	fmt.Println("compareFace::imgPath1::", imgPath1)
	fmt.Println("compareFace::imgPath2::", imgPath2)
	//Init frs client
	endpoint := "https://face.cn-north-1.myhuaweicloud.com"
	ak := "FIC1EXY74NXMOELTXLZC"
	sk := "1xEplN2aedTj7uhPRbZTxgF09Wt1jU0H3c17Ci7i"
	projectId := "05819bf0a6800f662ff8c0169b5c9fbc"
	authInfo := &param.AuthInfo{EndPoint: endpoint, Ak: ak, Sk: sk}
	frs := service.NewFrsClient(authInfo, projectId)

	//Get all service
	frs.GetCompareService()
	frs.GetDetectService()
	frs.GetFaceService()
	frs.GetFaceSetService()
	frs.GetLiveDetectService()
	frs.GetQualityService()
	frs.GetSearchService()

	//Use frs compare service
	res, err := frs.GetCompareService().CompareFaceByFile(imgPath1, imgPath2)
	if nil != err {
		fmt.Println("CompareFaceByFile err::", err)
	} else {
		fmt.Println(res.Similarity)
		fmt.Println(res.Image1Face.BoundingBox)
		fmt.Println(res.Image2Face.BoundingBox)
	}
}

func searchFace(imgPath string) []common.ComplexFace {
	fmt.Println("compareFace::imgPath::", imgPath)
	//Init frs client
	endpoint := "https://face.cn-north-1.myhuaweicloud.com"
	ak := "FIC1EXY74NXMOELTXLZC"
	sk := "1xEplN2aedTj7uhPRbZTxgF09Wt1jU0H3c17Ci7i"
	projectId := "05819bf0a6800f662ff8c0169b5c9fbc"
	authInfo := &param.AuthInfo{EndPoint: endpoint, Ak: ak, Sk: sk}
	frs := service.NewFrsClient(authInfo, projectId)

	res, err := frs.GetSearchService().SearchFaceByFile("zxl_test_1", imgPath)

	searchFaceId := ""
	fmt.Println("SearchFaceByFile start searchFaceId::", (searchFaceId == ""))
	if nil != err {
		fmt.Println("SearchFaceByFile err::", err)
	} else {
		return res.Faces
	}
	return nil
}

func StarEvaluateSelf(w http.ResponseWriter, r *http.Request) {
	fmt.Println("StarEvaluateSelf::start")
	r.ParseForm()
	testParam := r.URL.Query()["test_param"][0]
	fmt.Println("testParam", testParam)

	/**
	底层通过调用multipartReader.ReadForm来解析
	如果文件大小超过maxMemory,则使用临时文件来存储multipart/form中文件数据
	*/
	r.ParseMultipartForm(32 << 20)
	//fmt.Println("r.Form:         ", r.Form)
	//fmt.Println("r.PostForm:     ", r.PostForm)
	//fmt.Println("r.MultiPartForm:", r.MultipartForm)

	form := r.MultipartForm

	var evaluateSelfResponseBean data.EvaluateSelfResponseBean

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
			imgRealName := md5Str + value.Filename
			fmt.Println("getFormData::form.File::imgRealName = " + imgRealName)

			imgDir := "./img_file/"
			if !IsExist(imgDir) {
				fmt.Println(imgDir, "not exist")
				err := os.Mkdir(imgDir, os.ModePerm)
				if err != nil {
					fmt.Println(err)
				}
			}
			imgPath := "./img_file/" + imgRealName
			if !IsExist(imgPath) {
				fmt.Println(imgPath, "not exist")
				file, err := os.Create(imgPath)
				if err != nil {
					fmt.Println(err)
				}
				defer file.Close()
			}
			imgFile, err := os.OpenFile(imgPath, os.O_WRONLY|os.O_CREATE, 0666)
			defer imgFile.Close()
			if err != nil {
				fmt.Println("open new file error", err.Error())
			} else {
				imgFile.WriteString(fileContent)

				absPath, err := filepath.Abs(imgPath)
				if err != nil {
					fmt.Println("filepath.Abs error", err.Error())
				} else {
					searchResult := searchFace(absPath)

					os.Remove(absPath)

					fmt.Println(searchResult)
					for _, item := range searchResult {
						searchFaceId := item.Face.FaceId

						var sql_str = "SELECT " +
							"id, star_name,star_img_url,star_detail_url,face_id " +
							"FROM star_info WHERE face_id = '" + searchFaceId + "'"

						rows, stmt, db := base_db.Query(sql_str)

						var starInfoItem data.StarInfo
						var evaluateSelfInfoItem data.EvaluateSelfInfo
						for rows.Next() {
							var id, star_name, star_img_url, star_detail_url, face_id string
							rows.Scan(&id, &star_name, &star_img_url, &star_detail_url, &face_id)
							fmt.Println("query::result = ", id, star_name, star_img_url, star_detail_url, face_id)

							starInfoItem = data.StarInfo{id, star_name, star_img_url, star_detail_url, face_id}
							evaluateSelfInfoItem.StarInfo = starInfoItem
						}
						evaluateSelfInfoItem.Similarity = item.Similarity
						evaluateSelfResponseBean.EvaluateSelfInfoList = append(evaluateSelfResponseBean.EvaluateSelfInfoList, evaluateSelfInfoItem)

						defer rows.Close()
						defer stmt.Close()
						defer db.Close()
					}
				}

			}
		}
	}

	if len(evaluateSelfResponseBean.EvaluateSelfInfoList) > 0 {
		evaluateSelfResponseBean.BaseBean.Code = 0
		evaluateSelfResponseBean.BaseBean.Desc = "success"
	} else {
		evaluateSelfResponseBean.BaseBean.Code = 1
		evaluateSelfResponseBean.BaseBean.Desc = "fail"
		var starInfoItem data.StarInfo
		var evaluateSelfInfoItem data.EvaluateSelfInfo
		evaluateSelfInfoItem.StarInfo = starInfoItem
		evaluateSelfResponseBean.EvaluateSelfInfoList = append(evaluateSelfResponseBean.EvaluateSelfInfoList, evaluateSelfInfoItem)
	}

	responseResult, responseError := json.Marshal(evaluateSelfResponseBean)
	fmt.Println("query::responseError = ", responseError)
	fmt.Println("query::responseResult = ", string(responseResult))
	fmt.Fprint(w, string(responseResult))
}
