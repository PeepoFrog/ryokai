# Bootstrap
The script provides the initial setup for the localhost on which it runs.

What's installed:
1. apt-transport-https  
2. ca-certificates 
3. curl 
4. software-properties-common 
5. jq
6. docker

## Ansible runner
The final step of the initialization is building the Ansible runner image. It contains the Python `docker` module, so it can orchestrate containers.

Main purpose: To simplify the further deployment of infrastructure.

## How to use:

```bash
docker run --rm -it -v $(pwd)/src:/ansible/src ansible-runner:v*.*. ansible-playbook -i /ansible/src/inventory.ini /ansible/src/playbook.yml
```

* `--rm` - Deletes container after run.
* `--it` - Runs interactively, displaying execution logs.
*   `-v` - Attaches a volume with the playbook and inventory file.
* `ansible-runner:v*.*.*` - The image name. Replace `*.*.*` with the speific version you are using.
* `ansible-playbook` - Command to run Ansible playbook.
* `-i` - Flag that specifies the path in the inventory file.
* `/ansible/src/playbook.yml` - Path to the playbook file.



