package worldloader

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"world"
)

type JsonRoom struct {
	Name, Desc             string
	N_to, S_to, E_to, W_to string
}

type JsonGameData struct {
	StartRoom string
	Rooms     []JsonRoom
}

func LoadWorld() *world.GameWorld {
	// load the rooms into temporary structures
	b, err := ioutil.ReadFile("world.json")
	if err != nil {
		log.Fatal(err)
	}

	var data JsonGameData
	err = json.Unmarshal(b, &data)
	if err != nil {
		log.Fatal(err)
	}

	// build the final map
	roomMap := buildRoomMap(data)
	return &world.GameWorld{Rooms: roomMap, StartRoom: roomMap[data.StartRoom]}
}

func buildRoomMap(data JsonGameData) map[string]*world.Room {
	r := make(map[string]*world.Room)
	return r
}
