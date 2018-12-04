package main 
import ("fmt"
		s "strings"				// → has “s”
		)

func main() {
	str := "Hello World"
	str1 := str[0:len(str)]

	str2 := make([]string, 2)
	var str3 []string

	fmt.Println(str1[:1])
	fmt.Println("Contains :",s.Contains(str,"lll"))
	fmt.Println("Count : ",s.Count(str,"l"))
	fmt.Println("HasPrefix :",s.HasPrefix(str, "He"))
	fmt.Println("HasSuffix :",s.HasSuffix(str, "ld"))
	fmt.Println("Index :",s.Index(str, "ll"))

	fmt.Println("Join :",s.Join(str2, "-"))		// → OK with make([]string)
	fmt.Println("Join :",s.Join(str3, "-"))		// → OK with var … []string
	// fmt.Println("Join :",s.Join(str1, "-"))		// → NOK with [low:high]

	fmt.Println("Repeat: ",s.Repeat("a", 5))
	fmt.Println("Replace (\"fooooA\",\"o\",\"0\",2):",s.Replace("fooooA","o","0",2))	// → f00ooA
	fmt.Println("Replace (\"fooooA\",\"o\",\"0\",-1):",s.Replace("fooooA","o","0",-1))	// → f0000A
	

	split := s.Split("a-b-c","-")
	fmt.Println("Split a-b-c:", split)

	fmt.Println("ToUper abc:",s.ToUpper("abc"))
	fmt.Println("ToLower ABC:",s.ToLower("ABC"))
}
