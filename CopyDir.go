package main

import (
	"ethos/efmt"
	"ethos/syscall"
	"ethos/ethos"
	"log"
	"os"
	"io"
	"ethos/kernelTypes"
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
			var s kernelTypes.String = kernelTypes.String(string(fileData))
			s.Write(fd)
			//statusWrite := ethos.WriteVar(fd, fileName + "_Output", fileData)
			//efmt.Println("Status:",statusWrite)
			//cp(path+"_Output", path+fileName)
			efmt.Println("Copied file with status: ")
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

func cp(dst, src string) {
	s, err := os.Open(src)
	if err != nil {
		efmt.Println("1: ",err)
		return //err
	}
	// no need to check errors on read only file, we already got everything
	// we need from the filesystem, so nothing can go wrong now.
	defer s.Close()
	d, err := os.Create("/tmp/newnew")
	if err != nil {
		efmt.Println("2: ",err)
		return
	}
	if _, err := io.Copy(d, s); err != nil {
		d.Close()
		efmt.Println("3: ", err)
		return
	}
	efmt.Println("4: ", d.Close())
	return //d.Close()
}


