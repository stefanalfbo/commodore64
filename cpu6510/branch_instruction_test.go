package cpu6510

import (
	"testing"
)

func TestBranchInstruction(t *testing.T) {
	programCounterStart := uint16(0xC002)

	tests := []struct {
		instruction    string
		statusRegister StatusRegister
		jumpAddress    byte
		expectedPC     uint16
	}{
		{"BPL", StatusRegister{negativeFlag: false}, 0x40, 0xC044},
		{"BPL", StatusRegister{negativeFlag: true}, 0x40, 0xC003},
		{"BPL", StatusRegister{negativeFlag: false}, 0xFC, 0xC000},
		{"BMI", StatusRegister{negativeFlag: true}, 0x40, 0xC044},
		{"BMI", StatusRegister{negativeFlag: false}, 0x40, 0xC003},
		{"BMI", StatusRegister{negativeFlag: true}, 0xFC, 0xC000},
		{"BVC", StatusRegister{overflowFlag: false}, 0x40, 0xC044},
		{"BVC", StatusRegister{overflowFlag: true}, 0x40, 0xC003},
		{"BVC", StatusRegister{overflowFlag: false}, 0xFC, 0xC000},
		{"BVS", StatusRegister{overflowFlag: true}, 0x40, 0xC044},
		{"BVS", StatusRegister{overflowFlag: false}, 0x40, 0xC003},
		{"BVS", StatusRegister{overflowFlag: true}, 0xFC, 0xC000},
		{"BCC", StatusRegister{carryFlag: false}, 0x40, 0xC044},
		{"BCC", StatusRegister{carryFlag: true}, 0x40, 0xC003},
		{"BCC", StatusRegister{carryFlag: false}, 0xFC, 0xC000},
		{"BCS", StatusRegister{carryFlag: true}, 0x40, 0xC044},
		{"BCS", StatusRegister{carryFlag: false}, 0x40, 0xC003},
		{"BCS", StatusRegister{carryFlag: true}, 0xFC, 0xC000},
		{"BNE", StatusRegister{zeroFlag: false}, 0x40, 0xC044},
		{"BNE", StatusRegister{zeroFlag: true}, 0x40, 0xC003},
		{"BNE", StatusRegister{zeroFlag: false}, 0xFC, 0xC000},
		{"BEQ", StatusRegister{zeroFlag: true}, 0x40, 0xC044},
		{"BEQ", StatusRegister{zeroFlag: false}, 0x40, 0xC003},
		{"BEQ", StatusRegister{zeroFlag: true}, 0xFC, 0xC000},
	}

	for _, test := range tests {
		cpu := NewCPU()
		cpu.statusRegister = test.statusRegister
		cpu.programCounter = programCounterStart
		cpu.ram[cpu.programCounter+1] = test.jumpAddress

		cpu.execute(InstructionAsHex(test.instruction))

		if cpu.programCounter != test.expectedPC {
			t.Errorf("Expected PC to be 0x%02X, got 0x%02X for instruction %s", test.expectedPC, cpu.programCounter, test.instruction)
		}
	}

}
