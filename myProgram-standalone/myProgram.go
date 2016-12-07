 package main 

import (
	"ethos/syscall"
	"ethos/ethos"
	"ethos/efmt"
)

func main () {
	//path := "/user/"
	dirName := "TestDir"
	SeedData(dirName)
	
}

func SeedData(name string) {
	ID := "CopyDir:"
	efmt.Println(ID,"Creating directory named",name)
	var d1 String
	status := d1.CreateDirectoryPath("/user/" + name, "")
	if status != syscall.StatusOk {
		efmt.Print("%v Unable to create source directory. Status %v\n", ID, status)
	return
	}
	efmt.Println(ID,"Writing first file")
	d1 = "Hello"

	fd, s := ethos.OpenDirectoryPath("/user/" +name)
	if s != syscall.StatusOk {
	efmt.Print("%v Unable to open source directory. Status: %v\n", ID, s)
	}
	d1.WriteVar(fd,"F1")
	efmt.Println(ID,"Writing second file")
	d1 = "Hi!"
	d1.WriteVar(fd,"F2")
	syscall.Close(fd)

	var d2 Uint32
	efmt.Println(ID,"Creating subdirectory named IntDir")
	status = d2.CreateDirectoryPath("/user/" + name + "/IntDir","")
	if status != syscall.StatusOk {
	efmt.Print("%v Unable to open IntDir directory. Status: %v\n", ID, s)
	}

	efmt.Println(ID,"Writing first int file")
	d2 = 123
	d2.WriteVar(fd, "SF1")
	efmt.Println(ID,"Writing second int file")
	d2 = 456
	d2.WriteVar(fd, "SF2")
	syscall.Close(fd)
}

