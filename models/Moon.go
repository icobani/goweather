/*
   Name    : Kamil KAPLAN
   Date    : 25.07.2019
*/
package models

type Moon struct {
	Rise      string `json:"rise,omitempty"`
	EpochRise int64  `json:"epochRise,omitempty"`
	Set       string `json:"set,omitempty"`
	EpochSet  int64  `json:"epochSet,omitempty"`
	Phase     string `json:"phase,omitempty"`
	Age       int32  `json:"age,omitempty"`
}
