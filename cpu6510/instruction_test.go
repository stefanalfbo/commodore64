package cpu6510

import (
	"testing"
)

func TestConvertTwoBytesToAddress(t *testing.T) {
	tests := []struct {
		high     byte
		low      byte
		expected uint16
	}{
		{0x03, 0x00, 0x0300},
		{0x00, 0xFF, 0x00FF},
		{0xFF, 0xFF, 0xFFFF},
		{0x12, 0x34, 0x1234},
		{0xAB, 0xCD, 0xABCD},
		{0x00, 0x00, 0x0000},
		{0x02, 0x00, 0x0200},
	}

	for _, test := range tests {
		address := ConvertTwoBytesToAddress(test.high, test.low)
		if address != test.expected {
			t.Errorf("Address should be %04x, got %04x for high:%02x low:%02x", test.expected, address, test.high, test.low)
		}
	}
}

func TestBRK(t *testing.T) {
	cpu := NewCPU()
	expectedPC := cpu.programCounter + 2

	cpu.execute(InstructionAsHex("BRK"))

	if !cpu.statusRegister.interruptDisableFlag {
		t.Errorf("Interrupt disable flag should be set")
	}

	if !cpu.statusRegister.breakCommandFlag {
		t.Errorf("Break status flag should be set")
	}

	if cpu.programCounter != expectedPC {
		t.Errorf("Program counter should be incremented by 2")
	}
}

func TestRTS(t *testing.T) {
	t.Run("Return from subroutine", func(t *testing.T) {
		cpu := NewCPU()
		cpu.ram[0x01FE] = 0x03
		cpu.ram[0x01FF] = 0x00
		cpu.stackPointer = 0xFD

		cpu.execute(InstructionAsHex("RTS"))

		if cpu.programCounter != 0x0004 {
			t.Errorf("Program counter should be set to the address on the stack")
		}
	})
}

func TestNOP(t *testing.T) {
	cpu := NewCPU()
	expectedPC := cpu.programCounter + 1

	cpu.execute(InstructionAsHex("NOP"))

	if cpu.programCounter != expectedPC {
		t.Errorf("Program counter should be incremented")
	}
}
