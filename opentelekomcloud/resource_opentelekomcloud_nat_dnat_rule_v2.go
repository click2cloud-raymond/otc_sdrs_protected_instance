// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file at
//     https://www.github.com/huaweicloud/magic-modules
//
// ----------------------------------------------------------------------------

package opentelekomcloud

import (
	"fmt"
	"log"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/huaweicloud/golangsdk"
)

func resourceNatDnatRuleV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceNatDnatRuleCreate,
		Read:   resourceNatDnatRuleRead,
		Delete: resourceNatDnatRuleDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"floating_ip_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"internal_service_port": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},

			"nat_gateway_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"port_id": {
				Type:          schema.TypeString,
				ConflictsWith: []string{"private_ip"},
				Optional:      true,
				ForceNew:      true,
			},

			"private_ip": {
				Type:          schema.TypeString,
				ConflictsWith: []string{"port_id"},
				Optional:      true,
				ForceNew:      true,
			},

			"protocol": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"external_service_port": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},

			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"floating_ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"tenant_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceNatDnatUserInputParams(d *schema.ResourceData) map[string]interface{} {
	return map[string]interface{}{
		"external_service_port": d.Get("external_service_port"),
		"floating_ip_id":        d.Get("floating_ip_id"),
		"internal_service_port": d.Get("internal_service_port"),
		"nat_gateway_id":        d.Get("nat_gateway_id"),
		"port_id":               d.Get("port_id"),
		"private_ip":            d.Get("private_ip"),
		"protocol":              d.Get("protocol"),
	}
}

func resourceNatDnatRuleCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sdkClient(GetRegion(d, config), "nat", serviceProjectLevel)
	if err != nil {
		return fmt.Errorf("Error creating sdk client, err=%s", err)
	}

	opts := resourceNatDnatUserInputParams(d)

	params := make(map[string]interface{})

	floatingIPIDProp, err := navigateValue(opts, []string{"floating_ip_id"}, nil)
	if err != nil {
		return err
	}
	e, err := isEmptyValue(reflect.ValueOf(floatingIPIDProp))
	if err != nil {
		return err
	}
	if !e {
		params["floating_ip_id"] = floatingIPIDProp
	}

	internalServicePortProp, err := navigateValue(opts, []string{"internal_service_port"}, nil)
	if err != nil {
		return err
	}
	e, err = isEmptyValue(reflect.ValueOf(internalServicePortProp))
	if err != nil {
		return err
	}
	if !e {
		params["internal_service_port"] = internalServicePortProp
	}

	externalServicePortProp, err := navigateValue(opts, []string{"external_service_port"}, nil)
	if err != nil {
		return err
	}
	e, err = isEmptyValue(reflect.ValueOf(externalServicePortProp))
	if err != nil {
		return err
	}
	if !e {
		params["external_service_port"] = externalServicePortProp
	}

	natGatewayIDProp, err := navigateValue(opts, []string{"nat_gateway_id"}, nil)
	if err != nil {
		return err
	}
	e, err = isEmptyValue(reflect.ValueOf(natGatewayIDProp))
	if err != nil {
		return err
	}
	if !e {
		params["nat_gateway_id"] = natGatewayIDProp
	}

	portIDProp, err := navigateValue(opts, []string{"port_id"}, nil)
	if err != nil {
		return err
	}
	e, err = isEmptyValue(reflect.ValueOf(portIDProp))
	if err != nil {
		return err
	}
	if !e {
		params["port_id"] = portIDProp
	}

	privateIPProp, err := navigateValue(opts, []string{"private_ip"}, nil)
	if err != nil {
		return err
	}
	e, err = isEmptyValue(reflect.ValueOf(privateIPProp))
	if err != nil {
		return err
	}
	if !e {
		params["private_ip"] = privateIPProp
	}

	protocolProp, err := navigateValue(opts, []string{"protocol"}, nil)
	if err != nil {
		return err
	}
	e, err = isEmptyValue(reflect.ValueOf(protocolProp))
	if err != nil {
		return err
	}
	if !e {
		params["protocol"] = protocolProp
	}

	log.Printf("[DEBUG] Creating new Dnat: %#v", params)

	url, err := replaceVars(d, "dnat_rules", nil)
	if err != nil {
		return err
	}
	url = client.ServiceURL(url)

	r := golangsdk.Result{}
	_, r.Err = client.Post(
		url,
		&map[string]interface{}{"dnat_rule": params},
		&r.Body,
		&golangsdk.RequestOpts{OkCodes: successHTTPCodes})
	if r.Err != nil {
		return fmt.Errorf("Error creating Dnat: %s", r.Err)
	}

	id, err := navigateValue(r.Body, []string{"dnat_rule", "id"}, nil)
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id.(string))

	return resourceNatDnatRuleRead(d, meta)
}

func resourceNatDnatRuleRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sdkClient(GetRegion(d, config), "nat", serviceProjectLevel)
	if err != nil {
		return fmt.Errorf("Error creating sdk client, err=%s", err)
	}

	url, err := replaceVars(d, "dnat_rules/{id}", nil)
	if err != nil {
		return err
	}
	url = client.ServiceURL(url)

	r := golangsdk.Result{}
	_, r.Err = client.Get(
		url, &r.Body,
		&golangsdk.RequestOpts{MoreHeaders: map[string]string{"Accept": "application/json"}})
	if r.Err != nil {
		return fmt.Errorf("Error reading %s: %s", fmt.Sprintf("NatDnat %q", d.Id()), r.Err)
	}
	v, ok := r.Body.(map[string]interface{})
	if !ok {
		return fmt.Errorf("Error reading %s: the result is not map", fmt.Sprintf("NatDnat %q", d.Id()))
	}

	res := map[string]interface{}{"read": v}

	opts := resourceNatDnatUserInputParams(d)

	createdATProp, err := navigateValue(res, []string{"read", "dnat_rule", "created_at"}, nil)
	if err != nil {
		return fmt.Errorf("Error reading Dnat:created_at, err: %s", err)
	}
	if err = d.Set("created_at", createdATProp); err != nil {
		return fmt.Errorf("Error setting Dnat:created_at, err: %s", err)
	}

	floatingIPIDProp, ok := opts["floating_ip_id"]
	if floatingIPIDProp != nil {
		ok, _ = isEmptyValue(reflect.ValueOf(floatingIPIDProp))
		ok = !ok
	}
	if !ok {
		floatingIPIDProp, err = navigateValue(res, []string{"read", "dnat_rule", "floating_ip_id"}, nil)
		if err != nil {
			return fmt.Errorf("Error reading Dnat:floating_ip_id, err: %s", err)
		}
		if err = d.Set("floating_ip_id", floatingIPIDProp); err != nil {
			return fmt.Errorf("Error setting Dnat:floating_ip_id, err: %s", err)
		}
	}

	floatingIPAddrProp, err := navigateValue(res, []string{"read", "dnat_rule", "floating_ip_address"}, nil)
	if err != nil {
		return fmt.Errorf("Error reading Dnat:floating_ip_address, err: %s", err)
	}
	if err = d.Set("floating_ip_address", floatingIPAddrProp); err != nil {
		return fmt.Errorf("Error setting Dnat:floating_ip_address, err: %s", err)
	}

	internalServicePortProp, ok := opts["internal_service_port"]
	if internalServicePortProp != nil {
		ok, _ = isEmptyValue(reflect.ValueOf(internalServicePortProp))
		ok = !ok
	}
	if !ok {
		internalServicePortProp, err = navigateValue(res, []string{"read", "dnat_rule", "internal_service_port"}, nil)
		if err != nil {
			return fmt.Errorf("Error reading Dnat:internal_service_port, err: %s", err)
		}
		if err = d.Set("internal_service_port", internalServicePortProp); err != nil {
			return fmt.Errorf("Error setting Dnat:internal_service_port, err: %s", err)
		}
	}

	externalServicePortProp, ok := opts["external_service_port"]
	if externalServicePortProp != nil {
		ok, _ = isEmptyValue(reflect.ValueOf(externalServicePortProp))
		ok = !ok
	}
	if !ok {
		externalServicePortProp, err = navigateValue(res, []string{"read", "dnat_rule", "external_service_port"}, nil)
		if err != nil {
			return fmt.Errorf("Error reading Dnat:external_service_port, err: %s", err)
		}
		if err = d.Set("external_service_port", externalServicePortProp); err != nil {
			return fmt.Errorf("Error setting Dnat:external_service_port, err: %s", err)
		}
	}

	natGatewayIDProp, ok := opts["nat_gateway_id"]
	if natGatewayIDProp != nil {
		ok, _ = isEmptyValue(reflect.ValueOf(natGatewayIDProp))
		ok = !ok
	}
	if !ok {
		natGatewayIDProp, err = navigateValue(res, []string{"read", "dnat_rule", "nat_gateway_id"}, nil)
		if err != nil {
			return fmt.Errorf("Error reading Dnat:nat_gateway_id, err: %s", err)
		}
		if err = d.Set("nat_gateway_id", natGatewayIDProp); err != nil {
			return fmt.Errorf("Error setting Dnat:nat_gateway_id, err: %s", err)
		}
	}

	portIDProp, ok := opts["port_id"]
	if portIDProp != nil {
		ok, _ = isEmptyValue(reflect.ValueOf(portIDProp))
		ok = !ok
	}
	if !ok {
		portIDProp, err = navigateValue(res, []string{"read", "dnat_rule", "port_id"}, nil)
		if err != nil {
			return fmt.Errorf("Error reading Dnat:port_id, err: %s", err)
		}
		if err = d.Set("port_id", portIDProp); err != nil {
			return fmt.Errorf("Error setting Dnat:port_id, err: %s", err)
		}
	}

	privateIPProp, ok := opts["private_ip"]
	if privateIPProp != nil {
		ok, _ = isEmptyValue(reflect.ValueOf(privateIPProp))
		ok = !ok
	}
	if !ok {
		privateIPProp, err = navigateValue(res, []string{"read", "dnat_rule", "private_ip"}, nil)
		if err != nil {
			return fmt.Errorf("Error reading Dnat:private_ip, err: %s", err)
		}
		if err = d.Set("private_ip", privateIPProp); err != nil {
			return fmt.Errorf("Error setting Dnat:private_ip, err: %s", err)
		}
	}

	protocolProp, ok := opts["protocol"]
	if protocolProp != nil {
		ok, _ = isEmptyValue(reflect.ValueOf(protocolProp))
		ok = !ok
	}
	if !ok {
		protocolProp, err = navigateValue(res, []string{"read", "dnat_rule", "protocol"}, nil)
		if err != nil {
			return fmt.Errorf("Error reading Dnat:protocol, err: %s", err)
		}
		if err = d.Set("protocol", protocolProp); err != nil {
			return fmt.Errorf("Error setting Dnat:protocol, err: %s", err)
		}
	}

	statusProp, err := navigateValue(res, []string{"read", "dnat_rule", "status"}, nil)
	if err != nil {
		return fmt.Errorf("Error reading Dnat:status, err: %s", err)
	}
	if err = d.Set("status", statusProp); err != nil {
		return fmt.Errorf("Error setting Dnat:status, err: %s", err)
	}

	tenantIDProp, err := navigateValue(res, []string{"read", "dnat_rule", "tenant_id"}, nil)
	if err != nil {
		return fmt.Errorf("Error reading Dnat:tenant_id, err: %s", err)
	}
	if err = d.Set("tenant_id", tenantIDProp); err != nil {
		return fmt.Errorf("Error setting Dnat:tenant_id, err: %s", err)
	}

	return nil
}

func resourceNatDnatRuleDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sdkClient(GetRegion(d, config), "nat", serviceProjectLevel)
	if err != nil {
		return fmt.Errorf("Error creating sdk client, err=%s", err)
	}

	url, err := replaceVars(d, "dnat_rules/{id}", nil)
	if err != nil {
		return err
	}
	url = client.ServiceURL(url)

	log.Printf("[DEBUG] Deleting Dnat %q", d.Id())
	r := golangsdk.Result{}
	_, r.Err = client.Delete(url, &golangsdk.RequestOpts{
		OkCodes:      []int{204},
		JSONResponse: nil,
		MoreHeaders:  map[string]string{"Content-Type": "application/json"},
	})
	if r.Err != nil {
		return fmt.Errorf("Error deleting Dnat %q: %s", d.Id(), r.Err)
	}

	return nil
}
