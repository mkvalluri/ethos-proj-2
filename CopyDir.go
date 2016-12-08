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

func CleanUp(dirName string) {
	ID := "CopyDir:"

	ethos.RemoveFilePath("/user/" + dirName + "/IntDir/SF1")
	ethos.RemoveFilePath("/user/" + dirName + "/IntDir/SF2")
	fd, s := ethos.OpenDirectoryPath("/user/" + dirName)
	s = ethos.RemoveDirectory(fd, "IntDir")
	efmt.Println(ID, "Remove Int Directory: ", s)

	ethos.RemoveFilePath("/user/" + dirName + "/F2")
	ethos.RemoveFilePath("/user/" + dirName + "/F1")
	fd, s = ethos.OpenDirectoryPath("/user")
	s = ethos.RemoveDirectory(fd, dirName)
	efmt.Println(ID, "Remove Test Directory: ", s)
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

func CopyDir(sourceDirPath string, destDirPath string) {
	ID := "CopyDir:"
	sourcefd, sourceStatus := ethos.OpenDirectoryPath(sourceDirPath)

	if sourceStatus != syscall.StatusOk {
		efmt.Print("%v Unable to open source directory: %v\n", ID, sourceDirPath)
		return
	}

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

	/*for e, s := ethos.GetNextName(sourcefd, elem); s == syscall.StatusOk; e, s = ethos.GetNextName(sourcefd, elem) {
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

		if info.FileType == 1 {
			efmt.Println(ID, "TypeName:", typeName)
			if typeName == "string" {
				var t1 String
				t1.ReadVar(sourcefd, elem)
				efmt.Println(ID, "Data Read:", t1)
				t1.WriteVar(destfd, elem)
			}

			if typeName == "uint32" {
				var t1 Uint32
				t1.ReadVar(sourcefd, elem)
				efmt.Println(ID, "Data Read:", t1)
				t1.WriteVar(destfd, elem)
			}
		}

		if info.FileType == 2 {
			CopyDir(sourceDirPath+"/"+elem, destDirPath+"/"+elem)
		}
	}*/

	syscall.Close(sourcefd)
	syscall.Close(destfd)
}
