package nox

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"sort"
	"strings"

	noxflags "nox/v1/common/flags"
	"nox/v1/common/log"
	"nox/v1/internal/version"
)

var (
	apiLog    = log.New("api")
	apiTokens = make(map[string]struct{})
)

func init() {
	if tok := os.Getenv("NOX_API_TOKEN"); tok != "" {
		apiTokens[tok] = struct{}{}
	}
	const pref = "/api/v0/game"
	gameMux.HandleFunc(pref+"/info", handleServerInfo)
	gameMux.HandleFunc(pref+"/map", handleChangeMap)
	gameMux.HandleFunc(pref+"/cmd", handleRunCmd)
	gameMux.HandleFunc(pref+"/lua", handleRunLUA)
	var token string
	configStrPtr("server.api_token", "NOX_API_TOKEN", "", &token)
	registerOnConfigRead(func() {
		if token != "" {
			apiTokens[token] = struct{}{}
		}
	})
}

type gameInfoPlayers struct {
	Cur  int              `json:"cur"`
	Max  int              `json:"max"`
	List []gameInfoPlayer `json:"list"`
}

type gameInfoPlayer struct {
	Name  string `json:"name"`
	Class string `json:"class,omitempty"`
}

type gameInfoQuest struct {
	Stage int `json:"stage"`
}

type gameInfoResp struct {
	Name    string          `json:"name"`
	Map     string          `json:"map"`
	Mode    string          `json:"mode"`
	Vers    string          `json:"vers"`
	Players gameInfoPlayers `json:"players"`
	Quest   *gameInfoQuest  `json:"quest,omitempty"`
}

func getGameInfo(ctx context.Context) (*gameInfoResp, error) {
	ch := make(chan *gameInfoResp, 1)
	addGameTickHook(func() {
		v := &gameInfoResp{
			Name: getServerName(),
			Map:  strings.ToLower(getServerMap()),
			Vers: version.Version(),
			Mode: noxflags.GetGame().ModeString(),
		}
		if noxflags.HasGame(noxflags.GameModeQuest) {
			v.Quest = &gameInfoQuest{Stage: nox_game_getQuestStage_4E3CC0()}
		}
		v.Players.Max = getServerMaxPlayers()
		for _, p := range getPlayers() {
			v.Players.List = append(v.Players.List, gameInfoPlayer{
				Name:  p.Name(),
				Class: p.PlayerClass().String(),
			})
		}
		v.Players.Cur = len(v.Players.List)
		ch <- v
	})
	select {
	case v := <-ch:
		sort.Slice(v.Players.List, func(i, j int) bool {
			return v.Players.List[i].Name < v.Players.List[j].Name
		})
		return v, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func queueServerMapLoad(name string) {
	name = strings.TrimSpace(name)
	apiLog.Printf("load map: %q", name)
	addGameTickHook(func() {
		serverCmdLoadMap(name)
	})
}

func queueServerCmd(cmd string) {
	cmd = strings.TrimSpace(cmd)
	apiLog.Printf("run command: %q", cmd)
	addGameTickHook(func() {
		execServerCmd(cmd)
	})
}

func queueMapLUA(code string) {
	code = strings.TrimSpace(code)
	apiLog.Printf("run lua: %q", code)
	addGameTickHook(func() {
		runMapLUA(code)
	})
}

func handleServerInfo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	case http.MethodGet, http.MethodHead:
	}
	resp, err := getGameInfo(r.Context())
	if err != nil {
		// TODO: set correct status code
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

func handlePostStr(w http.ResponseWriter, r *http.Request, limit int, fnc func(str string)) {
	switch r.Method {
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	case http.MethodHead:
		return
	case http.MethodPost:
	}
	defer r.Body.Close()
	token := r.Header.Get("X-Token")
	if _, ok := apiTokens[token]; !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	data, err := io.ReadAll(io.LimitReader(r.Body, int64(limit)+1))
	if err != nil || len(data) > limit {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fnc(string(data))
}

func handleChangeMap(w http.ResponseWriter, r *http.Request) {
	handlePostStr(w, r, 256, queueServerMapLoad)
}

func handleRunCmd(w http.ResponseWriter, r *http.Request) {
	handlePostStr(w, r, 256, queueServerCmd)
}

func handleRunLUA(w http.ResponseWriter, r *http.Request) {
	handlePostStr(w, r, 4096, queueMapLUA)
}
