package model

type HttpTpyeEnum uint8

const (
	HttpTpyeAdd    HttpTpyeEnum = iota //添加
	HttpTpyeDelete                     //删除
	HttpTpyeChange                     //修改
)
