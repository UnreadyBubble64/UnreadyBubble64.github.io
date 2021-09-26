export CGO_ENABLED=1
export CFLAGS='-w -s -O3'
export CGO_CFLAGS="$CFLAGS"
export CGO_CPPFLAGS="$CFLAGS"
export CGO_CXXFLAGS="$CFLAGS"
export CGO_FFLAGS="$CFLAGS"
export CGO_LDFLAGS="$CFLAGS"

echo "Очиситка кэша компилятора..."
go clean -cache -testcache -modcache
ccache -C

# read -p 'Пожалуйста, введите номер версии: ' ver

echo 'Компиляция windows-x86 бинарника...'
GOOS=windows GOARCH=386 go build -o ./builds/windows/Infinite-Bomber-x86/infinite-bomber.exe -gcflags="all=-trimpath=$GOPATH" -asmflags="all=-trimpath=$GOPATH" -ldflags="-s -w"

# rcedit.exe ./builds/windows/Infinite-Bomber-x86/infinite-bomber.exe --set-version-string "Infinite Bomber v"$ver "Infinite Bomber" --set-file-version $ver --set-product-version $ver --set-icon ./icon.ico

echo 'Копирование файлов тора...'
mkdir ./builds/windows/Infinite-Bomber-x86/tor
cp -avr ./tor-for-copy/* ./builds/windows/Infinite-Bomber-x86/tor

echo 'Компиляция windows-x64 бинарника...'
GOOS=windows GOARCH=amd64 go build -o ./builds/windows/Infinite-Bomber-x64/infinite-bomber.exe -gcflags="all=-trimpath=$GOPATH" -asmflags="all=-trimpath=$GOPATH" -ldflags="-s -w"

# rcedit.exe ./builds/windows/Infinite-Bomber-x64/infinite-bomber.exe --set-version-string "Infinite Bomber v"$ver "Infinite Bomber" --set-file-version $ver --set-product-version $ver --set-icon ./icon.ico

echo 'Копирование файлов тора...'
mkdir ./builds/windows/Infinite-Bomber-x64/tor
cp -avr ./tor-for-copy/* ./builds/windows/Infinite-Bomber-x64/tor

echo 'Готово!'
