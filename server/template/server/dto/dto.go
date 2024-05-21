package TableNameDto

type TableNameCreateRequestDTO struct {
	Name string `json:"name" binding:"required"`
	Age  int8   `json:"age" binding:"required"`
}

type TableNameUpdateRequestDTO struct {
	Name string `json:"name" binding:"required"`
	Age  int8   `json:"age" binding:"required"`
}

type TableNameFindRequestDTO struct {
	Name   string `json:"name"`
	Age    int8   `json:"age"`
	Limit  int64  `json:"limit"`
	Offset int64  `json:"offset"`
}
