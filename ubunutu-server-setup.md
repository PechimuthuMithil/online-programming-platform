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

### Install Docker
This can be chosen to install during the the installation of the soperating system.

### Install Kubernetes
Like docker microK8s can be installed during the installation of the soperating system, however, I install Kubernetes separately using the following steps.

#### Download and add GPG key
The GPG key (GNU Privacy Guard key) is used to ensure the integrity and authenticity of the packages you download and install from a repository.   
`curl -fsSL https://pkgs.k8s.io/core:/stable:/v1.30/deb/Release.key | sudo gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg`

#### Add Software Repositories and Install K8s Tools
Kubernetes is not included in the default repositories. To add them, enter the following:
`echo 'deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v1.30/deb/ /' | sudo tee /etc/apt/sources.list.d/kubernetes.list`  
`sudo at-get update`  
`sudo apt-get install -y kubelet kubeadm kubectl`  
`sudo apt-mark hold kubelet kubeadm kubectl`    

