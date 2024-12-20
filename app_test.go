package main

import "testing"

func TestGenerateBSN(t *testing.T) {
	app := &App{}

	for i := 1; i <= 50; i++ {
		bsn := app.GenerateBSN(false)

		t.Run("GenerateBSN:"+bsn, func(t *testing.T) {
			if !app.IsValidBSN(bsn) {
				t.Errorf("Expected BSN %s to be valid, but got invalid", bsn)
			}
		})
	}

	for i := 1; i <= 50; i++ {
		bsn := app.GenerateBSN(true)

		t.Run("GenerateBSN w/ leading zeros:"+bsn, func(t *testing.T) {
			if !app.IsValidBSN(bsn) {
				t.Errorf("Expected BSN %s to be valid, but got invalid", bsn)
			}
		})
	}
}

func TestIsValidBSN(t *testing.T) {
	app := &App{}

	validBSNs := []string{
		"000006944",
		"000003086",
		"000006154",
		"999993239",
		"000003803",
		"000007900",
		"999992089",
		"799890054",
		"000005083",
		"000001533",
		"000000681",
		"999999941",
		"999998882",
		"999994505",
		"999997282",
		"000006506",
		"000005642",
		"000002082",
		"000001156",
		"999990202",
		"000004583",
		"999999606",
		"000006580",
		"999995315",
		"000007006",
		"000002264",
		"000001880",
		"000005551",
		"000007560",
		"000002963",
		"000008230",
		"999994657",
		"999999357",
		"999993768",
		"000009611",
		"000000929",
		"999995224",
		"000005873",
		"000007869",
		"999992260",
		"999994438",
		"999996526",
		"999991541",
		"999999229",
		"000005137",
		"999994256",
		"000002501",
		"999993562",
		"999990779",
		"999993306",
	}

	invalidBSNs := []string{
		"00001",
		"0000000001",
		"1212121212",
		"999993237",
		"000003802",
		"000007901",
		"999992189",
		"799890053",
		"000015083",
		"000001532",
		"0100006121",
		"999999942",
		"998882",
		"999994504",
		"99997282",
		"000006505",
		"200081642",
		"000082083",
		"000081156",
		"999080202",
		"000081880",
		"000085551",
		"000087560",
		"000082963",
		"000088230",
		"999984657",
		"999989357",
		"999983768",
		"000089611",
		"000080929",
		"999985224",
		"000085873",
		"fake bsn",
		"999982260",
		"999984438",
		"9999aa652",
		"999981541",
		"999989229",
		"000085137",
		"99998425x",
		"000082501",
		"999983562",
		"999980779",
		"979983306",
		"070008230",
		"97999465a",
		"979999357",
		"979993768",
		"070009611",
		"070000929",
	}

	for _, bsn := range validBSNs {
		t.Run("ValidBSN:"+bsn, func(t *testing.T) {
			if !app.IsValidBSN(bsn) {
				t.Errorf("Expected BSN %s to be valid, but got invalid", bsn)
			}
		})
	}

	for _, bsn := range invalidBSNs {
		t.Run("ValidBSN:"+bsn, func(t *testing.T) {
			if app.IsValidBSN(bsn) {
				t.Errorf("Expected BSN %s not to be valid, but got valid", bsn)
			}
		})
	}
}
