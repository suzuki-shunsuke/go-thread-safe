/*
Package safe provides API for thread safe data operation.
safe provides some struct which has a data internally.
These structs have some methods to do thead safe operation to their internal data.
Internally sync.RWMutex is used for thread safe operation.

The methods whose name ends with `Unsafe` operates internal data without lock,
which means these methods aren't thread safe.
We should use these methods carefully.
Note that they don't have nothing to do with the standard library "unsafe".
*/
package safe
