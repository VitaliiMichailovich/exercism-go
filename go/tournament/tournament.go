package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type team struct {
	name                           string
	plays, won, draw, lost, points int
}

type scoreBoard map[string]*team

type scoreArr []team

func Tally(reader io.Reader, writer io.Writer) error {
	scanner := bufio.NewScanner(reader)
	board := make(scoreBoard)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if err := board.game(line); err != nil {
			return err
		}
	}
	teamBoard := board.makeArrOfScoreBoard()
	sort.Sort(teamBoard)
	header := "Team                           | MP |  W |  D |  L |  P\n"
	io.WriteString(writer, header)
	for _, board := range teamBoard {
		io.WriteString(writer, board.String()+"\n")
	}
	return nil
}

func (board scoreBoard) game(game string) error {
	match := strings.Split(game, ";")
	if len(match) != 3 {
		return errors.New("Something wrong with game: " + game)
	}
	first, fFirst := board[match[0]]
	if !fFirst {
		first = &team{name: match[0]}
		board[match[0]] = first
	}
	second, fSecond := board[match[1]]
	if !fSecond {
		second = &team{name: match[1]}
		board[match[1]] = second
	}
	switch match[2] {
	case "win":
		first.winTeam()
		second.loseTeam()
	case "loss":
		first.loseTeam()
		second.winTeam()
	case "draw":
		first.drawTeam()
		second.drawTeam()
	default:
		return fmt.Errorf("Unknown win condition: %s", game)
	}
	return nil
}

func (t *team) winTeam() {
	t.plays++
	t.won++
	t.points += 3
}

func (t *team) loseTeam() {
	t.plays++
	t.lost++
}

func (t *team) drawTeam() {
	t.plays++
	t.draw++
	t.points++
}

func (s scoreArr) Len() int {
	return len(s)
}

func (s scoreArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s scoreArr) Less(i, j int) bool {
	if s[i].points != s[j].points {
		return s[i].points > s[j].points
	}
	if s[i].won != s[i].won {
		return s[i].won > s[j].won
	}
	return s[i].name < s[j].name
}

func (sb scoreBoard) makeArrOfScoreBoard() scoreArr {
	var arrTeam scoreArr
	for _, team := range sb {
		arrTeam = append(arrTeam, *team)
	}
	return arrTeam
}

func (t team) String() string {
	return fmt.Sprintf("%-30v | %2d | %2d | %2d | %2d | %2d", t.name, t.plays, t.won, t.draw, t.lost, t.points)
}
