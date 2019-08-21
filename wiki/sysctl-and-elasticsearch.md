Elasticsearch needs a vm.max_heap_count of at least 262144 (linux default is 65535).

A description of what this is can be found in the [linux kernel documentation](https://www.kernel.org/doc/Documentation/sysctl/vm.txt), in short is about the number of memory areas that a process can have (not the size of the areas just the amount a process may own) .

This can be set on a running machine with:

```
sysctl -w vm.max_map_count=262144
```

The above will revert to default when the system is rebooted, to persist these settings, do the following:

```
echo 'vm.max_map_count=262144' > /etc/sysctl.d/10-elasticsearch.conf # or /etc/sysctl.conf
```

> With minikube, the host can communicate with the minikube vm via ssh with the command `minikube ssh`
