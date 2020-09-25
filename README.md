# thinkpad-led

This little tool allows you to manage the red back led of your **Thinkpad**.

## Build

To build the tool, just type the following command:

```bash
go build thinkpad-led.go
```

Optionally, you can put it under `/usr/local/bin` to be available under your `PATH`.

## Usage

```
Usage: # thinkpad-led c̲o̲m̲m̲a̲n̲d̲
Available commands:
on → to enable the led
off → to disable the led
blink → to make it blink
help → show this page
version → show the version
```

## Persistence

To make this tool available at boot, without setup it at each boot, you can automatically load the kernel module at boot.

What you have to do is to create 2 files to allow the kernel to load the correct module, with right parameters.

#### Step 1

Create the entry that will be loaded by `systemd-modules-load.service`.

```
cat << EOF > /etc/modules-load.d/ec_sys.conf
ec_sys
EOF
```

#### Step 2

Add module options under the `/etc/modprobe.d` directory.

```
cat << EOF > /etc/modprobe.d/ec_sys.conf
options ec_sys write_support=1
EOF
```

#### Step 3

Reboot the system.

```
reboot
```

## Credits

I really need to thank [vali20](https://www.reddit.com/user/vali20/) for his amazing post on reddit: https://www.reddit.com/r/thinkpad/comments/7n8eyu/thinkpad_led_control_under_gnulinux/.
