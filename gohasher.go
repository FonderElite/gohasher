package main
import (
"crypto/md5"
"golang.org/x/crypto/md4"
"crypto/sha1"
"crypto/sha256"
"fmt"
"flag"
"github.com/fatih/color"
"github.com/common-nighthawk/go-figure"
"encoding/hex"
"strings"
)

type Hasher struct {
str_tohash string
hashtype string
}
func banner(){
banner := `
╔═╗┌─┐   ╦ ╦┌─┐┌─┐┬ ┬┌─┐┬─┐
║ ╦│ │───╠═╣├─┤└─┐├─┤├┤ ├┬┘
╚═╝└─┘   ╩ ╩┴ ┴└─┘┴ ┴└─┘┴└─`
cyan := color.New(color.FgCyan).Add(color.Underline)
green := color.New(color.FgGreen).Add(color.Underline)
red := color.New(color.FgRed).Add(color.Underline)
figure.NewFigure("Go-Hasher", "basic", true).Scroll(1000, 200, "left")
fmt.Println(banner)
fmt.Println("\nMade By Gian Paris C. Agsam | Github:https://github.com/FonderElite\n")
cyan.Println("Build: sudo go build <file_name.go>")
green.Println("Run: sudo go run <file_name.go>")
red.Println("Help: sudo go run <file_name.go> -h")
}
func hashstring(strtohash string, typeofhash string) *Hasher{
hasher := Hasher{str_tohash:strtohash, hashtype:typeofhash}
if strings.ToLower(hasher.hashtype) == "md4"{
md4hasher := md4.New()
md4hasher.Write([]byte(hasher.str_tohash))
md4_hashedstr :=  hex.EncodeToString(md4hasher.Sum(nil))
fmt.Printf("MD4 - %v : %v",hasher.str_tohash, md4_hashedstr)
}else if strings.ToLower(hasher.hashtype) == "md5"{
md5hasher := md5.New()
md5hasher.Write([]byte(hasher.str_tohash))
md5_hashedstr :=  hex.EncodeToString(md5hasher.Sum(nil))
fmt.Printf("MD5 - %v : %v",hasher.str_tohash, md5_hashedstr)
}else if strings.ToLower(hasher.hashtype) == "sha1"{
sha1hasher := sha1.New()
sha1hasher.Write([]byte(hasher.str_tohash))
sha1_hashedstr := hex.EncodeToString(sha1hasher.Sum(nil))
fmt.Printf("SHA1 - %v : %v",hasher.str_tohash, sha1_hashedstr)
}else if strings.ToLower(hasher.hashtype) == "sha256"{
sha256hasher := sha256.New()
sha256hasher.Write([]byte(hasher.str_tohash))
sha256_hashedstr := hex.EncodeToString(sha256hasher.Sum(nil))
fmt.Printf("SHA256 - %v : %v",hasher.str_tohash, sha256_hashedstr)
}else{
error_msg := color.New(color.FgYellow, color.Bold)
error_msg.Println("\nOptions: ")
fmt.Println("(md4 | md5 | sha1 | sha256)")
}
return &hasher
}
func main(){
flag1 := flag.String("s", "", "String to hash")
flag2 := flag.String("ht","","Type of hash(md4,md5,sha1,sha256)")
flag.Parse()
banner()
hashstring(*flag1,*flag2)
}
