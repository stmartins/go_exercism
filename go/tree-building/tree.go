package tree

import "sort"
import "errors"

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func errorChecker(records []Record) error {

	var err error

	err = nil
	if records[0].ID != 0 {
		err = errors.New("records have no root")
	} else if records[0].ID == 0 && records[0].Parent != 0 {
		err = errors.New("records have Parent")
	}

	lowerID := -1
	for _, rec := range records {
		if (rec.ID - lowerID) > 1 {
			err = errors.New("Non continuous")
			break
		} else if rec.ID != 0 && rec.ID == rec.Parent {
			err = errors.New("direct find")
			break
		} else if rec.ID < rec.Parent {
			err = errors.New("cycle indirect")
			break
		}
		if lowerID != rec.ID {
			lowerID = rec.ID
		} else {
			err = errors.New("Duplicate find")
			break
		}
	}
	return err
}

func Build(records []Record) (*Node, error) {

	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	err := errorChecker(records)
	if err != nil {
		return nil, err
	}

	node := make([]Node, len(records))
	for i, v := range records {
		node[i].ID = v.ID
		if v.ID != v.Parent {
			node[v.Parent].Children = append(node[v.Parent].Children, &node[v.ID])
		}
	}
	return &node[0], nil
}
