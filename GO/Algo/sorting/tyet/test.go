package main

import "fmt"

type book struct {
	name       string
	checkedout bool
}

type members struct {
	name string
}

func printmemberlist(memberList [4]members) {
	for i := 0; i < len(memberList); i++ {
		members := memberList[i]
		fmt.Println(members.name)
	}
}

func booklist(libraryBooks [4]book) (checkedout []book) {
	for i := 0; i < len(libraryBooks); i++ {
		if libraryBooks[i].checkedout {
			checkedout = append(checkedout, libraryBooks[i])
		}
	}
	return checkedout
}

//func checkedoutoflibrary(memberList[4]members, checkedout string) string {
//  fmt.Println("Checked out of Library", members.name & checkedout)
//}

func main() {
	memberList := [...]members{
		{name: "Julie"},
		{name: "Sally"},
		{name: "Fred"},
		{name: "Kyle"}}
	printmemberlist(memberList)
	libraryBooks := [...]book{
		{name: "Harry Potter"},
		{name: "Eragon"},
		{name: "Sun Stand Still"},
		{name: "Terraform In Action"},
	}
	booklist(libraryBooks)
	//fmt.Println("Member", memberList[1].name, "Has Checked out", libraryBooks[3].checkedout = true)
	//checkedoutoflibrary(memberList[1].name)
	libraryBooks[2].checkedout = true
	libraryBooks[3].checkedout = true
	fmt.Println("books which have been checked out:", booklist(libraryBooks))
}
