package system

type systemRes struct{}

func Start() *systemRes {
	DbStart()
	tgStart()
	timerStart()
	return &systemRes{}
}

func (*systemRes) Stop() {
	DbClose()
	tgClear()
	timerStop()
}
