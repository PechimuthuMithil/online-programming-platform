## Setup the ubuntu servers
I am using Ubuntu 24,04 (Noble Numbat) LTS server. The following are the steps I followed to setup some important services in the servers. 
`NOTE: This is a currently under-work projet to learn new technologies. So this file will be constantly updated.`
These servers run on VMs in my local machine and the following are steps I followed n my machine.

All the servers inside a NAT Netwrk to enable communication within them, and there are specific port forwarding rules for managing the rest.

### Enable SSH service
I use the SSH service to login into the server terminals from my host machine so I can copy, cut and paste commands. 

1. Open the VM and setup the ssh service by doing the following.
`sudo apt-get update`
`sudo apt-get install -y openssh-server`
`sudo systemctl enable ssh`
`sudo systemctl start ssh`

2. Test service by running `sudo systemctl status ssh`

3. Add port-forwarding rule imilar o below

| Name         | Host IP    | Host Port | Guest IP | Guest Port |  
|--------------|------------|-----------|----------|------------|  
| ssh-server-1 | 127.0.0.2  | 1111      | 10.0.2.4 | 22         |  
| ssh-server-2 | 127.0.0.2  | 1112      | 10.0.2.5 | 22         |  
| ssh-server-3 | 127.0.0.2  | 1113      | 10.0.2.6 | 22         |  

### Disable Swap
1. Stop and disable the ufw service.
`sudo systemctl stop ufw`
`sudo systemctl disable ufw`

2. Comment out the last line in `/etc/fstab`

3. Run `sudo init 6`

4. Check by running `free -m`

