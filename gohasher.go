package main
import (
"crypto/md5"
"crypto/rand"
"golang.org/x/crypto/md4"
"golang.org/x/crypto/ripemd160"
"crypto/sha1"
"crypto/sha256"
"golang.org/x/crypto/sha3"
"golang.org/x/crypto/blake2b"
"golang.org/x/crypto/scrypt"
"golang.org/x/crypto/argon2"
"golang.org/x/crypto/bcrypt"
"fmt"
"log"
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

func generateRandomBytes(n uint32) ([]byte, error) {
    b := make([]byte, n)
    _, err := rand.Read(b)
    if err != nil {
        return nil, err
    }

    return b, nil
}

func banner(){
banner := `
╔═╗┌─┐   ╦ ╦┌─┐┌─┐┬ ┬┌─┐┬─┐
║ ╦│ │───╠═╣├─┤└─┐├─┤├┤ ├┬┘
╚═╝└─┘   ╩ ╩┴ ┴└─┘┴ ┴└─┘┴└─`
cyan := color.New(color.FgCyan).Add(color.Underline)
green := color.New(color.FgGreen).Add(color.Underline)
red := color.New(color.FgRed).Add(color.Underline)
figure.NewFigure("Go-Hasher", "basic", true).Scroll(2000, 200, "left")
fmt.Println(banner)
fmt.Println("\nMade By Gian Paris C. Agsam | Github:https://github.com/FonderElite\n")
cyan.Println("Build: sudo go build <file_name.go>")
green.Println("Run: sudo go run <file_name.go>")
red.Println("Help: sudo go run <file_name.go> -h")
}
func hashstring(strtohash string, typeofhash string) *Hasher{
salt, err := generateRandomBytes(16)
if err != nil{
	log.Fatal(err)
}
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
}else if strings.ToLower(hasher.hashtype) == "ripemd160"{
ripemdhasher := ripemd160.New()
ripemdhasher.Write([]byte(hasher.str_tohash))
ripemd_hashedstr := hex.EncodeToString(ripemdhasher.Sum(nil))
fmt.Printf("RIPEMD160 - %v : %v",hasher.str_tohash, ripemd_hashedstr)
}else if strings.ToLower(hasher.hashtype) == "sha3"{
sha3hasher := sha3.Sum256([]byte(hasher.str_tohash))
fmt.Printf("SHA3 - %v : %x",hasher.str_tohash,sha3hasher)
}else if strings.ToLower(hasher.hashtype) == "blake2b"{
blake2bhasher := blake2b.Sum256([]byte(hasher.str_tohash))
blake2_512_bhasher := blake2b.Sum512([]byte(hasher.str_tohash))
fmt.Printf("BLAKE2B Sum256- %v : %x",hasher.str_tohash,blake2bhasher)
fmt.Printf("BLAKE2B Sum512- %v : %x",hasher.str_tohash,blake2_512_bhasher)
}else if strings.ToLower(hasher.hashtype) == "scrypt"{
scrypthasher, err := scrypt.Key([]byte(hasher.str_tohash), salt, 1<<15, 4, 1, 32)
if err != nil{
log.Fatal(err)
}
fmt.Printf("SCRYPT - %v : %x", hasher.str_tohash, scrypthasher)
}else if strings.ToLower(hasher.hashtype) == "argon2"{
argonhasher := argon2.IDKey([]byte(hasher.str_tohash), salt, 1,62 * 1024,1,32)
fmt.Printf("ARGON2 - %v : %x",hasher.str_tohash,argonhasher)
}else if strings.ToLower(hasher.hashtype) == "bcrypt"{
bcrypthasher,err:=bcrypt.GenerateFromPassword([]byte(hasher.str_tohash), 1)
if err != nil{
log.Fatal(err)
}
fmt.Printf("BCRYPT - %v : %x", hasher.str_tohash, bcrypthasher)
}else{
error_msg := color.New(color.FgYellow, color.Bold)
error_msg.Println("\nOptions: ")
fmt.Println("(md4 | md5 | sha1 | sha256 | sha512 | ripemd160 | sha3 | blake2b | scrypt | argon | bcrypt)")
}
return &hasher
}

func main(){
flag1 := flag.String("s", "", "String to hash")
flag2 := flag.String("ht","","Type of hash(md4,md5,sha1,sha256,ripemd160,sha3, blake2b, script, argon, bcrypt)")
flag.Parse()
banner()
hashstring(*flag1,*flag2)
}
