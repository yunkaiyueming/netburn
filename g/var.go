package g

type User struct {
	Email string
	Name  string
	Age   float64
	Memo  string
}

type CronJob struct {
	Name, IsRun, NextStartTime string
}

type ActionItems struct {
	Id   string
	Url  string
	Desc string
}

type HFile struct {
	Key        string
	Name       string
	Size       int64
	CreateTime string
	Tag        string
}
