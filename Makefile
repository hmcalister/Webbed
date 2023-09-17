dev:
	cd frontend; npm run dev &
	air -tmp_dir "airTMP" --build.cmd "go build -o airTMP/main main.go" --build.bin "./airTMP/main" --build.args_bin "-allowCORS -developmentServer" --build.exclude_dir "frontend"

build:
	cd frontend; npm run build
	go build main.go