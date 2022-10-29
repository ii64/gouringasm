# gouringasm

Go uring library backed by liburing-2.2 (dda4848a9911120a903bef6284fb88286f4464c9)

compile liburing:
```bash
./configure --nolibc --cc=clang --cxx=clang++
make CFLAGS="-fPIC -static"
make install
```
compile gouringasm:
```bash
make build-native
```


CGo is not Go: https://dave.cheney.net/2016/01/18/cgo-is-not-go