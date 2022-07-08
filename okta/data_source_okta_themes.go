package okta

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceThemes() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceThemesRead,
		Schema:      themesDataSourceSchema,
	}
}

func dataSourceThemesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var err error
	brandID, ok := d.GetOk("brand_id")
	if !ok {
		return diag.Errorf("brand_id required for themes: %v", err)
	}

	themes, _, err := getOktaClientFromMetadata(m).Brand.ListBrandThemes(ctx, brandID.(string))
	if err != nil {
		return diag.Errorf("failed to list brand themes: %v", err)
	}

	d.SetId(fmt.Sprintf("themes-%s", brandID.(string)))
	arr := make([]interface{}, len(themes))
	for i, theme := range themes {
		rawMap := flattenTheme("", theme)
		arr[i] = rawMap
	}

	themesDataSource := &schema.Resource{
		Schema: themeDataSourceSchema,
	}
	err = d.Set("themes", schema.NewSet(schema.HashResource(themesDataSource), arr))
	if err != nil {
		return diag.Errorf("failed to set themes: %v", err)
	}

	return nil
}
