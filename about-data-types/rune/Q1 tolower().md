# Q1. tolower()

Level: Very Easy
My Expertise: Easy
No. of times practiced: 1
Problem Type: Assignment
Done: Yes
Count Connection: Problem Count (https://www.notion.so/Problem-Count-5a9398227c8349cd8360893133fb7dd1?pvs=21)
Module Name: Introduction to Problem Solving (Intermediate) 2
Module Number: Module-4
Site: Scaler
Topics: Intermediate DSA : Strings (https://www.notion.so/Intermediate-DSA-Strings-91bdb8e64f4a41dd8f700aef85bfb1f2?pvs=21)

[In this Blog he told to us why we use UTF-8,Unicode character.](https://www.joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/)

In this Blog he told to us why we use UTF-8,Unicode character.

<aside>
🧘🏿‍♂️ Based on above article

- why we put **content type** and **meta tag** in our html page because then only the browser can understand which encoding which will use to render this page.
- Example
    
    > **Content-Type: text/plain; charset="UTF-8"**
    > 
    
    > **<html>**
    > 
    > 
    > **<head>**
    > 
    > **<meta http-equiv="Content-Type" content="text/html; charset=utf-8">**
    > 
    
- **byte** ⇒ Ascii ⇒ support only English alphabets which not able to render other language as well as type ⇒ uint8,range of a byte is **0 to 255**.
- Unicode ⇒ but Unicode  able to render all language.
- [UTF-8](http://www.utf-8.com/) ⇒ is a way of encoding the Unicode characters.
- **Rune use** Unicode as well as larger range than byte**.**
- **Explanation-Example**
    
    ```go
    package main
    
    import "fmt"
    
    func main() {
    	// byte example
    	var b byte = 'A'
    	fmt.Printf("Byte: %c\n", b)
    
    	// rune example
    	var r rune = '東' // This word isn't able to read form "byte" but rune can.
    	fmt.Printf("Rune: %c\n", r)
    }
    ```
    
</aside>

<aside>
🧘🏿‍♂️ what is the Difference between Array and Silice?

![Untitled](Q1%20tolower()%205af8809ffc1b45b3884746b688851cd8/Untitled.png)

> Array
> 
- Array has a fixed size.

> Slice
> 
- Different btw while iterating string in For loop vs For range
    - Reference
        
        [The difference between bytes and runes in Go](https://www.educative.io/answers/the-difference-between-bytes-and-runes-in-go)
        
    - String is a readable only slice of bytes.
    - when we use **For loop** in string ,we will get every character as a byte.
    - But when we use **For range loop** in string ,we will get every character as a rune.
- Slice-Header
    
    ```go
    package main
    
    import (
    	"fmt"
    )
    
    type sliceHeader struct {
    	Length        int
    	ZerothElement *byte
    }
    
    func main() {
    	var arr [5]byte
    	arr = [5]byte{11, 19, 50, 100, 87}
    
    	slice2 := sliceHeader{
    		Length:        5,
    		ZerothElement: &arr[4],
    	}
    
    	fmt.Println("slice2", slice2)
    
    	for i := 0; i < slice2.Length; i++ {
    		element := arr[i]
    		fmt.Printf("%d ", element)
    	}
    
    }
    ```
    
- inbuild function **Copy**
    - Here we can only copy the value to the destination, it won’t append all the elements,it only copies the data from the source
    
    ```go
    package main
    
    import "fmt"
    
    func main() {
        src := []int{1, 2, 3, 4, 5}
        dst := make([]int, len(src))
    
        // Copy elements from src to dst
        numCopied := copy(dst, src)
    
        fmt.Printf("Number of elements copied: %d\n", numCopied)
        fmt.Println("Source slice:", src)
        fmt.Println("Destination slice:", dst)
    }
    ```
    
</aside>

<aside>
💡 Default type of Single Quote

- The default type of the variable **`myData := 'a'`** will be **`int32`**. When you use a single-quoted character literal, it is interpreted as a **`rune`**, and **`rune`** is an alias for **`int32`**.
</aside>

<aside>
🧘🏿‍♂️ String

- String are just read-only slices of bytes.
- Go-Strings
    - **String Immutability:** Like Java, Go strings are immutable, meaning that once you create a string, you cannot change its contents. Any operation that appears to modify a string actually creates a new string with the desired changes.
    
</aside>

<aside>
🧘🏿‍♂️ Ascii Codes

![Untitled](Q1%20tolower()%205af8809ffc1b45b3884746b688851cd8/Untitled%201.png)

</aside>