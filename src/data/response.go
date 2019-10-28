package data

type ResponseBaseBean struct {
	Code int
	Desc string
}

type EvaluateSelfResponseBean struct {
	BaseBean             ResponseBaseBean
	EvaluateSelfInfoList []EvaluateSelfInfo
}

type MovieInfoListResponseBean struct {
	BaseBean      ResponseBaseBean
	MovieInfoList []MovieInfo
}

type MovieDetailInfoListResponseBean struct {
	BaseBean            ResponseBaseBean
	MovieDetailInfoList []MovieDetailInfo
}
