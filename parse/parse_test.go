/*
 * Copyright 2015 Ming-der Wang <ming@log4analytics.com> All rights reserved.
 * Licensed under MIT License.
 */

package parse

import ()

var (
	src = `
/*
 * Copyright 2015 Ming-der Wang <ming@log4analytics.com> All rights reserved.
 * Licensed under MIT License.
 */

package parse
import ()

// @generate_helloworld
const hello = "World"
`
)

func ExampleParse() {
	Scan(src, "")
	//Output:
	//"World"
}
