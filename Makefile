OUT = output.log
N = 50

all: bench tsv plot

bench:
	number=0 ; while [[ $$number -lt $(N) ]] ; do \
		echo $$number ; \
		for val in InsertRand InsertRep InsertSeq ReadRand ReadRepeat ReadSeq ; do \
			echo $$val ; \
			go test -run none -bench $$val -benchtime 1s -timeout 0 | grep Benchmark >> $(OUT) ; \
		done ; \
		((number = number + 1)) ; \
	done

tsv:
	# Remove ' ns/op', '-4', '\s', 'Benchmark' from output
	# TODO: Consolidate these
	sed -i "" 's- ns/op--g' $(OUT)
	sed -i "" 's/-4//g' $(OUT)
	sed -i "" 's/ //g' $(OUT)
	sed -i "" 's/Benchmark//g' $(OUT)

	# Split on /
	sed -i "" 's-/-	-g' $(OUT)

	# Split operations and input types into a tab-delimited field
	for val in Insert Read Random Repeated Sequential; do \
    	sed -i "" "s/$$val/$$val	/g" $(OUT) ; \
	done

plot:
	Rscript plot_bench.r

clean:
	rm $(OUT)