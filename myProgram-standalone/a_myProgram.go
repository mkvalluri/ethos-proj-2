 package main 

import (
	"ethos/syscall"
	"ethos/ethos"
	"ethos/efmt"
)
func main () {
	//name of the directory to be created
	dirName := "TestDir"
	//calling the function that creates the directory
	SeedData(dirName)
	
}

func SeedData(name string) {
	//ID for the print statements in the log
	ID := "CopyDir:"
	//Log statements to identify the directory created
	efmt.Println(ID,"Creating directory named",name)
	var d1 String
	//command to create the directory at the specifid path
	status := d1.CreateDirectoryPath("/user/" + name, "")
	//check if the directory was successfully created
	if status != syscall.StatusOk {
		efmt.Print("%v Unable to create source directory. Status %v\n", ID, status)
	return
	}
	efmt.Println(ID,"Writing first file")
	d1 = "Hello"
	//Open directory to write
	fd, s := ethos.OpenDirectoryPath("/user/" +name)
	//give error message if the directory could not be opened
	if s != syscall.StatusOk {
	efmt.Print("%v Unable to open source directory. Status: %v\n", ID, s)
	}
	//Write "Hello" into the file F1
	d1.WriteVar(fd,"F1")
	efmt.Println(ID,"Writing second file")
	d1 = "Hi!"
	//Write "Hi!" into the file F2
	d1.WriteVar(fd,"F2")
	// close all files with the file descriptor
	syscall.Close(fd)

	var d2 Uint32
	//create subdirectory called IntDir within this directory
	efmt.Println(ID,"Creating subdirectory named IntDir")
	status = d2.CreateDirectoryPath("/user/" + name + "/IntDir","")
	if status != syscall.StatusOk {
	efmt.Print("%v Unable to open IntDir directory. Status: %v\n", ID, s)
	}

	efmt.Println(ID,"Writing first int file")
	d2 = 123
	//Write the number 123 into the file SF1
	d2.WriteVar(fd, "SF1")
	efmt.Println(ID,"Writing second int file")
	d2 = 456
	//Write the number 456 into the file SF2
	d2.WriteVar(fd, "SF2")
	//close files with the file descriptor
	syscall.Close(fd)
}

