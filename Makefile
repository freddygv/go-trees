OUT = output.log
# TODO: Bump up iterations
N = 3

all: clean bench

bench:
	# TODO: Remove bench filter and bump up benchtime
	# Run benchmarks 
	number=0 ; while [[ $$number -lt $(N) ]] ; do \
		go test -run none -bench InsertSeq10/splay -benchtime 1s -timeout 0 | grep Benchmark >> $(OUT) ; \
		((number = number + 1)) ; \
	done

	# Remove ' ns/op', '\s', '-4' from output
	sed -i "" 's- ns/op--g' $(OUT)
	sed -i "" 's/-4//g' $(OUT)
	sed -i "" 's/ //g' $(OUT)
	
	# Split on /
	sed -i "" 's-/-	-g' $(OUT)

clean:
	rm $(OUT)