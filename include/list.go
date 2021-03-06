/*
Package include imports all processor packages so that they register with the global
registry. This package can be imported in the main package to automatically register
all of the standard supported apm-server processors.
*/
package include

import (
	// This list is automatically generated by `make imports`
	_ "github.com/elastic/apm-server/processor/error"
	_ "github.com/elastic/apm-server/processor/transaction"
)
