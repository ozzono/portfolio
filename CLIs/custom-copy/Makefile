default-tests:
	cd pkg/utils && ./utils.test -test.v && cd -

custom-tests:
	cd pkg/utils && ./utils.test -test.v -s $(s) -d $(d) && cd -

sample:
	sh new_sample.sh $(folder)