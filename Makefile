build-solution:
	docker build -t course .
.PHONY: build-solution

run:
	docker run -it course
.PHONY: run