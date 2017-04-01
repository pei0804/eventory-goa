#! /usr/bin/make
#
# Targets:
# - clean     delete all generated files
# - generate  (re)generate all goagen-generated files.
# - build     compile executable
# - ae-build  build appengine
# - ae-dev    deploy to local (dev) appengine
# - ae-deploy deploy to appengine
#
# Meta targets:
# - all is the default target, it runs all the targets in the order above.
#

all: clean generate

dev:
	@which direnv || go get -v github.com/zimbatm/direnv
	@direnv allow
	@which dep || go get -v github.com/zimbatm/dep

clean:
	@rm -rf app
	@rm -rf client
	@rm -rf tool
	@rm -rf swagger

generate:
	@goagen app     -d github.com/tikasan/eventory-goa/design
	@goagen swagger -d github.com/tikasan/eventory-goa/design
	@goagen client  -d github.com/tikasan/eventory-goa/design

model:
	@rm -rf models
	@goagen --design=github.com/tikasan/eventory-goa/design gen --pkg-path=github.com/goadesign/gorma
