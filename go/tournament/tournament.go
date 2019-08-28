package tournament

import "io"
import "fmt"
import "strings"
import "sort"
import "errors"

type Result struct {
	MP int
	W  int
	D  int
	L  int
	P  int
}

type Pair struct {
	name  string
	point int
}

type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool {
	if p[i].point == p[j].point && p[i].name > p[j].name {
		return true
	}
	return p[i].point < p[j].point
}
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func emptyString(str string) bool {
	if str == "" || str == "\n" || byte(str[0]) == 0 || byte(str[0]) == '#' {
		return true
	}
	return false
}

func Tally(r io.Reader, w io.Writer) error {

	buf := make([]byte, 400)
	for {
		_, err := r.Read(buf)
		if err == io.EOF {
			break
		}
	}
	teamArray := make(map[string]Result, 4)
	str := strings.Split(string(buf), "\n")
	var IOstr []string
	for _, s := range str {

		IOstr = strings.Split(s, ";")
		if emptyString(IOstr[0]) == false {
			if len(IOstr) != 3 {
				return errors.New("error input string")
			}
			switch IOstr[2] {
			case "win":
				var winner, loser = teamArray[IOstr[0]], teamArray[IOstr[1]]

				winner.MP += 1
				winner.W += 1
				winner.P += 3

				loser.MP += 1
				loser.L += 1
				loser.P += 0

				teamArray[IOstr[0]] = winner
				teamArray[IOstr[1]] = loser
			case "loss":
				var winner, loser = teamArray[IOstr[1]], teamArray[IOstr[0]]

				winner.MP += 1
				winner.W += 1
				winner.P += 3

				loser.MP += 1
				loser.L += 1
				loser.P += 0

				teamArray[IOstr[0]] = loser
				teamArray[IOstr[1]] = winner
			case "draw":
				var draw1, draw2 = teamArray[IOstr[0]], teamArray[IOstr[1]]

				draw1.MP += 1
				draw1.D += 1
				draw1.P += 1

				draw2.MP += 1
				draw2.D += 1
				draw2.P += 1

				teamArray[IOstr[0]] = draw1
				teamArray[IOstr[1]] = draw2
			default:
				return errors.New("error")
			}
		}
	}
	display_results(teamArray, w)
	return nil
}

func rankByPoints(teamsMap map[string]Result) PairList {
	pl := make(PairList, len(teamsMap))
	i := 0
	for k, v := range teamsMap {
		pl[i] = Pair{k, v.P}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

func display_results(teamArray map[string]Result, w io.Writer) {

	teamRanked := rankByPoints(teamArray)

	w.Write([]byte("Team                           | MP |  W |  D |  L |  P\n"))
	for _, v := range teamRanked {
		display_results_array(teamArray[v.name], v.name, w)
	}
}

func display_results_array(teamFinalResults Result, name string, w io.Writer) {
	w.Write([]byte(fmt.Sprintf("%s", name)))
	for i := 0; i+len(name) < 31; i++ {
		w.Write([]byte(" "))
	}
	w.Write([]byte(fmt.Sprintf("|  %d |  %d |  %d |  %d |  %d\n",
		teamFinalResults.MP, teamFinalResults.W, teamFinalResults.D,
		teamFinalResults.L, teamFinalResults.P)))
}
