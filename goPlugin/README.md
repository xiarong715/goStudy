cd plugin
go build -buildmode=plugin -o testPlugin.so

cd ..
go build
./goplugin