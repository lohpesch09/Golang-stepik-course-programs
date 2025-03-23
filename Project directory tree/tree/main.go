package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	depth := 0
	branches := make([]string, 0)
	branches = append(branches, "├───")
	if printFiles == false {
		dirTreeWithoutFiles(out, path, depth, branches)
	} else {
		dirTreeWithFiles(out, path, depth, branches)
	}

	return nil
}

func dirTreeWithoutFiles(out io.Writer, path string, depth int, branches []string) error {
	dirs, _ := os.ReadDir(path)
	if len(dirs) == 0 {
		fmt.Errorf("Empty directory")
	}

	dirsWithoutFiles := make([]os.DirEntry, 0)

	for _, dir := range dirs {
		if dir.IsDir() {
			dirsWithoutFiles = append(dirsWithoutFiles, dir)
		}
	}

	os.Chdir(path)

	for index, dirInternal := range dirsWithoutFiles {

		for i := 0; i < depth; i++ {
			fmt.Fprint(out, branches[i])
		}

		internalDirs, _ := os.ReadDir(dirInternal.Name())

		internalDirsWithoutFiles := make([]os.DirEntry, 0)

		for _, internalDir := range internalDirs {
			if internalDir.IsDir() {
				internalDirsWithoutFiles = append(internalDirsWithoutFiles, internalDir)
			}
		}

		if len(internalDirsWithoutFiles) != 0 {

			if index == len(dirsWithoutFiles)-1 {
				branches[depth] = "└───"
				fmt.Fprintf(out, "%s%s\n", branches[depth], dirInternal.Name())
				branches[depth] = "	"
			} else {
				fmt.Fprintf(out, "%s%s\n", branches[depth], dirInternal.Name())
				branches[depth] = "│	"
			}
			depth++
			branches = append(branches, "├───")
			dirTreeWithoutFiles(out, dirInternal.Name(), depth, branches)
			depth--
			os.Chdir("..")
		} else if len(internalDirsWithoutFiles) == 0 {
			if index == len(dirsWithoutFiles)-1 {
				branches[depth] = "└───"
				fmt.Fprintf(out, "%s%s\n", branches[depth], dirInternal.Name())
				depth--
				branches[depth] = "├───"
			} else {
				fmt.Fprintf(out, "%s%s\n", branches[depth], dirInternal.Name())
			}
			continue
		}
	}
	return nil
}

func dirTreeWithFiles(out io.Writer, path string, depth int, branches []string) error {
	dirsWithFiles, _ := os.ReadDir(path)
	if len(dirsWithFiles) == 0 {
		fmt.Errorf("Empty directory")
	}

	os.Chdir(path)

	for index, dir := range dirsWithFiles {
		for i := 0; i < depth; i++ {
			fmt.Fprint(out, branches[i])
		}

		internalDirs, _ := os.ReadDir(dir.Name())

		if len(internalDirs) != 0 {

			if index == len(dirsWithFiles)-1 {
				branches[depth] = "└───"
				fmt.Fprintf(out, "%s%s\n", branches[depth], dir.Name())
				branches[depth] = "	"
			} else {
				fmt.Fprintf(out, "%s%s\n", branches[depth], dir.Name())
				branches[depth] = "│	"
			}
			depth++
			branches = append(branches, "├───")
			dirTreeWithFiles(out, dir.Name(), depth, branches)
			depth--
			branches[depth] = "├───"
			branches = branches[:depth+1]
			os.Chdir("..")
		} else if len(internalDirs) == 0 {
			if index == len(dirsWithFiles)-1 {
				branches[depth] = "└───"
				fmt.Fprint(out, branches[depth])
				depth--
				if depth >= 0 {
					branches[depth] = "├───"
				}
				branches = branches[:depth+1]
			} else {
				fmt.Fprint(out, branches[depth])
			}
			if !dir.IsDir() {
				stat, _ := os.Stat(dir.Name())
				if stat.Size() != 0 {
					fmt.Fprintf(out, "%s", dir.Name())
					fmt.Fprintf(out, " (%db)\n", stat.Size())
				} else {
					fmt.Fprintf(out, "%s", dir.Name())
					fmt.Fprintln(out, " (empty)")
				}
			} else {
				fmt.Fprintf(out, "%s\n", dir.Name())
			}
			continue
		}
	}

	return nil
}
