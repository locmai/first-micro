config-local:
	kubectx docker-desktop

consul:
	$(MAKE) config-local
	helm install stable/consul -f consul.yaml --name consul --wait

db:
	$(MAKE) config-local
	helm install stable/postgresql -f postgresql.yaml --name postgresql --wait

nats:
	$(MAKE) config-local
	helm install stable/nats -f nats.yaml --name nats --wait