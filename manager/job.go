package manager

type Job struct {
	MapperCount		int		`json:"mapper_count" binding:"required"`
	ReducerCount		int		`json:"reducer_count" binding:"required"`
	
}

