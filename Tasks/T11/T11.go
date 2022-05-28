package main

import "fmt"

func main() {
	set1 := map[string]struct{}{"a": {}, "b": {}, "c": {}, "d": {}, "e": {}}
	set2 := map[string]struct{}{"d": {}, "e": {}, "b": {}, "i": {}, "l": {}}

	inter := make(map[string]struct{})

	if len(set2) > len(set1) {
		set1, set2 = set2, set1
	}

	for S1key := range set1 {
		for S2key := range set2 {
			if S1key == S2key {
				inter[S1key] = struct{}{}
			}
		}
	}
	fmt.Println(inter)
}
