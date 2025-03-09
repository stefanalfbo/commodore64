package cpu6510

import "testing"

func TestNewCPU(t *testing.T) {
	cpu := NewCPU()

	if cpu.programCounter != 0 {
		t.Errorf("Program counter should be initialized to 0")
	}

	if cpu.statusRegister.carry {
		t.Errorf("Carry flag should be initialized to false")
	}

	for i := 0; i < memorySize; i++ {
		if cpu.ram[i] != 0 {
			t.Errorf("RAM should be initialized to 0")
		}
	}
}
