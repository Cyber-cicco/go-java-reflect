package reflect

import (
	"context"
	"os"
	"testing"

	"github.com/Cyber-cicco/java-reflect/config"
)

func TestPackage(t *testing.T) {
    path := config.JAVA_DIR + "entites/Additif.java"
    file, err := os.ReadFile(path)

    if err != nil {
        t.Fatalf("expected no error, got %s", err) 
    }
    root, err := config.JavaParser.ParseCtx(context.TODO(), nil, file)

    document := &Document{
        root: root.RootNode(),
        content: file,
    }

    if document == nil {
        t.Fatalf("Wtf %p", document)
    }

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
    root, err = config.JavaParser.ParseCtx(context.TODO(), nil, errorCase)
    document = &Document{
        root: root.RootNode(),
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
