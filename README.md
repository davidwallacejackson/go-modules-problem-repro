# go-modules-problem-repro

If you load this project in VSCode by way of the workspace file in the repo, you'll see that as you load each of the go files, each module's `go.mod` is updated to contain more dependencies. Eventually, all three `go.mod` files will have the same deps.

I'd expect each module's `go.mod` to contain only its referenced dependencies. Interestingly, if you run `go mod tidy` in any of the three module dirs, the unused dependencies will be stripped out -- but then VSCode, if it's still running, will inject them again.
