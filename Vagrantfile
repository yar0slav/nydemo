Vagrant.configure("2") do |config|

## Provision VM

    config.vm.define 'minikube' do |machine|
      machine.vm.box = "ubuntu/xenial64"
      machine.disksize.size = "20GB"
      machine.vm.hostname = 'minikube'
      machine.vm.network :private_network,ip: "172.16.10.10"
      machine.vm.provider "virtualbox" do |vbox|
        vbox.gui = false        
        vbox.cpus = 2
        vbox.memory = 8192
      end
      machine.vm.synced_folder "./", "/home/vagrant/nydemo", id: "v-root", mount_options: ["rw", "tcp", "nolock", "noacl", "async"], type: "nfs", nfs_udp: false
  
## Run Ansible playbook to install minikube

      machine.vm.provision "ansible_local" do |ansible|
        ansible.playbook       = "ansible/playbook.yml"
        ansible.version        = "latest"      
        ansible.verbose        = false
        ansible.install        = true
        ansible.limit          = "minikube"      
        ansible.inventory_path = "ansible/hosts"
      end
    end
  end