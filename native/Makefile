default: build-dep

# ARCH := arm64
ARCH := amd64
# PREFIX := aarch64-linux-gnu-
# PREFIX := x86_64-w64-mingw32-
# PREFIX :=

CC := clang
# CC := $(PREFIX)gcc
CFLAGS := $(CFLAGS) -O3
CFLAGS += -target $(ARCH)
CFLAGS += -static
CFLAGS += -mavx
CFLAGS += -mavx2
CFLAGS += -mno-bmi
CFLAGS += -mno-red-zone
CFLAGS += -fPIE
CFLAGS += -fno-asynchronous-unwind-tables
CFLAGS += -fno-stack-protector
CFLAGS += -fno-exceptions
CFLAGS += -fno-builtin
CFLAGS += -fno-rtti
CFLAGS += -nostdlib

AR := $(PREFIX)ar
LD := $(PREFIX)ld

gen-asm:
	mkdir -p asm_out
	for fn in *.c; do                           	  \
		$(CC) -S $(CFLAGS) -o asm_out/$${fn}.s -c $${fn}; \
	done

build-dep:
	mkdir -p build
	mkdir -p dist
	for fn in *.c; do                       \
		$(CC) $(CFLAGS) -o build/$${fn}.o -c $${fn} ;\
	done
# liburing compiled with -fPIC
	ld --relocatable --whole-archive -o dist/lnk.o /usr/lib/liburing.a build/*.o
	$(AR) cr dist/libnative-$(ARCH).a dist/lnk.o

build-bin:
	$(CC) -o main lnk.o main.c

clean:
	rm -rf build dist asm_out *.[ao] main a.out