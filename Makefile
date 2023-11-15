SUBMOD = agentctl agent-server

all: clean build

push-hub:
	git remote set-url origin  git@github.com:baojingh/cagent.git
	git add .
	git commit -m "update"
	git push origin main


push-lab:
	git remote set-url origin git@gitlab.com:localdetector/cagent.git
	git add .
	git commit -m "update"
	git push origin main	



# cert:
# 	openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout server.key -out server.crt
# 	openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout client.key -out client.crt

build:
	@list='$(SUBMOD)'; for subdir in $$list; do \
		echo "\033[32mMake in $$subdir \033[0m";\
		$(MAKE) -C $$subdir;\
		echo "\n";\
	done	


clean:
	@list='$(SUBMOD)'; for subdir in $$list; do \
		echo "\033[32mClean in $$subdir \033[0m";\
		$(MAKE) -C $$subdir clean;\
		echo "\n";\
	done	

