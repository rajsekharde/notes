# File transfer between linux machines using scp

## Uploading file to a remote machine
```bash
scp -i <ssh-key-pem-file> <path-to-local-file> <remote-user>@<IP-address>:<upload-path>

scp -i ~/.ssh/key.pem ./file.txt ubuntu@90.92.93.94:/home/ubuntu

# File gets uploaded to /home/ubuntu/ in remote machine
```

## Downloading file from a remote machine
```bash
scp -i <ssh-key-pem-file> <remote-user>@<IP-address>:<file-path> <download-path>

scp -i ~/.ssh/key.pem ubuntu@90.92.93.94:/home/ubuntu/file.txt .

# File- 'file.txt'  gets downloaded from /home/ubuntu/ to local directory
```

## Transferring files between two remote machines
```bash
scp <user1>@<host1>:<path1> <user2>@<host2>:<path2> # Transfer from host1 to host2

scp ubuntu@1.2.3.4:/home/ubuntu/file.txt ubuntu@5.6.7.8:/home/ubuntu
```
