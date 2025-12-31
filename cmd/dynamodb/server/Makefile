main_package_path = ./src
binary_name = todo

.PHONY: run
run:
	gow run ${main_package_path}

.PHONY: build
build:
    # Include additional build steps, like TypeScript, SCSS or Tailwind compilation here...
	go build -o=/tmp/bin/${binary_name} ${main_package_path}