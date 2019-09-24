// device.go
package model

type Device struct {
    Name string // INV_CB_A01
    Code string // mb-101
    Type string // inverter|envkit|genmeter|combiner
    Project int
}

// p20_mb_101_inverter
func (z Device) TableName() string {
    code := strings.Replace(z.Code, "-", "_")
    return fmt.Sprintf("p%d_%s_%s", z.Project, code, z.Type)
}

//----------------------------------------------------------
// envkit.go
package model

type Envkit struct {
    Device
}

//----------------------------------------------------------
// genmeter.go
package model

type Genmeter struct {
    Device
}

//----------------------------------------------------------
// inverter.go
package model

type Inverter struct {
    Device
}

//----------------------------------------------------------
// combiner.go
package model

type Combiner struct {
    Device
}

//----------------------------------------------------------
// project.go
package model

type Project struct {
	ID        int
	Name      string
	Devices   []*Device
	Envkits   []*Envkit
	Genmeters []*Genmeter
	Inverters []*Inverter
}
