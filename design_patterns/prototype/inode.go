package main

import "fmt"

type inode interface {
	print(string)
	clone() inode
}

type file struct {
	name string
}

var _ inode = (*file)(nil)

func (f *file) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *file) clone() inode {
	return &file{
		name: f.name + "_clone",
	}
}

type folder struct {
	children []inode
	name     string
}

var _ inode = (*folder)(nil)

func (f *folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, child := range f.children {
		child.print(indentation + indentation)
	}
}

func (f *folder) clone() inode {
	cloneFolder := &folder{
		name: f.name + "_clone",
	}
	var tempChildren []inode
	for _, child := range f.children {
		clone := child.clone()
		tempChildren = append(tempChildren, clone)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}
