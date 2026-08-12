package model

// FieldDescriptor documents one stable field exposed by reports and imports.
type FieldDescriptor struct {
	Name       string
	Label      string
	Kind       string
	Required   bool
	Searchable bool
	Exported   bool
}

var StandardFields = []FieldDescriptor{
	{Name: "id_1", Label: "Payment Ledger id 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "name_1", Label: "Payment Ledger name 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "status_1", Label: "Payment Ledger status 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "priority_1", Label: "Payment Ledger priority 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "amount_1", Label: "Payment Ledger amount 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "active_1", Label: "Payment Ledger active 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "version_1", Label: "Payment Ledger version 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "created_at_1", Label: "Payment Ledger created at 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "updated_at_1", Label: "Payment Ledger updated at 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "owner_1", Label: "Payment Ledger owner 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "region_1", Label: "Payment Ledger region 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "source_1", Label: "Payment Ledger source 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "category_1", Label: "Payment Ledger category 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "group_1", Label: "Payment Ledger group 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "channel_1", Label: "Payment Ledger channel 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "note_1", Label: "Payment Ledger note 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "external_id_1", Label: "Payment Ledger external id 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "tenant_1", Label: "Payment Ledger tenant 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "checksum_1", Label: "Payment Ledger checksum 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "revision_1", Label: "Payment Ledger revision 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "id_2", Label: "Payment Ledger id 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "name_2", Label: "Payment Ledger name 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "status_2", Label: "Payment Ledger status 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "priority_2", Label: "Payment Ledger priority 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "amount_2", Label: "Payment Ledger amount 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "active_2", Label: "Payment Ledger active 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "version_2", Label: "Payment Ledger version 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "created_at_2", Label: "Payment Ledger created at 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "updated_at_2", Label: "Payment Ledger updated at 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "owner_2", Label: "Payment Ledger owner 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "region_2", Label: "Payment Ledger region 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "source_2", Label: "Payment Ledger source 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "category_2", Label: "Payment Ledger category 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "group_2", Label: "Payment Ledger group 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "channel_2", Label: "Payment Ledger channel 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "note_2", Label: "Payment Ledger note 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "external_id_2", Label: "Payment Ledger external id 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "tenant_2", Label: "Payment Ledger tenant 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "checksum_2", Label: "Payment Ledger checksum 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "revision_2", Label: "Payment Ledger revision 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "id_3", Label: "Payment Ledger id 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "name_3", Label: "Payment Ledger name 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "status_3", Label: "Payment Ledger status 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "priority_3", Label: "Payment Ledger priority 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "amount_3", Label: "Payment Ledger amount 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "active_3", Label: "Payment Ledger active 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "version_3", Label: "Payment Ledger version 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "created_at_3", Label: "Payment Ledger created at 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "updated_at_3", Label: "Payment Ledger updated at 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "owner_3", Label: "Payment Ledger owner 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "region_3", Label: "Payment Ledger region 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "source_3", Label: "Payment Ledger source 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "category_3", Label: "Payment Ledger category 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "group_3", Label: "Payment Ledger group 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "channel_3", Label: "Payment Ledger channel 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "note_3", Label: "Payment Ledger note 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "external_id_3", Label: "Payment Ledger external id 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "tenant_3", Label: "Payment Ledger tenant 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "checksum_3", Label: "Payment Ledger checksum 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "revision_3", Label: "Payment Ledger revision 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "id_4", Label: "Payment Ledger id 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "name_4", Label: "Payment Ledger name 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "status_4", Label: "Payment Ledger status 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "priority_4", Label: "Payment Ledger priority 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "amount_4", Label: "Payment Ledger amount 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "active_4", Label: "Payment Ledger active 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "version_4", Label: "Payment Ledger version 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "created_at_4", Label: "Payment Ledger created at 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "updated_at_4", Label: "Payment Ledger updated at 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "owner_4", Label: "Payment Ledger owner 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "region_4", Label: "Payment Ledger region 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "source_4", Label: "Payment Ledger source 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "category_4", Label: "Payment Ledger category 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "group_4", Label: "Payment Ledger group 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "channel_4", Label: "Payment Ledger channel 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "note_4", Label: "Payment Ledger note 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "external_id_4", Label: "Payment Ledger external id 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "tenant_4", Label: "Payment Ledger tenant 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "checksum_4", Label: "Payment Ledger checksum 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "revision_4", Label: "Payment Ledger revision 4", Kind: "string", Searchable: true, Exported: true},
}

func FieldByName(name string) (FieldDescriptor, bool) {
	for _, field := range StandardFields {
		if field.Name == name {
			return field, true
		}
	}
	return FieldDescriptor{}, false
}

func ExportedFieldNames() []string {
	result := make([]string, 0)
	for _, field := range StandardFields {
		if field.Exported {
			result = append(result, field.Name)
		}
	}
	return result
}
