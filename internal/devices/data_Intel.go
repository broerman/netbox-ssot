// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapIntel = map[string]*DeviceData{
    "H2312XXKR2": {
        Manufacturer: "Intel",
        Model: "H2312XXKR2",
        Slug: "intel-h2312xxkr2",
        UHeight: 2,
        PartNumber: "",
        IsFullDepth: true,
        Airflow: "front-to-rear",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "parent",
        Weight: 21.5,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
            { Name: "PSU Slot 1", Label: "Power Supply 1", Position: "1" },
            { Name: "PSU Slot 2", Label: "Power Supply 2", Position: "2" },
        },
			  DeviceBays: []DeviceBay{
            { Name: "Node 1", Label: "Compute Node 1" },
            { Name: "Node 2", Label: "Compute Node 2" },
            { Name: "Node 3", Label: "Compute Node 3" },
            { Name: "Node 4", Label: "Compute Node 4" },
        },
        InventoryItems: []InventoryItem{
            { Name: "FHW12X35HS12G", Label: "12Gb SAS backplane", Manufacturer: "Intel", PartID: "FHW12X35HS12G" },
        },
        Interfaces: []Interface{
        },
    },
    "H2312XXLR2": {
        Manufacturer: "Intel",
        Model: "H2312XXLR2",
        Slug: "intel-h2312xxlr2",
        UHeight: 2,
        PartNumber: "",
        IsFullDepth: true,
        Airflow: "front-to-rear",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "parent",
        Weight: 30.6,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
            { Name: "PSU Slot 1", Label: "Power Supply 1", Position: "1" },
            { Name: "PSU Slot 2", Label: "Power Supply 2", Position: "2" },
        },
			  DeviceBays: []DeviceBay{
            { Name: "Node 1", Label: "Compute Node 1" },
            { Name: "Node 2", Label: "Compute Node 2" },
            { Name: "Node 3", Label: "Compute Node 3" },
            { Name: "Node 4", Label: "Compute Node 4" },
        },
        InventoryItems: []InventoryItem{
            { Name: "FHW12X35HS12G", Label: "12Gb SAS backplane", Manufacturer: "Intel", PartID: "FHW12X35HS12G" },
        },
        Interfaces: []Interface{
        },
    },
    "HNS2600TPR": {
        Manufacturer: "Intel",
        Model: "HNS2600TPR",
        Slug: "intel-hns2600tpr",
        UHeight: 0,
        PartNumber: "",
        IsFullDepth: false,
        Airflow: "front-to-rear",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "child",
        Weight: 1,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
            { Name: "PCIe Riser 1", Label: "PCIe Riser 1", Position: "1" },
            { Name: "PCIe Riser 2", Label: "PCIe Riser 2", Position: "2" },
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
            { Name: "S2600TPR", Label: "Intel Server Board S2600TPR", Manufacturer: "Intel", PartID: "" },
        },
        Interfaces: []Interface{
            { Name: "GigabitEthernet1", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "GigabitEthernet2", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "IPMI", Label: "", Type: "1000base-t", MgmtOnly: true },
            { Name: "QSFPPlus", Label: "", Type: "40gbase-x-qsfpp", MgmtOnly: false },
        },
    },
    "NUC11PAHi5": {
        Manufacturer: "Intel",
        Model: "NUC11PAHi5",
        Slug: "intel-nuc11pahi5",
        UHeight: 0,
        PartNumber: "",
        IsFullDepth: false,
        Airflow: "passive",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 1.19,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU1", Label: "", Type: "dc-terminal", MaximumDraw: 0, AllocatedDraw: 90 },
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
            { Name: "eth0", Label: "", Type: "2.5gbase-t", MgmtOnly: false },
        },
    },
    "NUC11TNBi5": {
        Manufacturer: "Intel",
        Model: "NUC11TNBi5",
        Slug: "intel-nuc11tnbi5",
        UHeight: 0,
        PartNumber: "",
        IsFullDepth: false,
        Airflow: "passive",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0.64,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU1", Label: "", Type: "dc-terminal", MaximumDraw: 0, AllocatedDraw: 0 },
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
            { Name: "eth0", Label: "", Type: "2.5gbase-t", MgmtOnly: false },
        },
    },
}
