package questionsubmitstatus_enum

type Status struct {
	Text  string
	Value int
}

var (
	WAITING = Status{"等待中", 0}
	RUNNING = Status{"判题中", 1}
	SUCCEED = Status{"成功", 2}
	FAILED  = Status{"失败", 3}
)
