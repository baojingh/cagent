push:
	git add .
	git commit -m "update"
	git push origin main


cert:
	openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout server.key -out server.crt
	openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout client.key -out client.crt