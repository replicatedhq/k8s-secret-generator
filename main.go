package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"log"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	var name = flag.String("name", "", "secret name")
	var key = flag.String("key", "data", "secret key")
	var namespace = flag.String("namespace", "default", "namespace the secret will be deployed to")
	var length = flag.Int("length", 32, "number of bytes in the secret")
	var base64Encode = flag.Bool("base64encode", false, "encode as base64 after generating")

	flag.Parse()

	if *name == "" {
		log.Fatal("name is required")
	}

	b := make([]byte, *length)
	if _, err := rand.Read(b); err != nil {
		log.Panicf("Generate secret: %v", err)
	}

	config, err := rest.InClusterConfig()
	if err != nil {
		log.Panicf("Get K8s client config from cluster environment: %v", err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Panicf("Make new K8s client: %v", err)
	}

	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name: *name,
		},
	}
	if *base64Encode {
		// K8s will apply another round of base64 encoding to store the secret as data
		secret.StringData = map[string]string{
			*key: base64.StdEncoding.EncodeToString(b),
		}
	} else {
		secret.Data = map[string][]byte{
			*key: b,
		}
	}
	if _, err = clientset.CoreV1().Secrets(*namespace).Create(secret); err != nil {
		log.Panicf("Create Kubernetes Secret: %v", err)
	}
	log.Printf("Created secret %q in namespace %q", *name, *namespace)
}
