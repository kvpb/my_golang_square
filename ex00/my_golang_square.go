

package main

import (
	"os" // ARGS
	//"unicode" // IsDigit
	"strconv" // AToI
	"fmt" // PRINTF, PRINTLN
)

func IsNumber( s string ) bool {
	for i, c := range s {
		if i == 0 && ( ( c <= 48 || 57 <= c ) || c != 43 || c != 45 ) { //if i == 0 && ( ! ( 48 <= c && c <= 57 ) || c != 43 || c != 45 ) {
			return false
		}
		if i != 0 && ( c <= 48 || 57 <= c ) { //if i != 0 && ! ( 48 <= c && c <= 57 ) {
			return false
		}
	}
	return true
}

func sub( w /*u*/int/*64*/, h /*u*/int/*64*/ ) {
	var i /*u*/int/*64*/ = 1
	var j /*u*/int/*64*/ = 1

	for j <= h {
		if j == 1 || j == h {
			for 1 <= i && i <= w {
				if i == 1 {
					fmt.Printf("%c", 'o')
				} else if 1 < i && i < w {
					fmt.Printf("%c", '-')
				} else /*if i == w*/ {
					fmt.Println("o")
				}
				i++
			}
			i = 1
		} else if 1 < j && j < h {
			for 1 <= i && i <= w {
				if i == 1 {
					fmt.Printf("%c", '|')
				} else if 1 < i && i < w {
					fmt.Printf("%c", ' ')
				} else /*if i == w*/ {
					fmt.Println("|")
				}
				i++
			}
			i = 1
		}
		j++
	}
}

func main(/* c_a int, v_a []string */) {
	var c_a int = len( os.Args[ 1:] )
	var v_a []string = os.Args[ 1:]
	var w /*u*/int/*64*/ = 0
	var h /*u*/int/*64*/ = 0
	//fmt.Printf("%d", c_a )
	//fmt.Printf("%s", v_a )

	if c_a != 2 || v_a[ 0 ] == "" || v_a[ 1 ] == "" { //if c_a <= 2 || len( v_a[ 1 ] ) == 0 || len( v_a[ 2 ] ) == 0 {
		return
	} else if c_a != 2 || IsNumber( v_a[ 0 ] ) || IsNumber( v_a[ 1 ] ) {
		return
	} else {
		w, _ = strconv.Atoi( v_a[ 0 ] )
		h, _ = strconv.Atoi( v_a[ 1 ] )
		sub( w, h )
		return
	}
 }
 
 