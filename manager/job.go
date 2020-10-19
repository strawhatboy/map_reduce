package manager

type Job struct {
	ID  string `json:"id"`
	MapperCount  int `json:"mapper_count" binding:"required"`
	ReducerCount int `json:"reducer_count" binding:"required"`
	MapperFuncFile	string `json:"mapper_func_file" binding:"required"`
	ReducerFuncFile string `json:"reducer_func_file" binding:"required"`
	SourceFile	[]string `json:"source_file" binding:"required"`
	FilePattern	string `json:"file_pattern" binding:"required"`
	IsDirectory *bool `json:"is_directory" binding:"required"`
	DataProvider	string `json:"data_provider" binding:"required"`
}
