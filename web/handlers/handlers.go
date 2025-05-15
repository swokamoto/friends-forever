package handlers

import (
    "net/http"
    "html/template"
    "github.com/go-chi/chi/v5"
    "friends-forever/internal/game" // import the game package (adjust path as needed)
	"github.com/google/uuid"
	"fmt"
)

var (
    tmpl  = template.Must(template.ParseFiles("web/templates/index.html"))
    G     = game.NewGame() // global game instance for now
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    playerName := ""
    if cookie, err := r.Cookie("player_name"); err == nil {
        playerName = cookie.Value
    }
    data := struct {
        Players    map[string]*game.Player
        GameState  string
        Lobbies    map[string]*game.Lobby
        PlayerName string
    }{
        Players:    G.Players,
        GameState:  G.GameState,
        Lobbies:    game.Lobbies,
        PlayerName: playerName,
    }
    err := tmpl.Execute(w, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    name := r.FormValue("name")
    if name == "" {
        http.Error(w, "Missing name", http.StatusBadRequest)
        return
    }
    // Add player to G.Players if not present
    if _, exists := G.Players[name]; !exists {
        G.Players[name] = &game.Player{
            ID:    name, // or generate a unique ID if you prefer
            Name:  name,
            Score: 0,
        }
    }
    http.SetCookie(w, &http.Cookie{
        Name:  "player_name",
        Value: name,
        Path:  "/",
    })
    http.Redirect(w, r, "/", http.StatusSeeOther)
}

func StartGameHandler(w http.ResponseWriter, r *http.Request) {
    player := r.URL.Query().Get("player")
    if player == "" {
        http.Error(w, "Missing player", http.StatusBadRequest)
        return
    }

	// check if player is already in a game
    if playerInAnyLobby(player) {
        http.Error(w, "You are already in a game", http.StatusBadRequest)
        return
    }
	
    id := uuid.New().String()
    lobby := &game.Lobby{
        ID:      id,
        Host:    player,
        Players: map[string]*game.Player{player: G.Players[player]},
    }
    game.Lobbies[id] = lobby

    // Render the updated lobbies list partial
    tmpl := template.Must(template.ParseFiles("web/templates/lobbies_list.html"))
    data := struct {
        Lobbies map[string]*game.Lobby
    }{
        Lobbies: game.Lobbies,
    }
    err := tmpl.Execute(w, data)
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return // <-- add this!
	}
}

func LobbiesPartialHandler(w http.ResponseWriter, r *http.Request) {
    playerName := ""
    if cookie, err := r.Cookie("player_name"); err == nil {
        playerName = cookie.Value
    }
    tmpl := template.Must(template.ParseFiles("web/templates/lobbies_list.html"))
    data := struct {
        Lobbies    map[string]*game.Lobby
        PlayerName string
    }{
        Lobbies:    game.Lobbies,
        PlayerName: playerName,
    }
    err := tmpl.Execute(w, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func JoinGameHandler(w http.ResponseWriter, r *http.Request) {
    lobbyID := r.URL.Query().Get("id")
    player := r.URL.Query().Get("player")
    if lobbyID == "" || player == "" {
        http.Error(w, "Missing lobby ID or player", http.StatusBadRequest)
        return
    }
    lobby, ok := game.Lobbies[lobbyID]
    if !ok {
        http.Error(w, "Lobby not found", http.StatusNotFound)
        return
    }

    // Ensure player exists in G.Players
    if _, ok := G.Players[player]; !ok {
        G.Players[player] = &game.Player{
            ID:    player,
            Name:  player,
            Score: 0,
        }
    }

	// Check if player is already in any lobby
	if playerInAnyLobby(player) {
        http.Error(w, "You are already in a game", http.StatusBadRequest)
        return
    }

    // Add player to lobby if not already present
    if _, exists := lobby.Players[player]; !exists {
        lobby.Players[player] = G.Players[player]
    }

    // Debug output
    fmt.Printf("G.Players: %#v\n", G.Players)
    fmt.Printf("Lobby %s players: ", lobby.ID)
    for _, v := range lobby.Players {
        if v != nil {
            fmt.Printf("%s ", v.Name)
        }
    }
    fmt.Println()

    // Render the updated lobbies list partial
    tmpl := template.Must(template.ParseFiles("web/templates/lobbies_list.html"))
    playerName := ""
    if cookie, err := r.Cookie("player_name"); err == nil {
        playerName = cookie.Value
    }
    data := struct {
        Lobbies    map[string]*game.Lobby
        PlayerName string
    }{
        Lobbies:    game.Lobbies,
        PlayerName: playerName,
    }
    err := tmpl.Execute(w, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func playerInAnyLobby(player string) bool {
    for _, lobby := range game.Lobbies {
        if _, exists := lobby.Players[player]; exists {
            return true
        }
    }
    return false
}

func RegisterRoutes(r chi.Router) {
	r.Get("/", HomeHandler)
	r.Post("/login", LoginHandler)
	r.Get("/start-game", StartGameHandler)
	r.Get("/lobbies", LobbiesPartialHandler)
	r.Get("/join-game", JoinGameHandler)

}