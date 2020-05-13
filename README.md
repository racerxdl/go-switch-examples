# Golang Nintendo Switch Examples

Examples using TinyGo for Nintendo Switch

# Requirements

*   DevKitPro for Switch (see [Setting up Development Environment](https://switchbrew.org/wiki/Setting_up_Development_Environment))
*   TinyGo with Nintendo Switch Support (https://github.com/racerxdl/tinygo on nintendoswitch branch)
*   Linkle (see [https://github.com/MegatonHammer/linkle](https://github.com/MegatonHammer/linkle))

# Building a example

Go to the example folder and then run:

```bash
tinygo build -target nintendo-switch -o hello-world.elf
linkle nro hello-world.elf hello-world.nro
```

That should generate a `nro` file that runs on [Yuzu](https://yuzu-emu.org/) or [Atmosphere](https://github.com/Atmosphere-NX/Atmosphere/)

If you have linkle and tinygo in path, you can run `make`