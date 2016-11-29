export UDIR = .
export GOC = x86_64-xen-ethos-6g
export GOL = x86_64-xen-ethos-6l
export ETN2GO = etn2go
export ET2G = et2g
export EG2GO = eg2go

export GOARCH = amd64
export TARGET_ARCH = x86_64
export GOETHOSINCLUDE = /usr/lib64/go/pkg/ethos_$(GOARCH)
export GOLINUXINCLUDE = /usr/lib64/go/pkg/linux_$(GOARCH)

install.rootfs = /var/lib/ethos/ethos-default-$(TARGET_ARCH)/rootfs
install.minimaltd.rootfs = /var/lib/ethos/minimaltd/rootfs

.PHONY: all install

all: copyDir

install: copyDir
	install CopyDir $(install.rootfs)/programs
	install CopyDir $(install.minimaltd.rootfs)/programs
	echo -n /programs/CopyDir | ethosStringEncode > $(install.rootfs)/etc/init/console

copyDir: CopyDir.go
	ethosGo $^

clean:
	rm -f CopyDir
	rm -f CopyDir.goo.ethos
