package players

type Player interface {
	Run() Player
	Close()
}
