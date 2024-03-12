package main

import (
	"fmt"
	"sort"
)

type Object struct {
	ID int
}

func uniqObjects(objects []Object) []Object {
	var result []Object
	objectIdSet := make(map[int]struct{})

	for _, o := range objects {
		if _, ok := objectIdSet[o.ID]; !ok {
			result = append(result, o)
			objectIdSet[o.ID] = struct{}{}
		}
	}

	return result
}

func sortObjects(objects []Object) {
	sort.Slice(objects, func(i, j int) bool {
		return objects[i].ID < objects[j].ID
	})
}

func printObjects(objects []Object) {
	for _, o := range objects {
		fmt.Printf("<Object ID=%v>\n", o.ID)
	}
}

func sortAndUniqObjects(objects []Object) []Object {
	result := uniqObjects(objects)
	sortObjects(result)
	return result
}

func main() {
	unsortedDuplicatedObjects := []Object{Object{3}, Object{2}, Object{1}, Object{2}}

	fmt.Println("Income objects:")
	printObjects(unsortedDuplicatedObjects)

	fmt.Println("Uniq and sorted objects:")
	sortedUniqObjects := sortAndUniqObjects(unsortedDuplicatedObjects)
	printObjects(sortedUniqObjects)
}
