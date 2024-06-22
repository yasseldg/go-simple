package easyway

import "fmt"

type InterEwValues interface {
	String() string

	Head() float64
	Tail() float64
}

type EwValues struct {
	M_head float64 `bson:"h,omitempty" json:"h,omitempty"`
	M_tail float64 `bson:"t,omitempty" json:"t,omitempty"`
}

func (ew *EwValues) String() string {
	return fmt.Sprintf("%f - %f", ew.M_head, ew.M_tail)
}

func (ew *EwValues) Head() float64 {
	return ew.M_head
}

func (ew *EwValues) Tail() float64 {
	return ew.M_tail
}
