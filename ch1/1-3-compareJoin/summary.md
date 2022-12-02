# Experiment Result of benchmark



| array length | Join  | Customed Join |
| ------------ | ----- | ------------- |
| 0            | 2.869 | 0.9405        |
| 1            | 3.318 | 47.93         |
| 10           | 149.5 | 638.1         |
| 100          | 1472  | 12348         |
| 1000         | 13845 | 779353        |



> BenchmarkJoin
> BenchmarkJoin-4           	   81084	     13971 ns/op	    4096 B/op	       1 allocs/op
> BenchmarkCustomedJoin
> BenchmarkCustomedJoin-4   	    1441	    809479 ns/op	 2030678 B/op	    1000 allocs/op

Obviously, the reason why builtin Join function faster is that it perform allocs only once.