package main

import (
	"ethos/ethos"
	"ethos/efmt"
	"ethos/syscall"
)

func main() {
	dirName := "TestDir"
	CleanUp(dirName)
	CleanUp(dirName+"_Copy")
	SeedData(dirName)
	CopyDir("/user/" + dirName, "/user/" + dirName + "_Copy")
}

func CleanUp(dirName string) {
	ID := "CopyDir:"

	ethos.RemoveFilePath("/user/" + dirName + "/IntDir/SF1")
	ethos.RemoveFilePath("/user/" + dirName + "/IntDir/SF2")
	fd, s := ethos.OpenDirectoryPath("/user/" + dirName)
	s = ethos.RemoveDirectory(fd, "IntDir")
	efmt.Println(ID,"Remove Int Directory: ",s)

	ethos.RemoveFilePath("/user/" + dirName + "/F2")
	ethos.RemoveFilePath("/user/" + dirName + "/F1")
	fd, s = ethos.OpenDirectoryPath("/user")
	s = ethos.RemoveDirectory(fd, dirName)
	efmt.Println(ID,"Remove Test Directory: ",s)
	syscall.Close(fd)
}

func SeedData(name string) {
	ID := "CopyDir:"
	efmt.Println(ID, "Creating directory named", name)
	path := "/user/" + name

	var d1 String
	status := d1.CreateDirectory(path, "")
	if status != syscall.StatusOk {
		efmt.Print("%v Unable to create source directory. Status: %v\n", ID, status)
		return
	}

	efmt.Println(ID, "Writing first file")
	d1 = "Hello"

	d1.WriteVar(path + "/F1")

	efmt.Println(ID, "Writing second file")
	d1 = "Hi!"
	d1.WriteVar(path + "/F2")

	var d2 Uint32
	efmt.Println(ID, "Creating subdirectory named IntDir")
	status = d2.CreateDirectory(path+"/IntDir", "")
	if status != syscall.StatusOk {
		efmt.Print("%v Unable to create IntDir. Status: %v\n", ID, status)
	}

	path = path + "/IntDir"
	efmt.Println(ID, "Writing first int file")
	d2 = 456
	d2.WriteVar(path + "/SF1")

	efmt.Println(ID, "Writing second int file")
	d2 = 555
	d2.WriteVar(path + "/SF2")
}

func CopyDir(sourceDir string, destDir string) {

}
