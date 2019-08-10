package tree

import "fmt"
import "sort"

//import "errors"

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {

	//	rootID := 0

	if len(records) == 0 {
		return nil, nil
	}
	fmt.Println("======start============")
	fmt.Printf("records => ")
	fmt.Println(records)
	sort.Slice(records, func(i, j int) bool {
		fmt.Printf("i => %d et j =>%d\n", records[i].ID, records[j].ID)
		return records[i].ID < records[j].ID
	})

	fmt.Printf("len %d\n", len(records))
	node := make([]Node, len(records))
	for i, v := range records {
		fmt.Printf("value [%d]\n", v)
		node[i].ID = v.ID
		if v.ID != v.Parent {
			node[v.Parent].Children = append(node[v.Parent].Children, &node[v.ID])
		}
	}
	fmt.Printf("node value => [%v]\n\n", node)
	return &node[0], nil
}
