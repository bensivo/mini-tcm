package folders

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

type Folder struct {
	Name     string
	PathRel  string // Path relative to the root folder used when this folder object was created
	PathAbs  string // Absolute path on the underlying filesystem
	Children []*Folder
}

// LoadFromFs instantiates a folder struct based on a filesystem directory path, recursively
// walking the filesystem to populate the folder and all of its children.
func LoadFromFs(basePath string) (*Folder, error) {
	basePathAbs, err := filepath.Abs(basePath)
	if err != nil {
		return nil, err
	}

	foldersMap := map[string]*Folder{}

	err = filepath.WalkDir(basePathAbs, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		pathAbs := path // Becuase we called filepath.Abs earlier on basePath, this path will always be abs too

		if d.IsDir() {
			if pathAbs == basePathAbs { // Handle root folder
				folder := Folder{
					Name:     d.Name(),
					PathRel:  "/",
					PathAbs:  pathAbs,
					Children: []*Folder{},
				}
				foldersMap[path] = &folder
				return nil
			}

			// Handle any non-foot folder
			parentPath := filepath.Dir(pathAbs)
			parentFolder, ok := foldersMap[parentPath]
			if !ok {
				// NOTE: in theory this should never happen, because filepath.WalkDir uses a depth-first search,
				// we should always hit the parent folder before any of its children.
				return fmt.Errorf("parent folder not found for %s", pathAbs)
			}

			pathRel := strings.Replace(pathAbs, basePathAbs, "", 1)
			folder := Folder{
				Name:     d.Name(),
				PathRel:  pathRel,
				PathAbs:  pathAbs,
				Children: []*Folder{},
			}
			parentFolder.Children = append(parentFolder.Children, &folder) // Add this folder to its parent's "children"
			foldersMap[pathAbs] = &folder                                  // Add this folder to the foldersMap, in case it's a parent too
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	rootFolder := foldersMap[basePathAbs]
	return rootFolder, nil
}

// PrettyPrint prints a folder object and all of its children in a human-readable format.
func PrettyPrint(folder *Folder) {
	prettyPrintRecursive(folder, "")
}

func prettyPrintRecursive(folder *Folder, indent string) {
	fmt.Printf("%s%s\n", indent, folder.Name)
	for _, child := range folder.Children {
		prettyPrintRecursive(child, indent+"  ")
	}
}
