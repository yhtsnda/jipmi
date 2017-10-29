package jipmi

type mc_guid_req struct {
	netfn_lun  byte
	cmd        byte
	target_cmd byte
}

type mc_guid_res struct {
}
