package reflect

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/Cyber-cicco/java-reflect/config"
)

func testInit(t *testing.T) *Document {
    p := NewProject()
	path := config.JAVA_DIR + "entites/Additif.java"
	file, err := os.ReadFile(path)

	if err != nil {
		t.Fatalf("expected no error, got %s", err)
	}

	root, err := config.JavaParser.ParseCtx(context.TODO(), nil, file)
	absPath, err := filepath.Abs(path)

	if err != nil {
		t.Fatalf("Got unexpected error %s", err)
	}

	document, ok := p.NewDocument(root.RootNode(), absPath, file)
    fmt.Printf("document.diagnostics: %v\n", document.diagnostics)

    if !ok {
        t.Fatalf("Got unexpected errors %v", document.diagnostics)
    }

	if document == nil {
		t.Fatalf("Wtf %p", document)
	}
	return document
}

func TestPackage(t *testing.T) {
	document := testInit(t)
	exp := "fr.diginamic.entites"
	p := document.GetPackage()

	if p.GetFullScope() != exp {
		t.Fatalf("Expected %s, got %s", exp, p.GetFullScope())
	}

	errorCase := []byte(`
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;

@Entity
public class Additif extends BaseEntity{

}
`)
	root, _ := config.JavaParser.ParseCtx(context.TODO(), nil, errorCase)
    project := NewProject()
    document, ok := project.NewDocument(root.RootNode(), "", errorCase)

    if ok {
        t.Fatalf("Expected error, got valid document %v", document)
    }

    diag, ok := document.diagnostics[config.DIAG_IMPS]

	if ok {
		t.Fatalf("Error was expected to be %s, got %v", diag, document.diagnostics)
	}

}

func TestGetMainClass(t *testing.T) {

	document := testInit(t)
	main := document.GetMainClass()

	expected := "Additif"

	switch c := main.(type) {
	case *Class:
		actual := c.GetDeclaredName()

		if actual != expected {
			t.Fatalf("Error : expected %s, got %s", actual, expected)
		}
	}

}

func TestGetImports(t *testing.T) {
	document := testInit(t)
	imps := document.GetImports()

	expectedLen := 4

	if len(imps) != expectedLen {
		t.Fatalf("Error : expected %d, got %d", expectedLen, len(imps))
	}

	exp := "jakarta.persistence.Entity"

    fmt.Printf("imps: %v\n", imps)
	if imps[0].ToString() != exp {
		t.Fatalf("Error : expected %s, got %s", exp, imps[0].ToString())
	}
}
