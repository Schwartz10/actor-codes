# Actor codes on testnets 

This repo shows some confusion about loading actor codes. 

```
git clone git@github.com:Schwartz10/actor-codes.git
go get
go run main.go
```

You'll notice the logs:

```
Miner actor code from StateGetActor: bafk2bzacebkjnjp5okqjhjxzft5qkuv36u4tz7inawseiwi2kw4j43xpxvhpm
Miner actor code from GetActorCodeID: bafk2bzacec24okjqrp7c7rj3hbrs5ez5apvwah2ruka6haesgfngf37mhk6us
panic: unknown actor code bafk2bzacebkjnjp5okqjhjxzft5qkuv36u4tz7inawseiwi2kw4j43xpxvhpm

goroutine 1 [running]:
main.main()
        /Users/jonathanschwartz/Documents/glif/actor-codes/main.go:63 +0x4c0
exit status 2
```

How to get the actor codes to reflect calibrationnet?

