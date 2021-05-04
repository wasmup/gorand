.PHONY: all
all: 
	go run . &> /dev/null
	go tool pprof -http=":8787" cpu.out
	go tool pprof -text -output cpu.txt cpu.out
	go tool pprof -png -output cpu.png cpu.out
	go tool pprof -png -output cpu13.png --base ../1/cpu.out cpu.out

clean:
	rm *.out	
	rm *.png