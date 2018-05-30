OUT = output.log
# TODO: Bump up iterations
N = 1

all: clean bench

bench:
	# TODO: Remove bench filter and bump up benchtime
	# Run benchmarks 
	number=0 ; while [[ $$number -lt $(N) ]] ; do \
		go test -run none -bench InsertSequential100000/splay -benchtime 1s -timeout 0 | grep Benchmark >> $(OUT) ; \
		((number = number + 1)) ; \
	done

	# Remove ' ns/op', '-4', '\s', 'Benchmark' from output
	# TODO: Consolidate these
	sed -i "" 's- ns/op--g' $(OUT)
	sed -i "" 's/-4//g' $(OUT)
	sed -i "" 's/ //g' $(OUT)
	sed -i "" 's/Benchmark//g' $(OUT)

	# Split operations and input types into a tab-delimited field
	for val in Insert Read Random Repeated Sequential; do \
    	sed -i "" "s/$$val/$$val	/g" $(OUT) ; \
	done
	
	# Split on /
	sed -i "" 's-/-	-g' $(OUT)

clean:
	rm $(OUT)