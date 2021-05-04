.PHONY: all
all: 
	go run . &> /dev/null
	go tool pprof -png -output cpu.png cpu.out
	go tool pprof -http=":8787" cpu.out
	
	go run .
	go tool pprof -png -output cpu1terminal.png cpuTerminal.out
	go tool pprof -http=":8787" cpuTerminal.out

	go tool pprof -http=":8787" --base cpu.out cpuTerminal.out