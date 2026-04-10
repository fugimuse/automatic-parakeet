package home_test

import (
	"os"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
	"github.com/fugimuse/automatic-parakeet/home"
)

func TestHomeView(t *testing.T) {
	homeModel := home.NewModel(80, 50, "Budget App", []string{"q: Quit"})
	approvals.VerifyString(t, homeModel.View().Content)
}

func TestMain(m *testing.M) {
	approvals.UseFolder("testdata")
	os.Exit(m.Run())
}
