package file

import (
	"os"
	"testing"
)

var data = `
apiVersion: v1
kind: Pod
metadata:
  name: active-incident
  labels:
    app.kubernetes.io/name: active-incident
spec:
  volumes:
    - name: active-incident-data-volume
      persistentVolumeClaim:
        claimName: active-incident-storage
        readOnly: false
  containers:
    - name: active-incident-container
      image: us-central1-docker.pkg.dev/mchirico/public/activeincident:v0.1.3
      env:
      - name: MONGO_URI
        valueFrom:
          secretKeyRef:
            name: mongo-creds
            key: mongouri

      - name: OWM_API_KEY
        valueFrom:
          secretKeyRef:
            name: mongo-creds
            key: weatherkey

      - name: MONGO_WEATHER
        valueFrom:
          secretKeyRef:
            name: mongo-creds
            key: weathercollection


      volumeMounts:
        - mountPath: /etc/mongo
          name: active-incident-data-volume
  restartPolicy: Always

`

func CreatePodDeploymentFiles() {
	os.Remove("./testPod.yaml")
	os.Remove("./testDeployment.yaml")
	os.WriteFile("./testPod.yaml", podfile(), os.ModePerm)
	os.WriteFile("./testDeployment.yaml", deployment(), os.ModePerm)

}

func TestFindImages(t *testing.T) {
	CreatePodDeploymentFiles()
	pod, err := Pod("./testPod.yaml")
	if err != nil {
		t.Fatalf("pod: %v", err)
	}
	if pod.Name != "nginx" {
		t.FailNow()
	}


	_, err = Pod("./testDeployment.yaml")
	if err != nil {
		t.Fatalf("deployment: %v", err)
	}


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
