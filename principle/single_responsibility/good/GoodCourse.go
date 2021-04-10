package good

type IGoodCourse interface {
	ID() int
	Name() string

	Controller() IPlayControl
}

type IPlayControl interface {
	Play()
	Pause()
}

type IReplayControl interface {
	IPlayControl

	Forward(seconds int)
	Backward(seconds int)
}