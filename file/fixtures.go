package file

import (
	"os"
)

func CreatePodDeploymentFiles() {
	os.Remove("./testPod.yaml")
	os.Remove("./testDeployment.yaml")
	os.WriteFile("./testPod.yaml", podfile(), os.ModePerm)
	os.WriteFile("./testDeployment.yaml", deployment(), os.ModePerm)

}

func podfile() []byte {
	return []byte(`apiVersion: v1
kind: Pod
metadata:
  name: nginx
spec:
  containers:
  - name: nginx
    image: nginx:1.14.2
    ports:
    - containerPort: 80
`)
}

func deployment() []byte {
	return []byte(`apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
spec:
  replicas: 3
  selector:
    matchLabels:
    app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.14.2
        ports:
        - containerPort: 80		
`)
}
