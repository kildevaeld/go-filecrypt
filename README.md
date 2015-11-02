# go-filecrypt
Streaming file encryption based on nacl/secretbox

# Usage:

```go


key := filecrypt.Key([]byte("some super secret key"))

file, _ := os.Open("main.go")
defer file.Close()

encrypted, _ := os.Create("main.vau")
defer file.Close()
filecrypt.Encrypt(encrypted, file, &key)

encrypted.Seek(0, 0)

buf := bytes.NewBuffer(nil)
filecrypt.Decrypt(buf, encrypted, &key)

fmt.Printf("%v", string(buf.Bytes()))
```
