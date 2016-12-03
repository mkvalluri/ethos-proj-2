package main

import (
	"ethos/efmt"
	"ethos/syscall"
	"ethos/ethos"
	"log"
)

func main () {
	ID := "COPY:"
	path := "/user/"

	fd, status := ethos.OpenDirectoryPath(path)
	if status != syscall.StatusOk {
		log.Fatalf("%v Unable to open path. Error %v\n",ID, status)
	}
	
	//Directory operations	
	fileName := ""

	for fileNames, status := ethos.GetNextName(fd, fileName); status == syscall.StatusOk; fileNames, status = ethos.GetNextName(fd, fileName) {
		if status == syscall.StatusNotFound {
			break
		}

		fileName = string (fileNames)
	
		if fileName == "." || fileName == ".." || fileName == "" {
			continue
		}			

		fileInfo,status  := ethos.GetFileInformation(fd, fileName)
			
		if status != syscall.StatusOk {
			log.Fatalf("%v Some error!. File: %v", ID, fileName)
		}		

		efmt.Println("Fi:", fileName)
		efmt.Println("Fi: IsDirectory: ", fileInfo.FileType)

		if fileInfo.FileType == 1 {
			fileData,_ := ethos.ReadVar(fd, fileName)
			efmt.Println("FILEOUTPUT: ", string(fileData))
			statusWrite := ethos.WriteVar(fd, fileName + "_Output", fileData)
			efmt.Println("Status:",statusWrite)
		}
	}
	syscall.Close(fd)
 	//CopyDir(path, path + "test")
	
}

func CopyDir(sourceDirPath string, destDirPath string) {
	sourcefd, sourceStatus := ethos.OpenDirectoryPath(sourceDirPath)	
	
	if sourceStatus != syscall.StatusOk {
		log.Fatalf("Unable to open source directory: %v\n", sourceDirPath)
	}
		

	//sourceInfo,_ := ethos.GetFileInformation(sourcefd, "") 

	destTypeHash,_ := ethos.TypeNameToHash("ethosGeneratedTypes", "string")
	destStatus := ethos.CreateDirectoryPath(destDirPath, "", destTypeHash)
	
	if destStatus != syscall.StatusOk {
		log.Fatalf("Unable to create dest directory: %v\nStatus: %v\n", destDirPath, destStatus)
	}
	
	destfd, destStatus := ethos.OpenDirectoryPath(destDirPath)
	
	if destStatus != syscall.StatusOk {
		log.Fatalf("Unable to open destination directory: %v\n", destDirPath)
	}

	efmt.Println("Fi:",sourcefd,destfd)

	syscall.Close(sourcefd)
	syscall.Close(destfd)	
}


