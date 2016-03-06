package dnsimple

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/randomeizer/contegix-classic"
)

func resourceVirtualMachine() *schema.Resource {
	return &schema.Resource{
		Create: resourceVirtualMachineCreate,
		Read:   resourceVirtualMachineRead,
		Update: resourceVirtualMachineUpdate,
		Delete: resourceVirtualMachineDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"ip_addresses": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     schema.TypeString,
				Computed: true,
			},

			"template_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"template_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"package_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"package_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"zone_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"zone_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"vm_tools_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceVirtualMachineCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*contegixclassic.Client)

	// Create the new record
	newVM := &contegixclassic.CreateVirtualMachine{
		Name:         d.Get("name").(string),
		ZoneUUID:     d.Get("zone_uuid").(string),
		PackageUUID:  d.Get("package_uuid").(string),
		TemplateUUID: d.Get("template_uuid").(string),
	}

	log.Printf("[DEBUG] Contegix Classic VM create configuration: %#v", newVM)

	vm, err := client.CreateVirtualMachine(newVM)

	if err != nil {
		return fmt.Errorf("Failed to create Contegix Classic VM: %s", err)
	}

	d.SetId(vm.UUID)
	log.Printf("[INFO] Virtual Machine ID: %s", d.Id())

	return updateVirtualMachineResourceData(d, vm)
}

func updateVirtualMachineResourceData(d *schema.ResourceData, vm *contegixclassic.VirtualMachine) error {
	d.Set("name", vm.Name)
	d.Set("uuid", vm.UUID)
	d.Set("state", vm.State)
	d.Set("ip_addresses", vm.IPAddresses)
	d.Set("template_name", vm.TemplateName)
	d.Set("template_uuid", vm.TemplateUUID)
	d.Set("package_name", vm.PackageName)
	d.Set("package_uuid", vm.PackageUUID)
	d.Set("zone_name", vm.ZoneName)
	d.Set("zone_uuid", vm.ZoneUUID)
	d.set("vm_tools_status", vm.VMToolsStatus)

	return nil
}

func resourceVirtualMachineRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*contegixclassic.Client)

	vm, err := client.GetVirtualMachine(d.Id())
	if err != nil {
		return fmt.Errorf("Couldn't find Contegix Classic VM: %s", err)
	}

	return updateVirtualMachineResourceData(d, vm)
}

func resourceVirtualMachineUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*dnsimple.Client)

	details = *contegixclassic.UpdateVirtualMachine{
		PackageUUID: d.Get("package_uuid"),
	}
	log.Printf("[DEBUG] Contegix Classic VM update configuration: %#v", details)

	vm, err := client.UpdateRecord(d.Id(), details)
	if err != nil {
		return fmt.Errorf("Failed to update Contegix Classic VM: %s", err)
	}

	return updateVirtualMachineResourceData(d, vm)
}

func resourceVirtualMachineDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*dnsimple.Client)

	log.Printf("[INFO] Deleting Contegix Classic VM: %s", d.Id())

	err := client.DestroyRecord(d.Id())

	if err != nil {
		return fmt.Errorf("Error deleting Contegix Classic VM: %s", err)
	}

	return nil
}
