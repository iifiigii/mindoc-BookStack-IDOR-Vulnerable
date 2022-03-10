## Vulnerable Detail

MinDoc and BookStack are document management system. They both develop their web services based on Beego. Besides, BookStack is based on MinDoc. They have same IDOR vulnerables.

### Execution track
All function are implemented by different **Controller**. **BaseController** (like Parent class) run the **Prepare()** at every beginning.

![](https://github.com/iifiigii/mindoc-BookStack-IDOR-Vulnerable/blob/gh-pages/1.png)

In **Prepare()**, if applications do not get session information, they try to use cookie to access account.

![](https://github.com/iifiigii/mindoc-BookStack-IDOR-Vulnerable/blob/gh-pages/2.png)


### PoC


