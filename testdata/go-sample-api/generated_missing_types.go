package go_sample_api

type UNKNOWN_BASE_TYPE interface {
}
type OneOfObjAObjB interface {
}

func (_ *ObjA) IsOneOfObjAObjB() {
}
func (_ *ObjB) IsOneOfObjAObjB() {
}
