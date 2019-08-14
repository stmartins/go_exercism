package robotname

import "math/rand"
import "strconv"

type Robot struct {
	name string
}

var nameMap = make(map[string]int)

func randomName() string {
	var name string

	name = string(byte(rand.Intn(90-65) + 65))
	name += string(byte(rand.Intn(90-65) + 65))
	name += strconv.Itoa(rand.Intn(999-100) + 100)
	return name
}

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		for nameMap[r.name] == 1 || r.name == "" {
			r.name = randomName()
		}
	}
	nameMap[r.name]++
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}
