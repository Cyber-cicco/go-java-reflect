package reflect

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/Cyber-cicco/java-reflect/config"
)

func testInit(t *testing.T) *Document {
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

	document := NewDocument(root.RootNode(), absPath, file)

	if document == nil {
		t.Fatalf("Wtf %p", document)
	}
	return document
}

func TestPackage(t *testing.T) {
	document := testInit(t)
	exp := "fr.diginamic.entites"
	p, err := document.GetPackage()

	if err != nil {
		t.Fatalf("Got unexpected error %s", err)
	}

	if p != exp {
		t.Fatalf("Expected %s, got %s", exp, p)
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
	root, err := config.JavaParser.ParseCtx(context.TODO(), nil, errorCase)
	document = &Document{
		root:    root.RootNode(),
		content: errorCase,
	}

	p, err = document.GetPackage()

	if err == nil {
		t.Fatalf("Expected an error, got package %s", p)
	}

	exp = "Class doesn't have a package declaration"

	if err.Error() != exp {
		t.Fatalf("Error was expected to be %s, got %s", exp, err)
	}

}

func TestGetMainClass(t *testing.T) {

	document := testInit(t)
	main, err := document.GetMainClass()

	if err != nil {
		t.Fatalf("Got unexpected error %s", err)
	}

	expected := "Additif"
	actual := main.GetDeclaredName()

	if actual != expected {
        t.Fatalf("Error : expected %s, got %s", actual, expected)
	}

}
