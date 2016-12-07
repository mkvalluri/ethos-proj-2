/*****************************
Name: CopyDir.go
Description: Implemenation of Unix cp functionality.
Authors: Murali Valluri (mvallu2@uic.edu), Samrudhi Nayak (snayak6@uic.edu)
Date: 12/08/2016
*****************************/


package main

import (
	"ethos/efmt"
	"ethos/ethos"
	"ethos/syscall"
)

func main() {
	dirName := "TestDir"
	CleanUp(dirName)
	CleanUp(dirName + "_Copy")
	SeedData(dirName)
	CopyDir("/user/"+dirName, "/user/"+dirName+"_Copy")
}
//used to remove the previously copied directory sub tree
func CleanUp(dirName string) {
	ID := "CopyDir:"

	ethos.RemoveFilePath("/user/" + dirName + "/IntDir/SF1")
	ethos.RemoveFilePath("/user/" + dirName + "/IntDir/SF2")
	fd, s := ethos.OpenDirectoryPath("/user/" + dirName)
	s = ethos.RemoveDirectory(fd, "IntDir")
	efmt.Println(ID, "Remove Int Directory: ", s)

	ethos.RemoveFilePath("/user/" + dirName + "/TestTypeDir/SF3")
	s = ethos.RemoveDirectory(fd, "TestTypeDir")
	efmt.Println(ID, "Remove TestType Directory: ", s)

	ethos.RemoveFilePath("/user/" + dirName + "/F2")
	ethos.RemoveFilePath("/user/" + dirName + "/F1")
	fd, s = ethos.OpenDirectoryPath("/user")
	s = ethos.RemoveDirectory(fd, dirName)
	efmt.Println(ID, "Remove Test Directory: ", s)
	syscall.Close(fd)
}
//used to create a new directory subtree with 2 files and two sub directories, one which has two int files in it and another which has one file of type TestType
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

	path1 := path + "/IntDir"
	efmt.Println(ID, "Writing first int file")
	d2 = 456
	d2.WriteVar(path1 + "/SF1")

	efmt.Println(ID, "Writing second int file")
	d2 = 555
	d2.WriteVar(path1 + "/SF2")

	var d3 TestType
	efmt.Println(ID, "Creating subdirectory named TestTypeDir")
	status = d3.CreateDirectory(path+"/TestTypeDir", "")
	if status != syscall.StatusOk {
		efmt.Print("%v Unable to create TestTypeDir. Status: %v\n", ID, status)
	}

	path2 := path + "/TestTypeDir"
	efmt.Println(ID, "Writing first TestType file")
	d3.F1 = "Some String"
	d3.F2 = 656
	d3.WriteVar(path2 + "/SF3")
}
//function to copy directory structure
func CopyDir(sourceDirPath string, destDirPath string) {
	ID := "CopyDir:"
	sourcefd, sourceStatus := ethos.OpenDirectoryPath(sourceDirPath)

	if sourceStatus != syscall.StatusOk {
		efmt.Print("%v Unable to open source directory: %v\n", ID, sourceDirPath)
		return
	}
	//get type hash of the source directory to create new directory with same type hash
	sourceInfo, _ := ethos.GetFileInformation(sourcefd, "")
	_, typeName, _ := ethos.TypeHashToName(sourceInfo.TypeHash)

	destStatus := ethos.CreateDirectoryPath(destDirPath, "", sourceInfo.TypeHash)
	if destStatus != syscall.StatusOk {
		efmt.Print("%v Unable to create destination directory named %v. Status: %v\n", ID, destDirPath, destStatus)
		return
	}

	destfd, destStatus := ethos.OpenDirectoryPath(destDirPath)
	if destStatus != syscall.StatusOk {
		efmt.Print("%v Unable to open destination directory: %v\n", ID, destDirPath)
		return
	}

	elem := ""
	efmt.Println(ID, "Looping through each file in source")

	for e, s := ethos.GetNextName(sourcefd, elem); s == syscall.StatusOk; e, s = ethos.GetNextName(sourcefd, elem) {
		if s == syscall.StatusNotFound {
			break
		}

		elem = string(e)
	
		if elem == "." || elem == ".." || elem == "" {
			continue
		}

		efmt.Println(ID, "Element:", elem)
		info, status := ethos.GetFileInformation(sourcefd, elem)
		if status != syscall.StatusOk {
			efmt.Println("%v Unable to get status for file %v. Status %v\n", ID, elem, status)
			continue
		}
		efmt.Println(ID, "TypeName:", typeName)
		//if the element encountered is a file(FileType=1)		
		if info.FileType == 1 {
			efmt.Println(ID, "TypeName:", typeName)
			//if the file is a string file
			if typeName == "string" {
				var t1 String
				t1.ReadVar(sourceDirPath + "/" + elem)
				efmt.Println(ID, "Data Read:", t1)
				t1.WriteVar(destDirPath + "/" + elem)
			}
			//if the file is an integer file
			if typeName == "uint32" {
				var t1 Uint32
				t1.ReadVar(sourceDirPath + "/" + elem)
				efmt.Println(ID, "Data Read:", t1)
				t1.WriteVar(destDirPath + "/" + elem)
			}
			//if the file is of type TestType
			if typeName == "TestType" {
				var t1 TestType
				t1.ReadVar(sourceDirPath + "/" + elem)
				efmt.Println(ID, "Data Read:", t1)
				t1.WriteVar(destDirPath + "/" + elem)
			}
		}
		//if the element is a directory(FileType=2), recursively call CopyDir to copy each file
		if info.FileType == 2 {
			CopyDir(sourceDirPath+"/"+elem, destDirPath+"/"+elem)
		}
	}
	//close the source and destination directories
	syscall.Close(sourcefd)
	syscall.Close(destfd)
}
