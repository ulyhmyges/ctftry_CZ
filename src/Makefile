log:
	@docker -f ctf_api

restart:
	@echo "Restarting the app: ctf_api..."
	@docker compose down
	@docker compose rmi -f ctf_api
	@docker compose up -d