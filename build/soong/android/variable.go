package android
type Product_variables struct {
	Additional_gralloc_10_usage_bits struct {
		Cppflags []string
	}
	Bootloader_message_offset struct {
		Cflags []string
	}
	Target_init_vendor_lib struct {
		Whole_static_libs []string
	}
 	Supports_hw_fde struct {
		Cflags []string
		Header_libs []string
		Shared_libs []string
	}
	Supports_hw_fde_perf struct {
		Cflags []string
	}
	Supports_extended_compress_format struct {
		Cflags []string
	}
}

type ProductVariables struct {
	Additional_gralloc_10_usage_bits  *string `json:",omitempty"`
	Bootloader_message_offset  *int `json:",omitempty"`
	Target_init_vendor_lib  *string `json:",omitempty"`
	Supports_extended_compress_format  *bool `json:",omitempty"`
	Supports_hw_fde  *bool `json:",omitempty"`
	Supports_hw_fde_perf  *bool `json:",omitempty"`
}
