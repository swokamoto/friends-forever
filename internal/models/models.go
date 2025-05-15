package models

type Player struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GameState struct {
	ID      int      `json:"id"`
	Players []Player `json:"players"`
	Status  string   `json:"status"`
}

type Move struct {
	PlayerID int    `json:"player_id"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
}