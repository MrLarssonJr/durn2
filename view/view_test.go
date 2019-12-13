package view

import (
	"durn2/test_util"
	"fmt"
	"testing"
)

// Test that a factory creates view with supplied defaults.
func TestFactory(t *testing.T) {
	//Arrange
	title := "foo"
	path := "bar"
	f := NewFactory(title, path)
	//Act
	v := f.New("baz")
	//Assert
	if v.GetTitle() != title {
		t.Errorf("Factory produces view with incorrect title")
	}
	if v.templatePath != path {
		t.Errorf("Factory produces view with incorrect path")
	}
}

func TestView_Render(t *testing.T) {
	//Arrange
	title := "foo"
	text := "bar"
	path := "./res_test"
	template := "test.gohtml"
	f := NewFactory(title, path)
	v := f.New(template).Data(text)
	w := test_util.NewMockResponseWriter()
	//Act
	v.Render(&w)
	//Assert
	if w.GetWrittenText() != fmt.Sprintf("%s\n%s", title, text) {
		t.Errorf("View do not render as expected")
	}
}
