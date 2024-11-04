package judgeconfig

type JudgeConfig struct {
	MemoryLimit int `json:"memoryLimit"`
	StackLimit  int `json:"stackLimit"`
	TimeLimit   int `json:"timeLimit"`
}
