OUT = output.log
N = 50
STRIP = " ns/op" "-4" " " "Benchmark"
BENCHMARKS = InsertRand InsertRep InsertSeq ReadRand ReadRepeat ReadSeq
SPLIT = Insert Read Random Repeated Sequential


all: bench tsv plot

bench:
	number=0 ; while [[ $$number -lt $(N) ]] ; do \
		echo $$number ; \
		for val in $(BENCHMARKS) ; do \
			echo $$val ; \
			go test -run none -bench $$val -benchtime 1s -timeout 0 | grep Benchmark >> $(OUT) ; \
		done ; \
		((number = number + 1)) ; \
	done

tsv:
	for val in $(STRIP); do \
    	sed -i "" "s,$$val,,g" $(OUT) ; \
	done
	
	# Split on sub-bench marker "/"
	sed -i "" 's,/,	,g' $(OUT)

	for val in $(SPLIT); do \
    	sed -i "" "s,$$val,$$val	,g" $(OUT) ; \
	done

plot:
	Rscript plot_bench.r

clean:
	rm $(OUT)
