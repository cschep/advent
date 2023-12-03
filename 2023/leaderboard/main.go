package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Member struct {
	Name       string `json:"name,omitempty"`
	ID         int    `json:"id,omitempty"`
	Stars      int    `json:"stars,omitempty"`
	LocalScore int    `json:"local_score,omitempty"`
}

func (m *Member) prettyName() string {
	if m.Name != "" {
		return m.Name
	}

	return fmt.Sprintf("%d", m.ID)
}

type LeaderboardResponse struct {
	OwnerID int
	Members map[string]Member
}

func main() {
	session, err := os.ReadFile("session.txt")
	if err != nil {
		fmt.Print(err)
	}
	sessionStr := string(session)

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://adventofcode.com/2023/leaderboard/private/view/555564.json", nil)
	if err != nil {
		panic(err)
	}

	cookie := &http.Cookie{
		Name:  "session",
		Value: sessionStr,
	}

	req.AddCookie(cookie)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var r LeaderboardResponse
	json.NewDecoder(resp.Body).Decode(&r)

	result := ""
	for _, member := range r.Members {
		line := fmt.Sprintf("-- %s %d points %d stars ", member.prettyName(), member.LocalScore, member.Stars)
		fmt.Println(line)
		result += line
	}

	err = os.WriteFile("/users/cschep/Desktop/foo.txt", []byte(result), 0644)
	if err != nil {
		panic(err)
	}

	// b, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(b))
}
