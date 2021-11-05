window ip: 192.168.88.153
linux ip : 192.168.80.132

linux运行fileserver

  ./fileserver /root 1314

访问

  http://localhost:1314
  http://192.168.80.1:1314
  http://192.168.88.153:1314

windows运行fileserver

  fileserver.exe D:\soft 1315

访问

  curl -s 192.168.88.153:1315
  curl -s 192.168.80.1:1315