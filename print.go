// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"

// 	"github.com/google/gousb"
// )

// func main() {
// 	// Initialize USB context
// 	usbDevices := GetAllUsb()
// 	for index, device := range usbDevices {
// 		fmt.Println(index+1, ":", device.Name+" "+device.VID.String()+" "+device.PID.String())
// 		// fmt.Println(device.PID.String())
// 		// fmt.Println(device.VID.String())
// 	}
// 	var input int
// 	fmt.Println("Pilih device(dengan angka): ")
// 	fmt.Scan(&input)
// 	// usbDevices[input-1].VID = gousb.ID("0x" + usbDevices[input-1].VID.String())
// 	ctx := gousb.NewContext()
// 	defer ctx.Close()
// 	// var vid gousb.ID = useb
// 	// Open the USB device
// 	vid := gousb.ID(usbDevices[input-1].VID)
// 	pid := gousb.ID(usbDevices[input-1].PID)

// 	dev, err := ctx.OpenDeviceWithVIDPID(vid, pid)
// 	if err != nil {
// 		fmt.Printf("OpenDeviceWithVIDPID failed: %s\n", err)
// 		return
// 	}
// 	defer dev.Close()

// 	// Read file to print
// 	fileData, err := ioutil.ReadFile("/home/fikri/Downloads/asd.pdf")
// 	if err != nil {
// 		fmt.Printf("ReadFile failed: %s\n", err)
// 		return
// 	}

// 	// Send file data to USB device1
// 	_, err = dev.Control(gousb.ControlOut|gousb.ControlClass|gousb.ControlInterface, 0x02, 0x0000, 0x0000, fileData)

// 	if err != nil {
// 		fmt.Printf("Control failed: %s\n", err)
// 		return
// 	}

// 	fmt.Println("File printed successfully!")

// }

// type Device struct {
// 	Name string
// 	VID  gousb.ID
// 	PID  gousb.ID
// }

// var devices []Device

// func GetAllUsb() []Device {
// 	// Initialize USB context
// 	ctx := gousb.NewContext()
// 	defer ctx.Close()

// 	// Enumerate USB devices
// 	devs, err := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {
// 		// Filter for all USB devices
// 		return desc.Class == gousb.ClassPerInterface
// 	})
// 	if err != nil {
// 		fmt.Printf("OpenDevices failed: %s\n", err)
// 		return nil
// 	}
// 	defer func() {
// 		for _, dev := range devs {
// 			dev.Close()
// 		}
// 	}()

// 	// Print information about each USB device

// 	for _, dev := range devs {
// 		desc := dev.Desc
// 		product, err := dev.Product()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		// fmt.Println("Device: " + product)
// 		// fmt.Printf("0x%x\n", desc.Vendor)
// 		// fmt.Println(desc.Vendor)
// 		// decimal, err := strconv.ParseInt(desc.Vendor.String(), 10, 64)
// 		// if err != nil {
// 		// 	fmt.Println("Invalid hexadecimal number:", desc.Vendor.String())
// 		// 	return nil
// 		// }
// 		// fmt.Println(decimal)

// 		device := Device{
// 			Name: product,
// 			VID:  desc.Vendor,
// 			PID:  desc.Product,
// 		}
// 		devices = append(devices, device)
// 	}

// 	return devices
// }
