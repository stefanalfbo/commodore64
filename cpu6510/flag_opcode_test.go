package cpu6510

import "testing"

func TestCLC(t *testing.T) {
	cpu := NewCPU()
	expectedPC := cpu.programCounter + 1
	cpu.statusRegister.carryFlag = true

	cpu.execute(OpCodeAsHex("CLC"))

	if cpu.statusRegister.carryFlag {
		t.Errorf("Carry flag should be cleared")
	}

	if cpu.programCounter != expectedPC {
		t.Errorf("Program counter should be incremented")
	}
}

func TestCLD(t *testing.T) {
	cpu := NewCPU()
	expectedPC := cpu.programCounter + 1
	cpu.statusRegister.decimalModeFlag = true

	cpu.execute(OpCodeAsHex("CLD"))

	if cpu.statusRegister.decimalModeFlag {
		t.Errorf("Decimal mode flag should be cleared")
	}

	if cpu.programCounter != expectedPC {
		t.Errorf("Program counter should be incremented")
	}
}

func TestSEC(t *testing.T) {
	cpu := NewCPU()
	expectedPC := cpu.programCounter + 1

	cpu.execute(OpCodeAsHex("SEC"))

	if !cpu.statusRegister.carryFlag {
		t.Errorf("Carry flag should be set")
	}

	if cpu.programCounter != expectedPC {
		t.Errorf("Program counter should be incremented")
	}
}

func TestSED(t *testing.T) {
	cpu := NewCPU()
	expectedPC := cpu.programCounter + 1

	cpu.execute(OpCodeAsHex("SED"))

	if !cpu.statusRegister.decimalModeFlag {
		t.Errorf("Decimal mode flag should be set")
	}

	if cpu.programCounter != expectedPC {
		t.Errorf("Program counter should be incremented")
	}
}

func TestCLI(t *testing.T) {
	cpu := NewCPU()
	expectedPC := cpu.programCounter + 1
	cpu.statusRegister.interruptDisableFlag = true

	cpu.execute(OpCodeAsHex("CLI"))

	if cpu.statusRegister.interruptDisableFlag {
		t.Errorf("Interrupt disable flag should be cleared")
	}

	if cpu.programCounter != expectedPC {
		t.Errorf("Program counter should be incremented")
	}
}

func TestCLV(t *testing.T) {
	cpu := NewCPU()
	expectedPC := cpu.programCounter + 1
	cpu.statusRegister.overflowFlag = true

	cpu.execute(OpCodeAsHex("CLV"))

	if cpu.statusRegister.overflowFlag {
		t.Errorf("Overflow flag should be cleared")
	}

	if cpu.programCounter != expectedPC {
		t.Errorf("Program counter should be incremented")
	}
}

func TestSEI(t *testing.T) {
	cpu := NewCPU()
	expectedPC := cpu.programCounter + 1

	cpu.execute(OpCodeAsHex("SEI"))

	if !cpu.statusRegister.interruptDisableFlag {
		t.Errorf("Interrupt disable flag should be set")
	}

	if cpu.programCounter != expectedPC {
		t.Errorf("Program counter should be incremented")
	}
}
