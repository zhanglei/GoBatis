pushd cmd\gobatis
go install
@if %errorlevel% equ 1 goto :eof
popd
go generate ./...
@if %errorlevel% equ 1 goto :eof
go test -v ./...