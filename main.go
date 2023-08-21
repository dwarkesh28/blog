package main

import "blog/cmd"

func main() {
	cmd.Execute()
}

// func createFolder(path string) {
// 	err := syscall.Mkdir(path, 0755)
// 	if err != nil {
// 		fmt.Println("err: ", err)
// 	}
// }

// package main

// import (
// 	"log"
// 	"os"
// )

// func createFolderIfNotExist(folderPath string) error {
// 	// Get the current working directory
// 	currentDir, err := os.Getwd()
// 	fmt.Println(currentDir)
// 	if err != nil {
// 		return fmt.Errorf("failed to get the current working directory: %s", err)
// 	}

// 	// Construct the full path for the "dex" folder
// 	fullPath := filepath.Join(currentDir, folderPath)

// 	// Check if the folder already exists
// 	_, err = os.Stat(fullPath)

// 	// If the folder doesn't exist, create it
// 	if os.IsNotExist(err) {
// 		err = os.Mkdir("."+string(filepath.Separator)+folderPath, os.ModePerm)
// 		if err != nil {
// 			return fmt.Errorf("failed to create folder: %s", err)
// 		}
// 		fmt.Printf("Folder %s created successfully\n", fullPath)
// 	} else if err != nil {
// 		return fmt.Errorf("failed to check folder existence: %s", err)
// 	} else {
// 		fmt.Printf("Folder %s already exists\n", fullPath)
// 	}

// 	return nil
// }

// package main

// import (
// 	"fmt"
// 	"os"
// 	"path/filepath"
// )

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	// Specify the folder name to create
// 	folderName := "mydir"

// 	// Specify the absolute path where the folder will be created
// 	// Change this path to your desired location
// 	targetDir := "D:\\projects\\Golang\\Blog\\blog"

// 	// Combine the target directory and folder name to create the full path
// 	folderPath := targetDir + "\\" + folderName

// 	err := os.MkdirAll(folderPath, 0755)
// 	if err != nil {
// 		fmt.Printf("Failed to create directory: %s\n", err)
// 		return
// 	}

// 	fmt.Printf("Directory %s created successfully\n", folderPath)
// }

// package main

// import (
// 	"log"
// 	"os"
// )

// func main() {
// 	err := os.Mkdir("temp1", 0755)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
