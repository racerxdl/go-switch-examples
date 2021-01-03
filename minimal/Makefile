
export OUTPUT := $(CURDIR)
export PROGRAMNAME := $(notdir $(CURDIR))
export PROGRAM := $(OUTPUT)/$(PROGRAMNAME)

.PHONY: $(BUILD) clean all

all: $(PROGRAM).nro

clean:
	@rm -fr $(PROGRAM).elf $(PROGRAM).nro

$(PROGRAM).nro: $(PROGRAM).elf
	@echo "Building nro"
	@linkle nro $(PROGRAM).elf $(PROGRAM).nro --icon-path=../go-switch.jpg --nacp-path=$(OUTPUT)/nacp.json

$(PROGRAM).elf:
	@echo "Building elf"
	@tinygo build -target nintendoswitch -opt=1 -o $(PROGRAM).elf
