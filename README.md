Gitmin:
---


### Build and run:
```
make build_linux_only && ./build/gitmin
```
Test it by navigating to
- http://localhost:3030
- http://localhost:3030/healthz


### Using Makefile
```
$ make

help                 Show available options with this Makefile
test                 Run all the tests
clean                Clean the application
vendor               Go get vendor
release              Create a release build.
bench                Benchmark the code.
prof                 Run the profiler.
prof_svg             Run the profiler and generate image.
generate_swagger     Generate swagger definitions from the comments
package              Package the html, css, js files etc
compress             Run upx on the generated binary in `build` directory
build_linux_only     Helper target to quickly build for linux without creating tar
```