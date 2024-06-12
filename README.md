# compilebot

This BOT is in the development phase.

This bot is compile golang code.
<br/>

![demo](/images/demo.png)
<br/>

Full security measures.
<br/>

![demo](/images/demo2.png)

<br/>
Configuration can be done in config.json.
<br/>
The token should contain the discord token.
<br/>
The module should contain a package that can be banned.
<br/>
The maxint can be used to determine the maximum number of for statements to be executed.
<br/>

```json
{
    "maxint": 100,
    "module": [
        "os/exec",
        "net",
        "syscall",
        "os"
    ],
    "token": ""
}
```

<br/>

To install, run `go mod tidy` in the directory of this file.