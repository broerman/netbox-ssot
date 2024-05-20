// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapInfotecs = map[string]*DeviceData{
    "ViPNet Coordinator HW1000C": {
        Manufacturer: "Infotecs",
        Model: "ViPNet Coordinator HW1000C",
        Slug: "infotecs-vipnet-coordinator-hw1000c",
        UHeight: 1,
        PartNumber: "",
        IsFullDepth: true,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU0", Label: "", Type: "iec-60320-c14", MaximumDraw: 200, AllocatedDraw: 0 },
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "ethernet 0", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "ethernet 1", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "ethernet 2", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "ethernet 3", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "ethernet 4", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "ethernet 5", Label: "", Type: "1000base-t", MgmtOnly: false },
        },
    },
}
