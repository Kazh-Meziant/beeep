## beeep
`beeep` provides a cross-platform library for sending desktop notifications, alerts and beeps.

### Forked from
[gen2brain/beeep](https://github.com/gen2brain/beeep)
Modification : 
- Windows AppID customization. (PowerShell by default)

### Installation

    go get -u github.com/Kazh-Meziant/beeep

### Build tags

* `nodbus` - disable `godbus/dbus` and use only `notify-send`

### Examples

```go
err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
if err != nil {
    panic(err)
}
```

```go
err := beeep.Notify("Title", "Message body", "assets/information.png", "AppID")
if err != nil {
    panic(err)
}
```

```go
err := beeep.Alert("Title", "Message body", "assets/warning.png", "AppID")
if err != nil {
    panic(err)
}
```


## More

For cross-platform dialogs and input boxes see [dlgs](https://github.com/gen2brain/dlgs).
