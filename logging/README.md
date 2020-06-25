# Background
I wrote a log package and wanted to see how it stacks up.

# Results
Pretty tasty, to be honest.
```
goos: linux
goarch: amd64
pkg: github.com/jackhascodes/benchmarks/logging
BenchmarkLog_InfoJson
BenchmarkLog_InfoJson-4                 	  247386	      4549 ns/op
BenchmarkLogrus_InfoJson
BenchmarkLogrus_InfoJson-4              	  288896	      4972 ns/op
BenchmarkStandardLog_PrintJson
BenchmarkStandardLog_PrintJson-4        	  321733	      4265 ns/op
BenchmarkLog_InfoJson_withfields
BenchmarkLog_InfoJson_withfields-4      	  198938	      5880 ns/op
BenchmarkLogrus_InfoJson_withfields
BenchmarkLogrus_InfoJson_withfields-4   	  173342	      7047 ns/op
```