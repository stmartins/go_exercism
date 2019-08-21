package tournament

import "io"
import "fmt"
import "strings"

type result struct {
	MP int
	W  int
	D  int
	L  int
	P  int
}

func emptyString(str string) bool {

	fmt.Println("|" + str + "|")
	fmt.Printf("%d\n", len(str))
	if str == "" || str == "\n" || byte(str[0]) == 0 || byte(str[0]) == '#' {
		fmt.Println("string is empty")
		return true
	}
	return false
}

func Tally(r io.Reader, w io.Writer) error {

	buf := make([]byte, 300)
	for {
		_, err := r.Read(buf)
		//		fmt.Printf("n => %d buf => %s\n", n, buf[:n])
		if err == io.EOF {
			break
		}
	}
	//	fmt.Printf("=======>")
	//	fmt.Printf("%s\n", buf)
	//	fmt.Println("-----------")
	teamArray := make(map[string]result, 4)
	fmt.Println(string(buf))
	str := strings.Split(string(buf), "\n")
	fmt.Printf("len = %d\n", len(str))
	fmt.Println(str)
	var IOstr []string
	for _, s := range str {
		fmt.Printf("s =>|%s|<=\n", s)
		IOstr = strings.Split(s, ";")
		if emptyString(IOstr[0]) == false {
			fmt.Println("-----------")
			fmt.Printf("[%s]\n", IOstr[0])
			fmt.Println("-----------")
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
				fmt.Println("error")
			}
		}
	}
	display_results(teamArray, w)

	//	fmt.Println(teamArray["Allegoric Alaskians"])
	//	for k, v := range teamArray {
	//		fmt.Println(k)
	//		fmt.Println(v)
	//	}
	return nil
}

func display_results(teamArray map[string]result, w io.Writer) {
	fmt.Println("in display_results ===============")
	fmt.Println(len(teamArray))

	orderTeam := make([]result, 4)

	for index := 0; index < 4; index++ {
		for i, v := range teamArray {
			if orderTeam[0].MP == 0 {

				fmt.Println("order is nill")
			}
			fmt.Println(i)
			fmt.Println(v)
		}
	}
}
