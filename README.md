huffman-coding
----

huffman encoding / decoding tool

refer: https://en.wikipedia.org/wiki/Huffman_coding

# Usage

```bash
$ go get github.com/KoyamaSohei/huffman-coding
$ echo "Hello, World" > hello.txt 
$ huffman-coding encode hello.txt hello 
$ ls
hello hello.json hello.txt 
$ huffman-coding decode hello 
Hello, World
```