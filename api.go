package tbapi

import (
	"io"
	"log"
	"regexp"
	"strconv"
	"time"
)

type TabroomApi struct {
	username string
	password string
	client   httpClient
}

func (t *TabroomApi) GetTournaments() ([]Tournament, error) {
	if err := t.ensureAuthenticated(); err != nil {
		return nil, err
	}
	resp, err := t.client.get("/user/tourn/all.mhtml")
	if err != nil {
		return nil, err
	}
	content, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	return parseTournaments(string(content))
}

func parseTournaments(html string) ([]Tournament, error) {
	otherRe := regexp.MustCompile(`(?is)<td[^>]*>\s*([0-9]{4}-[0-9]{2}-[0-9]{2})\s*</td>.*?href="select\.mhtml\?tourn_id=([0-9]+)".*?>\s*([^<\s](?:[^<]*?[^<\s])?)\s*</a>`)

	matches := otherRe.FindAllStringSubmatch(html, -1)
	tournaments := make([]Tournament, 0, len(matches))
	for _, match := range matches {
		tournamentDate, err := time.Parse(time.DateOnly, match[1])
		if err != nil {
			log.Printf("Unable to convert %s to time.Time", matches[1])
			continue
		}
		tournamentId, err := strconv.Atoi(match[2])
		if err != nil {
			log.Printf("Unable to convert %s to int", match[2])
			continue
		}
		tournamentName := match[3]
		tournaments = append(tournaments, Tournament{
			Id:   tournamentId,
			Date: tournamentDate,
			Name: tournamentName,
		})
	}

	return tournaments, nil
}
