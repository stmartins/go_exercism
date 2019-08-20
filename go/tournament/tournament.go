package tournament

import "io"
import "fmt"
import "strings"

func emptyString(str string) bool {

	fmt.Println("|" + str + "|")
	fmt.Printf("%d\n", len(str))
	if str == "" || str == "\n" || byte(str[0]) == 0 {
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
	teamArray := make(map[string]int, 4)
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
				teamArray[IOstr[0]] += 3
				teamArray[IOstr[1]] = teamArray[IOstr[1]]
			case "loss":
				teamArray[IOstr[0]] = teamArray[IOstr[0]]
				teamArray[IOstr[1]] = teamArray[IOstr[1]]
			case "draw":
				teamArray[IOstr[0]]++
				teamArray[IOstr[1]]++
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

func display_results(teamArray map[string]int, w io.Writer) {

}
