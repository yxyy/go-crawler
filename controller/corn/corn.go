package corn

type Corn interface {
	Run()
}

func Run(corn Corn)  {
	corn.Run()
}
