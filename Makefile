g:
	git add .
	git commit -m"自动提交 git 代码"
	git push

dev:
	make build && make run

build:
	protoc -I . --go_out=plugins=micro:. proto/goods/goods.proto
	protoc -I . --go_out=plugins=micro:. proto/taxcode/taxcode.proto
	protoc -I . --go_out=plugins=micro:. proto/unspsc/unspsc.proto
	protoc -I . --go_out=plugins=micro:. proto/category/category.proto
	protoc -I . --go_out=plugins=micro:. proto/department/department.proto
	protoc -I . --go_out=plugins=micro:. proto/firm/firm.proto
	protoc -I . --go_out=plugins=micro:. proto/brand/brand.proto
