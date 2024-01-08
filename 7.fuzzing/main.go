

package main
import ("fmt"
"errors"
"unicode/utf8"
)

func Reverse(s string)(string, error){
	if !utf8.ValidString(s) {
        return s, errors.New("input is not valid UTF-8")
    }
	// ret := []byte(s)
	// ret := []rune(s)
	fmt.Printf("input: %q\n", s)
    ret := []rune(s)
    fmt.Printf("runes: %q\n", ret)
	for i,j := 0, len(ret) -1; i < len(ret ) / 2; i,j = i+1, j-1{
		ret[i], ret[j] = ret[j], ret[i]
	}
	return string(ret), nil
}

func main(){
	s := "function main is undeclared in the main package"
	rev,revErr := Reverse(s)
	drev,drevErr := Reverse(rev) 
	fmt.Printf("origin %q\n",s)
	fmt.Printf("reverse %q, err: %v \n", rev, revErr)
	fmt.Printf("dreverse %q, err: %v \n", drev, drevErr)
}



