## 介绍：
`ft`是一个实现在两台主机间互传文件的应用。
## 使用说明：
```bash
# setup initial directory, received file would be put there
mkdir -p $HOME/ft
make config
./config --default-directory $HOME/ft
cd $HOME/ft

# start client on your local machine
make client
# connect with remote server
./ft-client-daemon
# list usable ip addrs
./ft-client-list
# send a file to some of the usable ip addrs 
./ft-client-sender $FILENAME ip1 ip2...
# check out sended or received file
./ft-client-log

# start server on your remote machine
make server
# connect with client
./ft-server-daemon
```
