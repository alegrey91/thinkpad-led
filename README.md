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

To make this tool available at boot (without setup it each time), you can automatically load the kernel module using the following commands.

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

## Use without sudo password

**(Do at your own risk)**

To avoid to input the `sudo` password each time, you can set a specific capability. Of course, you have to make it more secure, to avoid some kind of exploitation.

First of all, you have to change the binary ownership to `root`.

```bash
sudo chown root:root thinkpad-led
```

Then, you have to reduce access to other users, giving the binary non-write permission.

```bash
sudo chmod 511 thinkpad-led
```

Finally, you have to set capability to enable access to kernel file.

```bash
sudo setcap CAP_DAC_OVERRIDE=ep thinkpad-led
```

Now, you can use `thinkpad-led` without input the `sudo` password.

## Credits

I really need to thank [vali20](https://www.reddit.com/user/vali20/) for his amazing post on reddit: https://www.reddit.com/r/thinkpad/comments/7n8eyu/thinkpad_led_control_under_gnulinux/.
