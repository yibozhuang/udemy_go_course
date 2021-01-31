package main

import "fmt"

func main() {
	files := make([]*file, 0)
	for i := 0; i < 3; i++ {
		name := fmt.Sprintf("File%d", i+1)
		f := &file{name: name}
		files = append(files, f)
	}

	folder1 := &folder{
		children: []inode{files[0]},
		name:     "Folder1",
	}

	folder2 := &folder{
		children: []inode{folder1, files[1], files[2]},
		name:     "Folder2",
	}

	fmt.Println("Folder2 details:")
	folder2.print("  ")

	cloneFolder := folder2.clone()
	fmt.Println("Cloned folder details:")
	cloneFolder.print("  ")
}
