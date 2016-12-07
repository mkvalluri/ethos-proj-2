export UDIR= .
export GOC = x86_64-xen-ethos-6g
export GOL = x86_64-xen-ethos-6l
export ETN2GO = etn2go
export ET2G   = et2g
export EG2GO  = eg2go

export GOARCH = amd64
export TARGET_ARCH = x86_64
export GOETHOSINCLUDE=/usr/lib64/go/pkg/ethos_$(GOARCH)
export GOLINUXINCLUDE=/usr/lib64/go/pkg/linux_$(GOARCH)


install.rootfs = /var/lib/ethos/ethos-default-$(TARGET_ARCH)/rootfs
install.minimaltd.rootfs = /var/lib/ethos/minimaltd/rootfs


.PHONY: all install
all: a_myProgram

install: a_myProgram
	ethosTypeInstall $(install.rootfs) $(install.minimaltd.rootfs) myType
	install a_myProgram $(install.rootfs)/programs
	install a_myProgram $(install.minimaltd.rootfs)/programs
	echo -n /programs/a_myProgram | ethosStringEncode > $(install.rootfs)/etc/init/console
	mkdir -p $(install.rootfs)/user/nobody/myDir
	setfattr -n user.ethos.typeHash -v $(shell egPrint myType hash MyType) $(install.rootfs)/user/nobody/myDir

myType.go : myType.t
	$(ETN2GO) . myType main $^

a_myProgram: a_myProgram.go myType.go
	ethosGo $^ 

clean:
	rm -f myType.go
	rm -rf myType/ myTypeIndex
	rm -f a_myProgram
	rm -f a_myProgram.goo.ethos
