// Package kuboutil provides type conversion utilities for flags.
//
// The code for getting the value of flags usually looks like this
//
// 	v, err := ctx.flag("flag")
// 	if err != nil {
//		return err
// 	}
//
// However, this only returns the value as a string. In order to convert it
// to other types, the strconv package must be used. However, due to ctx.Flag
// having a second return value of error, the value must be converted in a
// seperate statement, which is rather cumbersome.
//
// 	vStr, err := ctx.Flag("flag")
// 	if err != nil {
//		return err
// 	}
// 	v, err := strconv.Atoi(vStr)
// 	if err != nil {
//		return err
// 	}
//
// Therefore, the conversion functions in this package take in both the value
// and an error in order allow flag fetching and conversion to be done in the
// same line.
//
// 	v, err := kuboutil.Int(ctx.Flag("flag"))
// 	if err != nil {
//		return err
// 	}
//
// If ctx.Flag returns an error, it will be returns from the conversion function
// immediately.
package kuboutil
