package hello

type EchoRspDto struct {
	StrVal    string `json:"strVal"`
	IntVal    int    `json:"intVal"`
	IntPtrVal int    `json:"intPtrVal"`
	StructVal struct {
		Id int `json:"id"`
	} `json:"structVal"`
	SliceVal []int `json:"sliceVal"`
}
