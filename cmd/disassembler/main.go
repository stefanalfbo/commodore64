package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/stefanalfbo/commodore64/cpu6510"
)

func main() {
	filePath := flag.String("file", "", "Path to file (reads from stdin if empty)")
	flag.Parse()

	var reader io.Reader
	if *filePath == "" {
		// No file provided? Read from stdin
		reader = os.Stdin
	} else {
		file, err := os.Open(*filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to open file %q: %v\n", *filePath, err)
			os.Exit(1)
		}
		defer file.Close()
		reader = file
	}

	run(reader)
}

// Run loops through 'buffer' from PC=0 until we reach the end of the buffer,
// disassembling each instruction in turn.
func run(buffer io.Reader) error {
	code := make([]byte, 1)
	for {
		_, err := buffer.Read(code)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		disassemble(buffer, code[0])
	}

	return nil
}

// getByteAddress returns the next byte from the buffer as an uint8 address.
func getByteAddress(buffer io.Reader) (uint8, error) {
	address := make([]byte, 1)
	_, err := buffer.Read(address)
	if err != nil {
		return 0, err
	}

	return address[0], nil
}

// getWordAddress increments opBytes by 2 and returns the next two bytes.
func getWordAddress(buffer io.Reader) (uint16, error) {
	address := make([]byte, 2)
	_, err := buffer.Read(address)
	if err != nil {
		return 0, err
	}

	return cpu6510.ConvertTwoBytesToAddress(address[1], address[0]), nil
}

// zeroPage reads the next byte from the buffer and prints the mnemonic with
// the address in zero page.
func zeroPage(buffer io.Reader, mnemonic string) {
	address, err := getByteAddress(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s $%02X\n", mnemonic, address)
}

// zeroPageX reads the next byte from the buffer and prints the mnemonic with
// the address in zero page indexed by X.
func zeroPageX(buffer io.Reader, mnemonic string) {
	address, err := getByteAddress(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s $%02X, X\n", mnemonic, address)
}

// zeroPageY reads the next byte from the buffer and prints the mnemonic with
// the address in zero page indexed by Y.
func zeroPageY(buffer io.Reader, mnemonic string) {
	address, err := getByteAddress(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s $%02X, Y\n", mnemonic, address)
}

// immediate reads the next byte from the buffer and prints the mnemonic with
// the address in immediate mode.
func immediate(buffer io.Reader, mnemonic string) {
	address, err := getByteAddress(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s #$%02X\n", mnemonic, address)
}

// indirect reads the next two bytes from the buffer and prints the mnemonic
// with the address in indirect mode.
func indirect(buffer io.Reader, mnemonic string) {
	address, err := getWordAddress(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s (%08b)\n", mnemonic, address)
}

// indexedIndirectX reads the next byte from the buffer and prints the mnemonic
// with the address in indexed indirect X mode.
func indexedIndirectX(buffer io.Reader, mnemonic string) {
	indexedIndirect(buffer, mnemonic, "X")
}

// indexedIndirectY reads the next byte from the buffer and prints the mnemonic
// with the address in indexed indirect Y mode.
func indexedIndirectY(buffer io.Reader, mnemonic string) {
	indexedIndirect(buffer, mnemonic, "Y")
}

func indexedIndirect(buffer io.Reader, mnemonic string, register string) {
	address, err := getByteAddress(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s ($%02X), %s\n", mnemonic, address, register)
}

// absolute reads the next two bytes from the buffer and prints the mnemonic
// with the address in absolute mode.
func absolute(buffer io.Reader, mnemonic string) {
	address, err := getWordAddress(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s $%04X\n", mnemonic, address)
}

// absoluteX reads the next two bytes from the buffer and prints the mnemonic
// with the address in absolute mode indexed by X.
func absoluteX(buffer io.Reader, mnemonic string) {
	address, err := getWordAddress(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s $%04X, X\n", mnemonic, address)
}

// absoluteY reads the next two bytes from the buffer and prints the mnemonic
// with the address in absolute mode indexed by Y.
func absoluteY(buffer io.Reader, mnemonic string) {
	address, err := getWordAddress(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s $%04X, Y\n", mnemonic, address)
}

// relative reads the next byte from the buffer and prints the mnemonic with
// the relative address.
func relative(buffer io.Reader, mnemonic string) {
	address, err := getByteAddress(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s $%02X\n", mnemonic, address)
}

// accumulator prints the mnemonic with the accumulator register.
func accumulator(mnemonic string) {
	fmt.Printf("%s A\n", mnemonic)
}

// Disassemble reads the opcode at 'pc' from 'buffer', prints its assembly, and
// returns how many bytes that opcode consumed (i.e. how much to advance 'pc').
func disassemble(buffer io.Reader, code byte) {
	switch code {
	case 0x00:
		fmt.Println("BRK")
	case 0x01:
		indexedIndirectX(buffer, "ORA")
	case 0x05:
		zeroPage(buffer, "ORA")
	case 0x06:
		zeroPage(buffer, "ASL")
	case 0x08:
		fmt.Println("PHP")
	case 0x09:
		immediate(buffer, "ORA")
	case 0x0A:
		accumulator("ASL")
	case 0x0D:
		absolute(buffer, "ORA")
	case 0x0E:
		absolute(buffer, "ASL")
	case 0x10:
		relative(buffer, "BPL")
	case 0x11:
		indexedIndirectY(buffer, "ORA")
	case 0x15:
		zeroPageX(buffer, "ORA")
	case 0x16:
		zeroPageX(buffer, "ASL")
	case 0x18:
		fmt.Println("CLC")
	case 0x19:
		absoluteY(buffer, "ORA")
	case 0x1D:
		absoluteX(buffer, "ORA")
	case 0x1E:
		absoluteX(buffer, "ASL")
	case 0x20:
		absolute(buffer, "JSR")
	case 0x21:
		indexedIndirectX(buffer, "AND")
	case 0x24:
		zeroPage(buffer, "BIT")
	case 0x25:
		zeroPage(buffer, "AND")
	case 0x26:
		zeroPage(buffer, "ROL")
	case 0x28:
		fmt.Println("PLP")
	case 0x29:
		immediate(buffer, "AND")
	case 0x2A:
		accumulator("ROL")
	case 0x2C:
		absolute(buffer, "BIT")
	case 0x2D:
		absolute(buffer, "AND")
	case 0x2E:
		absolute(buffer, "ROL")
	case 0x30:
		relative(buffer, "BMI")
	case 0x31:
		indexedIndirectY(buffer, "AND")
	case 0x35:
		zeroPageX(buffer, "AND")
	case 0x36:
		zeroPageX(buffer, "ROL")
	case 0x38:
		fmt.Println("SEC")
	case 0x39:
		absoluteY(buffer, "AND")
	case 0x3D:
		absoluteX(buffer, "AND")
	case 0x3E:
		absoluteX(buffer, "ROL")
	case 0x40:
		fmt.Println("RTI")
	case 0x41:
		indexedIndirectX(buffer, "EOR")
	case 0x45:
		zeroPage(buffer, "EOR")
	case 0x46:
		zeroPage(buffer, "LSR")
	case 0x48:
		fmt.Println("PHA")
	case 0x49:
		immediate(buffer, "EOR")
	case 0x4A:
		accumulator("LSR")
	case 0x4C:
		absolute(buffer, "JMP")
	case 0x4D:
		absolute(buffer, "EOR")
	case 0x4E:
		absolute(buffer, "LSR")
	case 0x50:
		relative(buffer, "BVC")
	case 0x51:
		indexedIndirectY(buffer, "EOR")
	case 0x55:
		zeroPageX(buffer, "EOR")
	case 0x56:
		zeroPageX(buffer, "LSR")
	case 0x58:
		fmt.Println("CLI")
	case 0x59:
		absoluteY(buffer, "EOR")
	case 0x5D:
		absoluteX(buffer, "EOR")
	case 0x5E:
		absoluteX(buffer, "LSR")
	case 0x60:
		fmt.Println("RTS")
	case 0x61:
		indexedIndirectX(buffer, "ADC")
	case 0x65:
		zeroPage(buffer, "ADC")
	case 0x66:
		zeroPage(buffer, "ROR")
	case 0x68:
		fmt.Println("PLA")
	case 0x69:
		immediate(buffer, "ADC")
	case 0x6A:
		fmt.Println("ROR A")
	case 0x6C:
		indirect(buffer, "JMP")
	case 0x6D:
		absolute(buffer, "ADC")
	case 0x6E:
		absolute(buffer, "ROR")
	case 0x70:
		relative(buffer, "BVS")
	case 0x71:
		indexedIndirectY(buffer, "ADC")
	case 0x75:
		zeroPageX(buffer, "ADC")
	case 0x76:
		zeroPageX(buffer, "ROR")
	case 0x78:
		fmt.Println("SEI")
	case 0x79:
		absoluteY(buffer, "ADC")
	case 0x7D:
		absoluteX(buffer, "ADC")
	case 0x7E:
		absoluteX(buffer, "ROR")
	case 0x81:
		indexedIndirectX(buffer, "STA")
	case 0x84:
		zeroPage(buffer, "STY")
	case 0x85:
		zeroPage(buffer, "STA")
	case 0x86:
		zeroPage(buffer, "STX")
	case 0x88:
		fmt.Println("DEY")
	case 0x8A:
		fmt.Println("TXA")
	case 0x8C:
		absolute(buffer, "STY")
	case 0x8D:
		absolute(buffer, "STA")
	case 0x8E:
		absolute(buffer, "STX")
	case 0x90:
		relative(buffer, "BCC")
	case 0x91:
		indexedIndirectY(buffer, "STA")
	case 0x94:
		zeroPageX(buffer, "STY")
	case 0x95:
		zeroPageX(buffer, "STA")
	case 0x96:
		zeroPageY(buffer, "STX")
	case 0x98:
		fmt.Println("TYA")
	case 0x99:
		absoluteY(buffer, "STA")
	case 0x9A:
		fmt.Println("TXS")
	case 0x9D:
		absoluteX(buffer, "STA")
	case 0xA0:
		immediate(buffer, "LDY")
	case 0xA1:
		indexedIndirectX(buffer, "LDA")
	case 0xA2:
		immediate(buffer, "LDX")
	case 0xA4:
		zeroPage(buffer, "LDY")
	case 0xA5:
		zeroPage(buffer, "LDA")
	case 0xA6:
		zeroPage(buffer, "LDX")
	case 0xA8:
		fmt.Println("TAY")
	case 0xA9:
		immediate(buffer, "LDA")
	case 0xAA:
		fmt.Println("TAX")
	case 0xAC:
		absolute(buffer, "LDY")
	case 0xAD:
		absolute(buffer, "LDA")
	case 0xAE:
		absolute(buffer, "LDX")
	case 0xB0:
		relative(buffer, "BCS")
	case 0xB1:
		indexedIndirectY(buffer, "LDA")
	case 0xB4:
		zeroPageX(buffer, "LDY")
	case 0xB5:
		zeroPageX(buffer, "LDA")
	case 0xB6:
		zeroPageY(buffer, "LDX")
	case 0xB8:
		fmt.Println("CLV")
	case 0xB9:
		absoluteY(buffer, "LDA")
	case 0xBA:
		fmt.Println("TSX")
	case 0xBC:
		absoluteX(buffer, "LDY")
	case 0xBD:
		absoluteX(buffer, "LDA")
	case 0xBE:
		absoluteY(buffer, "LDX")
	case 0xC0:
		immediate(buffer, "CPY")
	case 0xC1:
		indexedIndirectX(buffer, "CMP")
	case 0xC4:
		zeroPage(buffer, "CPY")
	case 0xC5:
		zeroPage(buffer, "CMP")
	case 0xC6:
		zeroPage(buffer, "DEC")
	case 0xC8:
		fmt.Println("INY")
	case 0xC9:
		immediate(buffer, "CMP")
	case 0xCA:
		fmt.Println("DEX")
	case 0xCC:
		absolute(buffer, "CPY")
	case 0xCD:
		absolute(buffer, "CMP")
	case 0xCE:
		absolute(buffer, "DEC")
	case 0xD0:
		relative(buffer, "BNE")
	case 0xD1:
		indexedIndirectY(buffer, "CMP")
	case 0xD5:
		zeroPageX(buffer, "CMP")
	case 0xD6:
		zeroPageX(buffer, "DEC")
	case 0xD8:
		fmt.Println("CLD")
	case 0xD9:
		absoluteY(buffer, "CMP")
	case 0xDD:
		absoluteX(buffer, "CMP")
	case 0xDE:
		absoluteX(buffer, "DEC")
	case 0xE0:
		immediate(buffer, "CPX")
	case 0xE1:
		indexedIndirectX(buffer, "SBC")
	case 0xE4:
		zeroPage(buffer, "CPX")
	case 0xE5:
		zeroPage(buffer, "SBC")
	case 0xE6:
		zeroPage(buffer, "INC")
	case 0xE8:
		fmt.Println("INX")
	case 0xE9:
		immediate(buffer, "SBC")
	case 0xEA:
		fmt.Println("NOP")
	case 0xEC:
		absolute(buffer, "CPX")
	case 0xED:
		absolute(buffer, "SBC")
	case 0xEE:
		absolute(buffer, "INC")
	case 0xF0:
		relative(buffer, "BEQ")
	case 0xF1:
		indexedIndirectY(buffer, "SBC")
	case 0xF5:
		zeroPageX(buffer, "SBC")
	case 0xF6:
		zeroPageX(buffer, "INC")
	case 0xF8:
		fmt.Println("SED")
	case 0xF9:
		absoluteY(buffer, "SBC")
	case 0xFD:
		absoluteX(buffer, "SBC")
	case 0xFE:
		absoluteX(buffer, "INC")
	default:
		message := fmt.Sprintf("Unknown operation code, %x", code)
		panic(message)
	}
}
