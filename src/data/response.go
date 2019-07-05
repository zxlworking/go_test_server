package data

type ResponseBaseBean struct {
	Code int
	Desc string
}

type EvaluateSelfResponseBean struct {
	BaseBean         ResponseBaseBean
	EvaluateSelfInfo EvaluateSelfInfo
}
