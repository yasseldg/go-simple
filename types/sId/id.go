package sId

type ObjectId interface {
	Get(interface{}) interface{}
	IsValid() bool
}
