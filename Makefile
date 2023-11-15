SUBMOD = agentctl agent-server

push:
	git add .
	git commit -m "update"
	git push origin main


# cert:
# 	openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout server.key -out server.crt
# 	openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout client.key -out client.crt

build:
	@list='$(SUBMOD)'; for subdir in $$list; do \
		echo "make in $$subdir";\
		$(MAKE) -C $$subdir;\
	done	


clean:
	@list='$(SUBMOD)'; for subdir in $$list; do \
		echo "Clean in $$subdir";\
		$(MAKE) -C $$subdir clean;\
	done	

