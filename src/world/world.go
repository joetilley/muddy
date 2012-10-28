package world

type Room struct {
	Desc       string
	N, E, S, W *Room
}

type GameWorld struct {
	//Areas map[string]Area <-- Will need this later
	Rooms     map[string]*Room
	StartRoom *Room
}
